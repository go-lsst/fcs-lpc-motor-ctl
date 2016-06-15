// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/go-lsst/ncs/drivers/m702"

	"golang.org/x/net/websocket"
)

//go:generate go-bindata-assetfs -prefix=root-fs/ ./root-fs

const (
	paramReady     = "0.08.15"
	paramRotation0 = "2.02.15"
	paramRotation1 = "2.02.16"
	paramRPMs      = "0.20.22"
	paramPosition  = "3.70.000"
	paramTemp0     = "0.07.004"
	paramTemp1     = "0.07.005"
	paramTemp2     = "0.07.006"
	paramTemp3     = "0.07.034"
)

var (
	codec    = binary.BigEndian
	addrFlag = flag.String("addr", "", "address:port to serve web-app")

	errMotorOffline     = fcsError{1, "fcs: motor OFFLINE"}
	errOpNotSupported   = fcsError{2, "fcs: operation not supported"}
	errUserAuth         = fcsError{100, "fcs: user not authenticated"}
	errUserPerm         = fcsError{101, "fcs: insufficient user permissions"}
	errInvalidReq       = fcsError{102, "fcs: invalid request"}
	errInvalidMotorName = fcsError{200, "fcs: invalid motor name"}
)

func newParameter(name string) m702.Parameter {
	p, err := m702.NewParameter(name)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func main() {
	flag.Parse()

	srv := newServer()
	mux := http.NewServeMux()
	mux.Handle("/", srv)
	mux.HandleFunc("/login", srv.handleLogin)
	mux.HandleFunc("/webcam", srv.handleWebcam)
	mux.Handle("/cmds", websocket.Handler(srv.cmdsHandler))
	mux.Handle("/data", websocket.Handler(srv.dataHandler))
	err := http.ListenAndServe(srv.Addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}

type registry struct {
	clients    map[*client]bool
	register   chan *client
	unregister chan *client
}

func newRegistry() registry {
	return registry{
		clients:    make(map[*client]bool),
		register:   make(chan *client),
		unregister: make(chan *client),
	}
}

type server struct {
	Addr string
	fs   http.Handler
	tmpl *template.Template

	session *authRegistry

	motor struct {
		x motor
		z motor
	}

	dataReg registry // clients interested in motor-statuses
	cmdsReg registry // clients interested in sending/receiving motor commands

	datac chan motorStatus
}

func newServer() *server {
	addr := *addrFlag
	if addr == "" {
		addr = getHostIP() + ":5555"
	}
	srv := &server{
		Addr:    addr,
		fs:      http.FileServer(http.Dir("./root-fs")),
		tmpl:    template.Must(template.New("fcs").Parse(string(MustAsset("index.html")))),
		session: newAuthRegistry(),
		dataReg: newRegistry(),
		cmdsReg: newRegistry(),
		datac:   make(chan motorStatus),
	}

	srv.motor.x = newMotor("x", "195.221.117.245:5021") // master-x
	srv.motor.z = newMotor("z", "195.221.117.245:5023") // master-z

	go srv.run()

	log.Printf("server created at %s...\n", srv.Addr)
	return srv
}

func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/" {
		client, _, err := srv.checkCredentials(w, r)
		if err != nil {
			http.Error(w, "credentials error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if !client.auth {
			srv.handleLogin(w, r)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "FCS_USER",
			Value: client.name,
		})
		srv.tmpl.Execute(w, srv)
		return
	}
	srv.fs.ServeHTTP(w, r)
}

func (srv *server) run() {
	go srv.monitor()
	for {
		select {
		case c := <-srv.dataReg.register:
			srv.dataReg.clients[c] = true

		case c := <-srv.dataReg.unregister:
			if _, ok := srv.dataReg.clients[c]; ok {
				delete(srv.dataReg.clients, c)
				close(c.datac)
				log.Printf(
					"client disconnected [%v]\n",
					c.ws.LocalAddr(),
				)
			}

		case c := <-srv.cmdsReg.register:
			srv.cmdsReg.clients[c] = true

		case c := <-srv.cmdsReg.unregister:
			if _, ok := srv.cmdsReg.clients[c]; ok {
				delete(srv.cmdsReg.clients, c)
				close(c.datac)
				log.Printf(
					"client disconnected [%v]\n",
					c.ws.LocalAddr(),
				)
			}

		case data := <-srv.datac:
			dataBuf := new(bytes.Buffer)
			err := json.NewEncoder(dataBuf).Encode(data)
			if err != nil {
				log.Printf("error marshalling data: %v\n", err)
				continue
			}
			for c := range srv.dataReg.clients {
				select {
				case c.datac <- dataBuf.Bytes():
				default:
					close(c.datac)
					delete(srv.dataReg.clients, c)
				}
			}
		}
	}
}

func (srv *server) motors() []*motor {
	return []*motor{
		&srv.motor.x,
		&srv.motor.z,
	}
}

