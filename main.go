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
	//mux.Handle("/", http.FileServer(http.Dir("./root-fs")))
	mux.Handle("/", srv)
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
	params struct {
		Ready     m702.Parameter
		Rotation0 m702.Parameter
		Rotation1 m702.Parameter
		RPMs      m702.Parameter
		Angle     m702.Parameter
		Temps     [4]m702.Parameter
	}

	dataReg registry // clients interested in motor-statuses
	cmdsReg registry // clients interested in sending/receiving motor commands

	datac chan motorStatus
	cmdsc chan command
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
	cmdsBuf := new(bytes.Buffer)
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

		case cmd := <-srv.cmdsc:
			log.Printf("received command: %v\n", cmd)
			cmdsBuf.Reset()
			err := json.NewEncoder(cmdsBuf).Encode(cmd)
			if err != nil {
				log.Printf("error marshalling data: %v\n", err)
				continue
			}
			for c := range srv.cmdsReg.clients {
				select {
				case c.datac <- cmdsBuf.Bytes():
				default:
					close(c.datac)
					delete(srv.cmdsReg.clients, c)
				}
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

	motor := m702.New(c.srv.Motors[0])
	for {
		log.Printf("waiting for commands...\n")
		var req cmdRequest
		err := websocket.JSON.Receive(c.ws, &req)
		if err != nil {
			log.Printf("error rcv: %v\n", err)
			return
		}
		log.Printf("received: %v\n", req)
		nretries := 0
	retry:
		params := make([]m702.Parameter, 1)
		switch req.Name {
		case "ready":
			params[0] = newParameter(paramReady)
			codec.PutUint32(params[0].Data[:], uint32(req.Value))

		case "rotation-direction":
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

		case "rpm":
			params[0] = newParameter(paramRPMs)
			codec.PutUint32(params[0].Data[:], uint32(req.Value))

		case "angle-position":
			websocket.JSON.Send(c.ws, cmdReply{Err: fmt.Errorf("not supported"), Req: req})
			continue

		default:
			log.Printf("invalid request: %#v\n", req)
			return
		}

		log.Printf("sending command %v to motor...\n", params)
		for _, p := range params {
			err = motor.WriteParam(p)
			if err != nil {
				log.Printf("error writing param Pr-%v: %v\n", p, err)
				if err == io.EOF && nretries < maxRetries {
					goto retry
				}
				websocket.JSON.Send(c.ws, cmdReply{Err: err, Req: req})
				return
			}
		}
		websocket.JSON.Send(c.ws, cmdReply{Err: nil, Req: req})
	}
}

type client struct {
	srv   *server
	reg   *registry
	ws    *websocket.Conn
	datac chan []byte
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

type motorStatus struct {
	Ready    bool       `json:"ready"`
	Rotation int        `json:"rotation_direction"`
	RPMs     int        `json:"rpms"`
	Angle    int        `json:"angle"`
	Temps    [4]float64 `json:"temps"`
}

type cmdRequest struct {
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type cmdReply struct {
	Err error      `json:"err"`
	Req cmdRequest `json:"req"`
}

type command struct {
}

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
