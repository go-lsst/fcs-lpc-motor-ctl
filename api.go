// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errInvalidHTTPMethod = errors.New("invalid HTTP method")
)

func (srv *server) apiMonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		srv.apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "/api/mon: %#v", r)
}

func (srv *server) apiCmdReqReadyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api/cmd/req-ready: %#v", r)
}

func (srv *server) apiCmdReqFindHomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api/cmd/req-find-home: %#v", r)
}

func (srv *server) apiCmdReqPosHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api/cmd/req-pos: %#v", r)
}

func (srv *server) apiCmdReqRPMHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api/cmd/req-rpm: %#v", r)
}

func (srv *server) apiCmdReqAnglePosHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api/cmd/req-angle-pos: %#v", r)
}

func (srv *server) apiCmdReqUploadCmdsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api/cmd/req-upload-cmds: %#v", r)
}

func (srv *server) apiError(w http.ResponseWriter, err error, code int) {
	http.Error(w, fmt.Sprintf("{error:%q, code:%v}", err.Error(), code), code)
}

func (srv *server) apiAuthenticated(h func(w http.ResponseWriter, r *http.Request), needACL bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cli, cookie, err := srv.checkCredentials(w, r)
		if err != nil {
			srv.apiError(w, err, http.StatusInternalServerError)
			return
		}

		err = r.ParseForm()
		if err != nil {
			srv.apiError(w, fmt.Errorf("could not parse form: %v", err), http.StatusInternalServerError)
			return
		}

		user := r.FormValue("username")
		pass := r.FormValue("password")
		if user == "" {
			u, p, ok := r.BasicAuth()
			if ok {
				user = u
				pass = p
			}
		}
		if !srv.authenticate(user, pass) {
			srv.apiError(w, fmt.Errorf("invalid user/password"), http.StatusForbidden)
			return
		}

		cli.auth = true
		cli.name = user
		srv.session.set(cookie, cli)

		r.SetBasicAuth(user, pass)
		if !cli.auth {
			srv.apiError(w, fmt.Errorf("authentication error"), http.StatusForbidden)
			return
		}

		c := client{}
		c.setACL(cli.name)

		if needACL && c.acl != 1 {
			srv.apiError(w, fmt.Errorf("authorization error"), http.StatusForbidden)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "FCS_USER",
			Value: cli.name,
		})

		w.Header().Set("Content-Type", "application/json")
		h(w, r)
		return
	}
}