func (srv *server) publishData() {
	for imotor, motor := range srv.motors() {
		// make sure the amount of memory used for the histos is under control
		switch {
		case len(motor.histos.rows) >= 128:
			for i, row := range motor.histos.rows {
				if i%2 == 0 {
					motor.histos.rows[i/2] = row
				}
			}
			motor.histos.rows = motor.histos.rows[:len(motor.histos.rows)/2]
		case len(motor.histos.rows) == 0:
			// no-op
		default:
			if time.Since(motor.histos.rows[0].id) >= 6*time.Hour {
				motor.histos.rows[0] = motor.histos.rows[len(motor.histos.rows)-1]
				motor.histos.rows = motor.histos.rows[:1]
			}
		}

		{
			c, err := net.DialTimeout("tcp", motor.addr, 2*time.Second)
			if err != nil || c == nil {
				motor.online = false
			} else {
				motor.online = true
			}
			if c != nil {
				c.Close()
			}
			if !motor.online {
				motor.histos.rows = append(motor.histos.rows, monData{id: time.Now()})
				plots := srv.makeMonPlots(imotor)
				srv.datac <- motorStatus{
					Motor:  motor.name,
					Online: false,
					Histos: plots,
					Webcam: srv.fetchWebcamImage(),
				}
				continue
			}
		}

		mm := m702.New(motor.addr)
		for _, p := range []*m702.Parameter{
			&motor.params.Ready,
			&motor.params.Rotation0,
			&motor.params.Rotation1,
			&motor.params.RPMs,
			&motor.params.Angle,
			&motor.params.Temps[0],
			&motor.params.Temps[1],
			&motor.params.Temps[2],
			&motor.params.Temps[3],
		} {
			err := mm.ReadParam(p)
			if err != nil {
				log.Printf("error reading %v (motor-%s) Pr-%v: %v\n", motor.addr, motor.name, *p, err)
			}
		}

		mon := monData{
			id:   time.Now(),
			rpms: codec.Uint32(motor.params.RPMs.Data[:]),
			temps: [4]float64{
				float64(codec.Uint32(motor.params.Temps[0].Data[:])),
				float64(codec.Uint32(motor.params.Temps[1].Data[:])),
				float64(codec.Uint32(motor.params.Temps[2].Data[:])),
				float64(codec.Uint32(motor.params.Temps[3].Data[:])),
			},
		}

		switch {
		case codec.Uint32(motor.params.Rotation0.Data[:]) == 1:
			mon.rotation = -1
		case codec.Uint32(motor.params.Rotation1.Data[:]) == 1:
			mon.rotation = +1
		}
		motor.histos.rows = append(motor.histos.rows, mon)
		plots := srv.makeMonPlots(imotor)

		status := motorStatus{
			Motor:    motor.name,
			Online:   motor.online,
			Ready:    codec.Uint32(motor.params.Ready.Data[:]) == 1,
			Rotation: mon.rotation,
			RPMs:     int(mon.rpms),
			Angle:    int(codec.Uint32(motor.params.Angle.Data[:])),
			Temps:    mon.temps,
			Histos:   plots,
			Webcam:   srv.fetchWebcamImage(),
		}

		srv.datac <- status
	}
}

func (srv *server) dataHandler(ws *websocket.Conn) {
	c := &client{
		srv:   srv,
		reg:   &srv.dataReg,
		datac: make(chan []byte, 256),
		ws:    ws,
	}
	c.reg.register <- c
	defer c.Release()

	c.run()
}

