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
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
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
}
func apiCmdReqFindHomeHandler(w http.ResponseWriter, r *http.Request) {
}

func apiCmdReqPosHandler(w http.ResponseWriter, r *http.Request) {
}

func apiCmdReqRPMHandler(w http.ResponseWriter, r *http.Request) {
}

func apiCmdReqAnglePosHandler(w http.ResponseWriter, r *http.Request) {
}

func apiCmdReqUploadCmdsHandler(w http.ResponseWriter, r *http.Request) {
}

func apiCmdReqUploadScriptHandler(w http.ResponseWriter, r *http.Request) {
}

func apiError(w http.ResponseWriter, err error, code int) {
	http.Error(w, fmt.Sprintf(`{"error":%q, "code":%v}`, err.Error(), code), code)
}

func apiOK(w http.ResponseWriter, code int) {
	http.Error(w, fmt.Sprintf(`{"code":%v}`, code), code)
}
