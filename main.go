// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
	"github.com/go-lsst/ncs/drivers/m702"
	"golang.org/x/net/websocket"
)

//go:generate go-bindata-assetfs -prefix=root-fs/ ./root-fs

var (
	codec       = binary.BigEndian
	addrFlag    = flag.String("addr", "", "address:port to serve web-app")
	localFlag   = flag.Bool("local", false, "enable/disable local IPs")
	verboseFlag = flag.Bool("verbose", false, "enable/disable verbose mode")
	webcamFlag  = flag.Bool("webcam", true, "enable/disable webcam")
	mockFlag    = flag.Bool("mock", false, "enable/disable mock motors")

	errUserAuth = bench.FcsError{100, "fcs: user not authenticated"}
	errUserPerm = bench.FcsError{101, "fcs: insufficient user permissions"}
)

func dbgPrintf(format string, args ...interface{}) {
	if *verboseFlag {
		log.Printf(format, args...)
	}
}

func newParameter(name string) m702.Parameter {
	p, err := m702.NewParameter(name)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func main() {
	flag.Parse()

	srv := newServer(*addrFlag)
	mux := http.NewServeMux()
	mux.Handle("/", srv)
	mux.HandleFunc("/login", srv.handleLogin)
	mux.HandleFunc("/logout", srv.handleLogout)
	if *webcamFlag {
		mux.HandleFunc("/webcam", srv.handleWebcam)
	}
	mux.Handle("/cmds", websocket.Handler(srv.cmdsHandler))
	mux.Handle("/data", websocket.Handler(srv.dataHandler))
	mux.Handle("/video", websocket.Handler(srv.videoHandler))

	for _, v := range []struct {
		name string
		h    http.HandlerFunc
		acl  bool
	}{
		{
			name: "/api/mon",
			h:    srv.apiMonHandler,
			acl:  false,
		},
		{
			name: "/api/cmd/req-ready",
			h:    srv.apiCmdReqReadyHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-find-home",
			h:    srv.apiCmdReqFindHomeHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-pos",
			h:    srv.apiCmdReqPosHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-get-rpm",
			h:    srv.apiCmdReqGetRPMHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-rpm",
			h:    srv.apiCmdReqRPMHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-get-angle-pos",
			h:    srv.apiCmdReqGetAnglePosHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-angle-pos",
			h:    srv.apiCmdReqAnglePosHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-wait-pos",
			h:    srv.apiCmdReqAngleWaitHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-upload-cmds",
			h:    srv.apiCmdReqUploadCmdsHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-upload-script",
			h:    srv.apiCmdReqUploadScriptHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-reset",
			h:    srv.apiCmdReqResetHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-stop",
			h:    srv.apiCmdReqStopHandler,
			acl:  true,
		},
	} {
		mux.HandleFunc(v.name, srv.apiAuthenticated(v.h, v.acl))
	}

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

	webcam string // address:port of webcam CGI endpoint
	motor  struct {
		x motor
		z motor
	}
	bench bench.Bench

	dataReg  registry // clients interested in motor-statuses
	cmdsReg  registry // clients interested in sending/receiving motor commands
	videoReg registry // clients interested in receiving webcam data

	datac chan bench.MotorInfos
}

func newServer(addr string) *server {
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
		datac:   make(chan bench.MotorInfos),
	}

	switch {
	case *mockFlag:
		srv.motor.x = newMotorMock("x", "127.0.0.1:5020")           // master-x
		slave := newMotorSlave("x", "127.0.0.1:5022", &srv.motor.x) // slave-x
		srv.motor.x.slave = &slave
		srv.motor.z = newMotorMock("z", "127.0.0.1:5021") // master-z
	case *localFlag:
		srv.motor.x = newMotor("x", "192.168.0.21:502")               // master-x
		slave := newMotorSlave("x", "192.168.0.22:502", &srv.motor.x) // slave-x
		srv.motor.x.slave = &slave
		srv.motor.z = newMotor("z", "192.168.0.23:502") // master-z
		if *webcamFlag {
			srv.webcam = "192.168.0.30:80"
		}
	default:
		ip := "134.158.155.16"
		srv.motor.x = newMotor("x", ip+":5021")               // master-x
		slave := newMotorSlave("x", ip+":5022", &srv.motor.x) // slave-x
		srv.motor.x.slave = &slave
		srv.motor.z = newMotor("z", ip+":5023") // master-z
		if *webcamFlag {
			srv.webcam = ip + ":80"
		}
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
	go srv.monitor()
	for {
		dbgPrintf("run-tick\n")
		select {
		case c := <-srv.videoReg.register:
			dbgPrintf("video-register")
			srv.videoReg.clients[c] = true

		case c := <-srv.videoReg.unregister:
			dbgPrintf("video-unregister")
			if _, ok := srv.videoReg.clients[c]; ok {
				delete(srv.videoReg.clients, c)
				close(c.datac)
				log.Printf(
					"client disconnected [%v]\n",
					c.ws.LocalAddr(),
				)
			}

		case c := <-srv.dataReg.register:
			dbgPrintf("data-register")
			srv.dataReg.clients[c] = true

		case c := <-srv.dataReg.unregister:
			dbgPrintf("data-unregister")
			if _, ok := srv.dataReg.clients[c]; ok {
				delete(srv.dataReg.clients, c)
				close(c.datac)
				log.Printf(
					"client disconnected [%v]\n",
					c.ws.LocalAddr(),
				)
			}

		case c := <-srv.cmdsReg.register:
			dbgPrintf("cmds-register")
			srv.cmdsReg.clients[c] = true

		case c := <-srv.cmdsReg.unregister:
			dbgPrintf("cmds-unregister")
			if _, ok := srv.cmdsReg.clients[c]; ok {
				delete(srv.cmdsReg.clients, c)
				close(c.datac)
				log.Printf(
					"client disconnected [%v]\n",
					c.ws.LocalAddr(),
				)
			}

		case data := <-srv.datac:
			dbgPrintf("data on srv.datac\n")
			if len(srv.dataReg.clients) == 0 {
				dbgPrintf("no client connected")
				continue
			}
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
		dbgPrintf("run-untick")
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
		dbgPrintf("-- motor-%v (%s)...\n", motor.name, motor.addr)
		// make sure the amount of memory used for the histos is under control
		switch {
		case len(motor.histos.rows) >= 128:
			n := len(motor.histos.rows)
			copy(motor.histos.rows[:n/2], motor.histos.rows[n/2:n])
			motor.histos.rows = motor.histos.rows[:n/2]
		case len(motor.histos.rows) == 0:
			// no-op
		default:
			if time.Since(motor.histos.rows[0].id) >= 6*time.Hour {
				motor.histos.rows[0] = motor.histos.rows[len(motor.histos.rows)-1]
				motor.histos.rows = motor.histos.rows[:1]
			}
		}

		// FIXME(sbinet): use motor.infos() instead.

		{
			var err error
			motor.online, err = motor.isOnline(motorTimeout)
			if err != nil {
				log.Printf("-- motor-%v: offline (err=%v)\n", motor.name, err)
			}
			if !motor.online {
				motor.histos.rows = append(motor.histos.rows, monData{id: time.Now()})
				plots := srv.makeMonPlots(imotor)
				srv.datac <- bench.MotorInfos{
					Motor:  motor.name,
					Online: false,
					Status: "N/A",
					FSM:    "N/A",
					Sync:   false,
					Mode:   "N/A",
					Histos: plots,
					Webcam: srv.fetchWebcamImage(),
				}
				dbgPrintf("-- motor-%v: continue\n", motor.name)
				continue
			}
		}

		errs := motor.poll()
		if len(errs) > 0 {
			for _, err := range errs {
				log.Printf("%v", err)
			}
		}

		if motor.isManual() {
			// make sure we won't override what manual-mode did
			// when we go back to sw-mode/ready-mode
			err := motor.updateAnglePos()
			if err != nil {
				log.Printf("-- motor-%v: standby: %v\n", motor.name, err)
			}
		}

		mon := monData{
			id:    time.Now(),
			rpms:  motor.rpms(),
			angle: motor.angle(),
			temps: [4]float64{
				float64(codec.Uint32(motor.params.Temps[0].Data[:])),
				float64(codec.Uint32(motor.params.Temps[1].Data[:])),
				float64(codec.Uint32(motor.params.Temps[2].Data[:])),
				float64(codec.Uint32(motor.params.Temps[3].Data[:])),
			},
		}

		status := "N/A"

		manual := motor.isManual()
		ready := !manual
		hwsafetyON := motor.isHWLocked()
		fsm := motor.fsm()

		switch {
		case hwsafetyON:
			status = "h/w safety"
		case manual:
			status = "manual"
		case ready:
			status = "ready"
		}

		if motor.online {
			switch {
			case codec.Uint32(motor.params.Home.Data[:]) == 1:
				mon.mode = motorModeHome
			case codec.Uint32(motor.params.ModePos.Data[:]) == 1:
				mon.mode = motorModePos
			}
		}
		motor.histos.rows = append(motor.histos.rows, mon)
		plots := srv.makeMonPlots(imotor)

		dbgPrintf("-- %s: online=%v status=%v mode=%v\n", motor.name, motor.online, status, mon.Mode())
		infos := bench.MotorInfos{
			Motor:  motor.name,
			Online: motor.online,
			Status: status,
			FSM:    fsm,
			Sync:   motor.isSyncOK(),
			Mode:   mon.Mode(),
			RPMs:   int(mon.rpms),
			Angle:  mon.angle,
			Temps:  mon.temps,
			Histos: plots,
			Webcam: srv.fetchWebcamImage(),
		}

		srv.datac <- infos
		dbgPrintf("-- %s: done\n", motor.name)
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

func (srv *server) videoHandler(ws *websocket.Conn) {
	c := &client{
		srv:   srv,
		reg:   &srv.videoReg,
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
		req.tstamp = time.Now().UTC()
		log.Printf(
			"received: {type=%q token=%q name=%q value=%v cmds=%q}\n",
			req.Type, req.Token, req.Name, req.Value, req.Cmds,
		)
		if acl == 0 {
			usr, ok := srv.session.get(req.Token)
			if usr.name == "" || !usr.auth || !ok {
				srv.sendReply(c.ws, cmdReply{Err: errUserAuth.Error(), Req: req})
				continue
			}
			acl++
			c.setACL(usr.name)
		}
		if c.acl == 0 {
			srv.sendReply(c.ws, cmdReply{Err: errUserPerm.Error(), Req: req})
			continue
		}

		nretries := 0
		var srvMotor *motor
		switch strings.ToLower(req.Motor) {
		case "x":
			srvMotor = &c.srv.motor.x
		case "z":
			srvMotor = &c.srv.motor.z
		case "":
			srv.sendReply(c.ws, cmdReply{Err: bench.ErrInvalidMotorName.Error(), Req: req})
			continue
		default:
			srv.sendReply(c.ws, cmdReply{Err: bench.ErrInvalidMotorName.Error(), Req: req})
			continue
		}
		{
			online, err := srvMotor.isOnline(motorTimeout)
			if err != nil || !online {
				srv.sendReply(c.ws, cmdReply{Err: bench.ErrMotorOffline.Error(), Req: req})
				continue
			}
		}

		if srvMotor.isHWLocked() {
			srv.sendReply(c.ws, cmdReply{Err: bench.ErrMotorHWLock.Error(), Req: req})
			continue
		}

		if srvMotor.isManual() {
			srv.sendReply(c.ws, cmdReply{Err: bench.ErrMotorManual.Error(), Req: req})
			continue
		}

		var motor bench.Motor = srvMotor.Motor()
		script := newScripter(srv, motor)

	retry:
		params := make([]m702.Parameter, 1)
		switch req.Name {
		case cmdReqStop:
			dbgPrintf("cmd-req-stop:\n")
			err := srvMotor.stop()
			reply := cmdReply{Req: req}
			if err != nil {
				reply.Err = err.Error()
			}
			srv.sendReply(c.ws, reply)

		case cmdReqReset:
			dbgPrintf("cmd-req-reset:\n")
			err := srvMotor.reset()
			reply := cmdReply{Req: req}
			if err != nil {
				reply.Err = err.Error()
			}
			srv.sendReply(c.ws, reply)
			continue

		case cmdReqReady:
			dbgPrintf("cmd-req-ready: %v\n", uint32(req.Value))
			params[0] = newParameter(bench.ParamCmdReady)
			codec.PutUint32(params[0].Data[:], uint32(req.Value))

		case cmdReqFindHome:
			dbgPrintf("cmd-req-find-home\n")
			err := srvMotor.findHome()
			reply := cmdReply{Req: req}
			if err != nil {
				reply.Err = err.Error()
			}
			srv.sendReply(c.ws, reply)
			continue

		case cmdReqPos:
			dbgPrintf("cmd-req-pos\n")
			params = append([]m702.Parameter{},
				newParameter(bench.ParamCmdReady),
				newParameter(bench.ParamModePos),
				newParameter(bench.ParamHome),
				newParameter(bench.ParamCmdReady),
			)

			codec.PutUint32(params[0].Data[:], 0)
			codec.PutUint32(params[1].Data[:], 1)
			codec.PutUint32(params[2].Data[:], 0)
			codec.PutUint32(params[3].Data[:], 1)

		case cmdReqRPM:
			dbgPrintf("cmd-req-rpm\n")
			params[0] = newParameter(bench.ParamRPMs)
			codec.PutUint32(params[0].Data[:], uint32(req.Value))

		case cmdReqAnglePos:
			dbgPrintf("cmd-req-angle-pos\n")
			params[0] = newParameter(bench.ParamWritePos)
			codec.PutUint32(params[0].Data[:], uint32(math.Floor(req.Value*10)))

		case cmdReqUploadCmds:
			dbgPrintf("cmd-req-upload-cmds\n")
			r := bytes.NewReader([]byte(req.Cmds))
			err := script.run(motor, r, ioutil.Discard) // FIXME(sbinet): redirect to expert-textarea
			reply := cmdReply{Req: req}
			if err != nil {
				reply.Err = err.Error()
			}
			srv.sendReply(c.ws, reply)
			continue

		default:
			srv.sendReply(c.ws, cmdReply{Err: bench.ErrInvalidReq.Error(), Req: req})
			continue
		}

		log.Printf("sending command %v to motor-%s %s...\n", params, srvMotor.name, srvMotor.addr)
		for _, p := range params {
			err = motor.WriteParam(p)
			if err != nil {
				log.Printf("error writing param Pr-%v: %v\n", p, err)
				if err == io.EOF && nretries < maxRetries {
					goto retry
				}
				srv.sendReply(c.ws, cmdReply{Err: err.Error(), Req: req})
				goto cmdLoop
			}
		}
		srv.sendReply(c.ws, cmdReply{Err: "", Req: req})
	}
}

func (srv *server) sendReply(ws *websocket.Conn, reply cmdReply) {
	log.Printf("reply: {err=%q, req={token=%q, type=%q, motor=%q, name=%q}}\n",
		reply.Err, reply.Req.Token, reply.Req.Type, reply.Req.Motor, reply.Req.Name,
	)
	websocket.JSON.Send(ws, reply)
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
	if *mockFlag {
		if user == "faux-fcs" {
			c.acl = 1
		}
	}
}

type cmdRequest struct {
	Motor  string  `json:"motor"` // Motor is the motor name "x" | "z"
	Type   string  `json:"type"`
	Token  string  `json:"token"` // Token is the web-client requestor
	Name   string  `json:"name"`
	Value  float64 `json:"value"`
	Cmds   string  `json:"cmds"`
	tstamp time.Time
}

type cmdReply struct {
	Err string     `json:"err"`
	Req cmdRequest `json:"req"`
}

// list of all possible and known command-request names
const (
	cmdReqStop       = "stop"
	cmdReqReset      = "reset"
	cmdReqFindHome   = "find-home"
	cmdReqPos        = "pos"
	cmdReqReady      = "ready"
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