func (srv *server) cmdsHandler(ws *websocket.Conn) {
	c := &client{
		srv:   srv,
		reg:   &srv.cmdsReg,
		datac: make(chan []byte, 256),
		ws:    ws,
	}
	c.reg.register <- c
	defer c.Release()

	const maxRetries = 10

	acl := 0
cmdLoop:
	for {
		log.Printf("waiting for commands...\n")
		var req cmdRequest
		err := websocket.JSON.Receive(c.ws, &req)
		if err != nil {
			log.Printf("error rcv: %v\n", err)
			return
		}
		log.Printf(
			"received: {type=%q token=%q name=%q value=%v cmds=%q}\n",
			req.Type, req.Token, req.Name, req.Value, req.Cmds,
		)
		if acl == 0 {
			wc, ok := srv.session.get(req.Token)
			if wc.name == "" || !wc.auth || !ok {
				websocket.JSON.Send(c.ws, cmdReply{Err: errUserAuth.Error(), Req: req})
				continue
			}
			acl++
			c.setACL(wc.name)
		}
		if c.acl == 0 {
			websocket.JSON.Send(c.ws, cmdReply{Err: errUserPerm.Error(), Req: req})
			continue
		}

		nretries := 0
		var srvMotor *motor
		switch strings.ToLower(req.Motor) {
		case "x":
			srvMotor = &c.srv.motor.x
		case "z":
			srvMotor = &c.srv.motor.z
		default:
			websocket.JSON.Send(c.ws, cmdReply{Err: errInvalidMotorName.Error(), Req: req})
			continue
		}
		motor := m702.New(srvMotor.addr)
		script := newScripter(motor)

	retry:
		params := make([]m702.Parameter, 1)
		switch req.Name {
		case cmdReqReady:
			params[0] = newParameter(paramReady)
			codec.PutUint32(params[0].Data[:], uint32(req.Value))

		case cmdReqRotDir:
			params = make([]m702.Parameter, 2)
			switch int(req.Value) {
			case +1:
				params[0] = newParameter(paramRotation0)
				params[1] = newParameter(paramRotation1)
				codec.PutUint32(params[0].Data[:], 0)
				codec.PutUint32(params[1].Data[:], 1)

			case -1:
				params[0] = newParameter(paramRotation0)
				params[1] = newParameter(paramRotation1)
				codec.PutUint32(params[0].Data[:], 1)
				codec.PutUint32(params[1].Data[:], 0)

			case 0:
				params[0] = newParameter(paramRotation0)
				params[1] = newParameter(paramRotation1)
				codec.PutUint32(params[0].Data[:], 0)
				codec.PutUint32(params[1].Data[:], 0)
			}

		case cmdReqRPM:
			params[0] = newParameter(paramRPMs)
			codec.PutUint32(params[0].Data[:], uint32(req.Value))

		case cmdReqAnglePos:
			params[0] = newParameter(paramPosition)
			codec.PutUint32(params[0].Data[:], uint32(math.Floor(req.Value*10)))
			continue

		case cmdReqUploadCmds:
			r := bytes.NewReader([]byte(req.Cmds))
			err := script.run(motor, r)
			if err != nil {
				websocket.JSON.Send(c.ws, cmdReply{Err: err.Error(), Req: req})
			} else {
				websocket.JSON.Send(c.ws, cmdReply{Err: "", Req: req})
			}
			continue

		default:
			websocket.JSON.Send(c.ws, cmdReply{Err: errInvalidReq.Error(), Req: req})
			return
		}

		log.Printf("sending command %v to motor-%s %s...\n", params, srvMotor.name, srvMotor.addr)
		{
			conn, err := net.DialTimeout("tcp", srvMotor.addr, 1*time.Second)
			if err != nil || conn == nil {
				websocket.JSON.Send(c.ws, cmdReply{Err: errMotorOffline.Error(), Req: req})
				if conn != nil {
					conn.Close()
				}
				continue
			}
			conn.Close()
		}

		for _, p := range params {
			err = motor.WriteParam(p)
			if err != nil {
				log.Printf("error writing param Pr-%v: %v\n", p, err)
				if err == io.EOF && nretries < maxRetries {
					goto retry
				}
				websocket.JSON.Send(c.ws, cmdReply{Err: err.Error(), Req: req})
				goto cmdLoop
			}
		}
		websocket.JSON.Send(c.ws, cmdReply{Err: "", Req: req})
	}
}

type client struct {
	srv   *server
	reg   *registry
	ws    *websocket.Conn
	datac chan []byte
	acl   byte // acl notes whether the client is authentified and has r/w access
}

func (c *client) Release() {
	c.reg.unregister <- c
	c.ws.Close()
	c.reg = nil
	c.srv = nil
}

func (c *client) run() {
	//c.ws.SetReadLimit(maxMessageSize)
	//c.ws.SetReadDeadline(time.Now().Add(pongWait))
	//c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for data := range c.datac {
		err := websocket.Message.Send(c.ws, string(data))
		if err != nil {
			log.Printf(
				"error sending data to [%v]: %v\n",
				c.ws.LocalAddr(),
				err,
			)
			break
		}
	}
}

func (c *client) setACL(user string) {
	switch user {
	case "fcs":
		c.acl = 1
	case "visitor":
		c.acl = 0
	}
}

type cmdRequest struct {
	Motor string  `json:"motor"` // Motor is the motor name "x" | "z"
	Type  string  `json:"type"`
	Token string  `json:"token"` // Token is the web-client requestor
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Cmds  string  `json:"cmds"`
}

type cmdReply struct {
	Err string     `json:"err"`
	Req cmdRequest `json:"req"`
}

// list of all possible and known command-request names
const (
	cmdReqReady      = "ready"
	cmdReqRotDir     = "rotation-direction"
	cmdReqRPM        = "rpm"
	cmdReqAnglePos   = "angle-position"
	cmdReqUploadCmds = "upload-commands"
)

func getHostIP() string {
	host, err := os.Hostname()
	if err != nil {
		log.Fatalf("could not retrieve hostname: %v\n", err)
	}

	addrs, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("could not lookup hostname IP: %v\n", err)
	}

	for _, addr := range addrs {
		ipv4 := addr.To4()
		if ipv4 == nil {
			continue
		}
		return ipv4.String()
	}

	log.Fatalf("could not infer host IP")
	return ""
}

type fcsError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e fcsError) Error() string {
	return fmt.Sprintf("[%03d]: %s", e.Code, e.Msg)
}

func (e fcsError) String() string {
	return e.Error()
}
