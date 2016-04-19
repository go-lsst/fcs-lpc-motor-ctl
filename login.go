// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/satori/go.uuid"
)

const loginPage = `
<!DOCTYPE html>
<!-- Copyright 2016 The go-lsst Authors. All rights reserved.
  -- Use of this source code is governed by a BSD-style
  -- license that can be found in the LICENSE file.
  -->
<html>
  <head>
    <meta name="viewport" content="width=device-width, minimum-scale=1.0, initial-scale=1.0, user-scalable=yes">
    <meta charset="utf-8">
    <title>FCS LPC Testbench</title>
    <script src="bower_components/webcomponentsjs/webcomponents-lite.min.js"></script>
    <link rel="import" href="bower_components/paper-styles/paper-styles-classes.html">
    <link rel="import" href="fcs-lpc-motor.html">
    <style>
      html {
        overflow-y: auto;
      }
      body {
        font-family: 'Roboto', 'Helvetica Neue', Helvetica, Arial, sans-serif;
        font-weight: 300;
      }
    </style>
  </head>
  <script>
    </script>
	<body unresolved class="layout vertical center-center">

		<div id="fcs-app">
			<fcs-lpc-motor-login></fcs-lpc-motor-login>
			%s
		</div>

	</body>
</html>
`

func (srv *server) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, loginPage, "")
		return
	}
	client, cookie, err := srv.checkCredentials(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, loginPage, "<h3>ERROR: Parsing form ("+err.Error()+")")
		return
	}

	user := r.FormValue("username")
	pass := r.FormValue("password")
	if !srv.authenticate(user, pass) {
		fmt.Fprintf(w, loginPage, "<h3>ERROR: Wrong username/pass</h3>")
		return
	}

	client.auth = true
	client.name = user
	srv.session.set(cookie, client)

	r.SetBasicAuth(user, pass)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (srv *server) authenticate(user, pass string) bool {
	v, ok := srv.session.password(user)
	if !ok {
		return false
	}

	return subtle.ConstantTimeCompare([]byte(pass), []byte(v)) == 1
}

func (srv *server) checkCredentials(w http.ResponseWriter, r *http.Request) (webClient, string, error) {
	var (
		ok     = false
		client webClient
	)
	cookie, err := r.Cookie("FCS_TOKEN")
	if err != nil {
		if err != http.ErrNoCookie {
			return client, "", err
		}
		err = nil
	}

	if cookie != nil {
		client, ok = srv.session.get(cookie.Value)
	}

	if !ok {
		cookie = &http.Cookie{
			Name:  "FCS_TOKEN",
			Value: uuid.NewV4().String(),
		}
		client = webClient{auth: false, token: cookie.Value}
		srv.session.set(cookie.Value, client)
	}

	http.SetCookie(w, cookie)
	return client, cookie.Value, nil
}

type webClient struct {
	name  string
	token string
	auth  bool
}

type authRegistry struct {
	store map[string]webClient
	mu    sync.RWMutex
	db    map[string]string
}

func newAuthRegistry() *authRegistry {
	db := make(map[string]string)
	f, err := os.Open("passwd.json")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&db)
	if err != nil {
		log.Fatalf("error decoding JSON db: %v\n", err)
	}

	return &authRegistry{
		store: make(map[string]webClient),
		db:    db,
	}
}

func (reg *authRegistry) password(user string) (string, bool) {
	v, ok := reg.db[user]
	return v, ok
}

func (reg *authRegistry) get(cookie string) (webClient, bool) {
	reg.mu.RLock()
	client, ok := reg.store[cookie]
	reg.mu.RUnlock()
	return client, ok
}

func (reg *authRegistry) set(cookie string, client webClient) {
	reg.mu.Lock()
	reg.store[cookie] = client
	reg.mu.Unlock()
}
