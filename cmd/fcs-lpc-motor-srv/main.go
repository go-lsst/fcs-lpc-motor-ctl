// Copyright ©2019 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command fcs-lpc-motor-srv is a simple REST server.
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
	"github.com/go-lsst/ncs/drivers/m702"
)

const (
	motorTimeout = 5 * time.Second
)

var (
	errInvalidHTTPMethod = errors.New("invalid HTTP method")
)

var (
	codec      = binary.BigEndian
	xmotorAddr = "134.158.155.16:5021"
	zmotorAddr = "134.158.155.16:5023"
	motors     = []motor{
		newMotor("x", xmotorAddr),
		//		newMotor("z", zmotorAddr),
	}
)

func main() {

	var (
		addr = flag.String("addr", "134.158.155.17:4444", "address:port to serve web-app")
		// verbose = flag.Bool("verbose", false, "enable/disable verbose mode")
		// mock = flag.Bool("mock", false, "enable/disable mock motors")

	)

	flag.Parse()

	for _, v := range []struct {
		name string
		h    http.HandlerFunc
		acl  bool
	}{
		{
			name: "/api/mon",
			h:    apiMonHandler,
			acl:  false,
		},
		{
			name: "/api/cmd/req-ready",
			h:    apiCmdReqReadyHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-find-home",
			h:    apiCmdReqFindHomeHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-pos",
			h:    apiCmdReqPosHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-rpm",
			h:    apiCmdReqRPMHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-angle-pos",
			h:    apiCmdReqAnglePosHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-upload-cmds",
			h:    apiCmdReqUploadCmdsHandler,
			acl:  true,
		},
		{
			name: "/api/cmd/req-upload-script",
			h:    apiCmdReqUploadScriptHandler,
			acl:  true,
		},
	} {
		// http.HandleFunc(v.name, srv.apiAuthenticated(v.h, v.acl))
		http.HandleFunc(v.name, v.h)
	}

	for i := range motors {
		m := &motors[i]
		errs := m.poll()
		if len(errs) > 0 {
			log.Fatalf("could not poll motor %q: %v", m.name, errs[0])
		}
	}

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func apiMonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}

	var infos [2]bench.MotorInfos
	for i, m := range motors {
		info, err := m.infos(motorTimeout)
		if err != nil {
			apiError(w, err, http.StatusServiceUnavailable)
			return
		}
		infos[i] = info
	}
	var resp = struct {
		Err   string              `json:"error,omitempty"`
		Code  int                 `json:"code"`
		Infos [2]bench.MotorInfos `json:"infos"`
	}{
		Err:   "",
		Code:  http.StatusOK,
		Infos: infos,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(resp)
	if err != nil {
		apiError(w, fmt.Errorf("could not encode monitoring infos to JSON: %v", err), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(w, &buf)
	if err != nil {
		apiError(w, fmt.Errorf("error writing JSON response: %v", err), http.StatusInternalServerError)
		return
	}
}

func apiCmdReqReadyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	m, ok := apiCheck(req, w, r)
	if !ok {
		return
	}

	p := newParameter(bench.ParamCmdReady)
	codec.PutUint32(p.Data[:], uint32(req.Value))
	err = m.Motor().WriteParam(p)
	if err != nil {
		apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
		return
	}
	apiOK(w, http.StatusOK)
}

func apiCmdReqFindHomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	m, ok := apiCheck(req, w, r)
	if !ok {
		return
	}

	params := append([]m702.Parameter{},
		newParameter(bench.ParamCmdReady),
		newParameter(bench.ParamModePos),
		newParameter(bench.ParamHome),
		newParameter(bench.ParamCmdReady),
	)

	codec.PutUint32(params[0].Data[:], 0)
	codec.PutUint32(params[1].Data[:], 0)
	codec.PutUint32(params[2].Data[:], 1)
	codec.PutUint32(params[3].Data[:], 1)

	for _, p := range params {
		err = m.Motor().WriteParam(p)
		if err != nil {
			apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
			return
		}
	}

	apiOK(w, http.StatusOK)
}

func apiCmdReqPosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	m, ok := apiCheck(req, w, r)
	if !ok {
		return
	}

	params := append([]m702.Parameter{},
		newParameter(bench.ParamCmdReady),
		newParameter(bench.ParamModePos),
		newParameter(bench.ParamHome),
		newParameter(bench.ParamCmdReady),
	)

	codec.PutUint32(params[0].Data[:], 0)
	codec.PutUint32(params[1].Data[:], 1)
	codec.PutUint32(params[2].Data[:], 0)
	codec.PutUint32(params[3].Data[:], 1)

	for _, p := range params {
		err = m.Motor().WriteParam(p)
		if err != nil {
			apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
			return
		}
	}

	apiOK(w, http.StatusOK)
}

func apiCmdReqRPMHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	if req.Value > 3000 {
		apiError(w, fmt.Errorf("invalid RPM value (%v > 3000)", req.Value), http.StatusBadRequest)
		return
	}
	if req.Value < 0 {
		apiError(w, fmt.Errorf("invalid RPM value (%v < 0)", req.Value), http.StatusBadRequest)
		return
	}

	m, ok := apiCheck(req, w, r)
	if !ok {
		return
	}

	p := newParameter(bench.ParamRPMs)
	codec.PutUint32(p.Data[:], uint32(req.Value))
	err = m.Motor().WriteParam(p)
	if err != nil {
		apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
		return
	}

	apiOK(w, http.StatusOK)
}

func apiCmdReqAnglePosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"
	if req.Value > +90 {
		apiError(w, fmt.Errorf("invalid angle position (%v > +90.0)", req.Value), http.StatusBadRequest)
		return
	}
	if req.Value < -90 {
		apiError(w, fmt.Errorf("invalid angle position (%v < -90.0)", req.Value), http.StatusBadRequest)
		return
	}
	req.Value *= 10

	m, ok := apiCheck(req, w, r)
	if !ok {
		return
	}

	p := newParameter(bench.ParamWritePos)
	codec.PutUint32(p.Data[:], uint32(req.Value))
	err = m.Motor().WriteParam(p)
	if err != nil {
		apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
		return
	}

	apiOK(w, http.StatusOK)
}

func apiCmdReqUploadCmdsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	m, ok := apiCheck(req, w, r)
	if !ok {
		return
	}

	buf := new(bytes.Buffer)
	script := newScripter(m.Motor())
	cmds := bytes.NewReader([]byte(req.Cmds))
	err = script.run(m.Motor(), cmds, buf)
	if err != nil {
		apiError(w, fmt.Errorf("error running script: %v", err), http.StatusInternalServerError)
		return
	}

	var reply = struct {
		Err    string `json:"error,omitempty"`
		Code   int    `json:"code"`
		Script string `json:"script"`
	}{
		Err:    "",
		Code:   http.StatusOK,
		Script: string(buf.Bytes()),
	}

	var o = new(bytes.Buffer)
	err = json.NewEncoder(o).Encode(reply)
	if err != nil {
		apiError(w, fmt.Errorf("error encoding JSON reply: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, o)
	if err != nil {
		apiError(w, fmt.Errorf("error sending JSON reply: %v", err), http.StatusInternalServerError)
		return
	}

	apiOK(w, http.StatusOK)
}

func apiCmdReqUploadScriptHandler(w http.ResponseWriter, r *http.Request) {
}

func apiError(w http.ResponseWriter, err error, code int) {
	http.Error(w, fmt.Sprintf(`{"error":%q, "code":%v}`, err.Error(), code), code)
}

func apiOK(w http.ResponseWriter, code int) {
	http.Error(w, fmt.Sprintf(`{"code":%v}`, code), code)
}

func apiCheck(req cmdRequest, w http.ResponseWriter, r *http.Request) (*motor, bool) {
	m, err := getMotor(req.Motor)
	if err != nil {
		apiError(w, err, http.StatusBadRequest)
		return nil, false
	}

	if online, err := m.isOnline(motorTimeout); err != nil || !online {
		if err != nil {
			apiError(w, err, http.StatusServiceUnavailable)
			return nil, false
		}
		apiError(w, bench.ErrMotorOffline, http.StatusServiceUnavailable)
		return nil, false
	}

	if m.isHWLocked() {
		apiError(w, bench.ErrMotorHWLock, http.StatusServiceUnavailable)
		return nil, false
	}

	if m.isManual() {
		apiError(w, bench.ErrMotorManual, http.StatusServiceUnavailable)
		return nil, false
	}

	return m, true
}

func getMotor(name string) (*motor, error) {
	var m *motor
	switch strings.ToLower(name) {
	case "x":
		m = &motors[0]
	case "z":
		// FIXME(sbinet): motor-z isn't available right now.
		return nil, bench.ErrInvalidMotorName
		m = &motors[1]
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
