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
	"net"
	"net/http"
	"os"
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
	paramTemp0     = "0.07.004"
	paramTemp1     = "0.07.005"
	paramTemp2     = "0.07.006"
	paramTemp3     = "0.07.034"
)

var (
	codec    = binary.BigEndian
	addrFlag = flag.String("addr", "", "address:port to serve web-app")

	errMotorOffline   = fcsError{1, "fcs: motor OFFLINE"}
	errOpNotSupported = fcsError{2, "fcs: operation not supported"}
	errUserAuth       = fcsError{100, "fcs: user not authenticated"}
	errUserPerm       = fcsError{101, "fcs: insufficient user permissions"}
	errInvalidReq     = fcsError{102, "fcs: invalid request"}
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
	Motors []string
	Addr   string
	fs     http.Handler
	tmpl   *template.Template

	session *authRegistry

	params struct {
		Ready     m702.Parameter
		Rotation0 m702.Parameter
		Rotation1 m702.Parameter
		RPMs      m702.Parameter
		Angle     m702.Parameter
		Temps     [4]m702.Parameter
	}
	online bool // whether motors are online/connected

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
		Motors: []string{
			"134.158.125.223:502",
			"134.158.125.224:502",
		},
		Addr:    addr,
		fs:      http.FileServer(http.Dir("./root-fs")),
		tmpl:    template.Must(template.New("fcs").Parse(string(MustAsset("index.html")))),
		session: newAuthRegistry(),
		dataReg: newRegistry(),
		cmdsReg: newRegistry(),
		datac:   make(chan motorStatus),
	}

	srv.params.Ready = newParameter(paramReady)
	srv.params.Rotation0 = newParameter(paramRotation0)
	srv.params.Rotation1 = newParameter(paramRotation1)
	srv.params.RPMs = newParameter(paramRPMs)
	// srv.params.Angle = newParameter("") // FIXME(sbinet)
	srv.params.Temps = [4]m702.Parameter{
		newParameter(paramTemp0),
		newParameter(paramTemp1),
		newParameter(paramTemp2),
		newParameter(paramTemp3),
	}

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
	go func() {
		tick := time.NewTicker(5 * time.Second)
		srv.publishData()
		for range tick.C {
			srv.publishData()
		}
	}()

	dataBuf := new(bytes.Buffer)
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
			dataBuf.Reset()
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

func (srv *server) publishData() {
	master := srv.Motors[0]
	{
		c, err := net.DialTimeout("tcp", master, 2*time.Second)
		if err != nil || c == nil {
			srv.online = false
		}
		if c != nil {
			c.Close()
		}
		if !srv.online {
			srv.datac <- motorStatus{Online: false}
			return
		}
	}

	motor := m702.New(master)
	for _, p := range []*m702.Parameter{
		&srv.params.Ready,
		&srv.params.Rotation0,
		&srv.params.Rotation1,
		&srv.params.RPMs,
		// &srv.params.Angle,
		&srv.params.Temps[0],
		&srv.params.Temps[1],
		&srv.params.Temps[2],
		&srv.params.Temps[3],
	} {
		err := motor.ReadParam(p)
		if err != nil {
			log.Printf("error reading %v Pr-%v: %v\n", master, *p, err)
		}
	}

	status := motorStatus{
		Online:   srv.online,
		Ready:    codec.Uint32(srv.params.Ready.Data[:]) == 1,
		Rotation: 0,
		RPMs:     int(codec.Uint32(srv.params.RPMs.Data[:])),
		// Angle: int(codec.Uint32(srv.params.Angle.Data[:])),
		Temps: [4]float64{
			float64(codec.Uint32(srv.params.Temps[0].Data[:])),
			float64(codec.Uint32(srv.params.Temps[1].Data[:])),
			float64(codec.Uint32(srv.params.Temps[2].Data[:])),
			float64(codec.Uint32(srv.params.Temps[3].Data[:])),
		},
	}
	switch {
	case codec.Uint32(srv.params.Rotation0.Data[:]) == 1:
		status.Rotation = -1
	case codec.Uint32(srv.params.Rotation1.Data[:]) == 1:
		status.Rotation = +1
	}

	srv.datac <- status
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
	motor := m702.New(c.srv.Motors[0])
	script := newScripter(motor)

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
			websocket.JSON.Send(c.ws, cmdReply{Err: errOpNotSupported.Error(), Req: req})
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

		log.Printf("sending command %v to motor...\n", params)
		{
			conn, err := net.DialTimeout("tcp", c.srv.Motors[0], 1*time.Second)
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

type motorStatus struct {
	Online   bool       `json:"online"`
	Ready    bool       `json:"ready"`
	Rotation int        `json:"rotation_direction"`
	RPMs     int        `json:"rpms"`
	Angle    int        `json:"angle"`
	Temps    [4]float64 `json:"temps"`
}

type cmdRequest struct {
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
