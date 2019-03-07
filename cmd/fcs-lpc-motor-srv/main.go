// Copyright ©2019 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command fcs-lpc-motor-srv is a simple REST server.
package main

import (
	"encoding/binary"
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
)

const (
	motorTimeout = 5 * time.Second
)

var (
	codec = binary.BigEndian
)

func main() {

	var (
		addr = flag.String("addr", "134.158.155.17:4444", "address:port to serve web-app")
		// verbose = flag.Bool("verbose", false, "enable/disable verbose mode")
		// mock = flag.Bool("mock", false, "enable/disable mock motors")

	)

	flag.Parse()

	srv := newServer()

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
			name: "/api/cmd/req-rpm",
			h:    srv.apiCmdReqRPMHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-angle-pos",
			h:    srv.apiCmdReqAnglePosHandler,
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
	} {
		http.HandleFunc(v.name, srv.apiAuthenticated(v.h, v.acl))
	}

	for i := range srv.motors {
		m := &srv.motors[i]
		errs := m.poll()
		if len(errs) > 0 {
			log.Fatalf("could not poll motor %q: %v", m.name, errs[0])
		}
	}

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func (srv *server) getMotor(name string) (*motor, error) {
	var m *motor
	switch strings.ToLower(name) {
	case "x":
		m = &srv.motors[0]
	case "z":
		// FIXME(sbinet): motor-z isn't available right now.
		return nil, bench.ErrInvalidMotorName
		m = &srv.motors[1]
	default:
		return nil, bench.ErrInvalidMotorName
	}
	return m, nil
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
	cmdReqFindHome   = "find-home"
	cmdReqPos        = "pos"
	cmdReqReady      = "ready"
	cmdReqRPM        = "rpm"
	cmdReqAnglePos   = "angle-position"
	cmdReqUploadCmds = "upload-commands"
)
