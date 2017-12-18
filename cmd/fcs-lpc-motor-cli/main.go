// Copyright ©2017 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command fcs-lpc-motor-cli is a simple REST client for the fcs-lpc-motor-ctl server.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
)

var (
	usr  = flag.String("u", "faux-fcs", "user name for the authentication")
	pwd  = flag.String("p", "faux-fcs", "user password for the authentication")
	addr = flag.String("addr", "http://clrbinetsrv.in2p3.fr:5555", "address:port of the fcs-lpc-motor-cli")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("fcs-lpc-motor-cli: ")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: fcs-lpc-motor-cli [options] [cmd-or-script-file]

ex:
 $> fcs-lpc-motor-cli ./test.script
 $> fcs-lpc-motor-cli x-angle-pos
 $> fcs-lpc-motor-cli z-angle-pos +20
 $> fcs-lpc-motor-cli

options:
`)
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	addr := *addr
	if !strings.HasPrefix(addr, "http") {
		addr = "http://" + addr
	}

	switch flag.NArg() {
	case 0:
		runMon(*usr, *pwd, addr)
	case 1:
		f, err := os.Open(flag.Arg(0))
		switch err {
		case nil:
			defer f.Close()
			runScript(f, *usr, *pwd, addr)
		default:
			runCommand(flag.Args(), *usr, *pwd, addr)
		}
	default:
		runCommand(flag.Args(), *usr, *pwd, addr)
	}
}

func runCommand(cmds []string, usr, pwd, addr string) {
	f, err := ioutil.TempFile("", "fcs-lpc-motor-cli-")
	if err != nil {
		log.Fatalf("could not create temporary script file: %v", err)
	}
	defer f.Close()
	defer os.Remove(f.Name())
	_, err = f.Write([]byte(strings.Join(cmds, " ")))
	if err != nil {
		log.Fatalf("could not generate temporary script file: %v", err)
	}
	f.Write([]byte("\n"))
	f.Seek(0, 0)
	runScript(f, usr, pwd, addr)
}

func runScript(f io.Reader, usr, pwd, addr string) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	part, err := w.CreateFormFile("upload-file", filepath.Base(flag.Arg(0)))
	if err != nil {
		log.Fatalf("could not create form-file: %v", err)
	}
	_, err = io.Copy(part, f)
	if err != nil {
		log.Fatalf("could not fill multipart: %v", err)
	}

	err = w.Close()
	if err != nil {
		log.Fatalf("could not close multipart: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, addr+"/api/cmd/req-upload-script", body)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.SetBasicAuth(usr, pwd)

	var cli http.Client
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatalf("could not POST request: %v", err)
	}
	defer resp.Body.Close()

	out := new(bytes.Buffer)
	var cmd struct {
		Code   int    `json:"code"`
		Error  string `json:"error"`
		Script string `json:"script"`
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("could not stream out response: %v", err)
	}

	err = json.Unmarshal(out.Bytes(), &cmd)
	if err != nil {
		log.Fatalf("could not unmarshal response: %v", err)
	}

	log.Printf("code: %v", cmd.Code)
	if cmd.Error != "" {
		log.Fatalf("err: %q", cmd.Error)
	}
	log.Printf("script:\n%s\n", cmd.Script)
}

func runMon(usr, pwd, addr string) {
	req, err := http.NewRequest(http.MethodGet, addr+"/api/mon", nil)
	if err != nil {
		log.Fatalf("could not create GET request: %v", err)
	}
	req.SetBasicAuth(usr, pwd)

	var cli http.Client
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatalf("could not send GET request: %v", err)
	}
	defer resp.Body.Close()

	var mon struct {
		Code  int                `json:"code"`
		Error string             `json:"error"`
		Infos []bench.MotorInfos `json:"infos"`
	}
	err = json.NewDecoder(resp.Body).Decode(&mon)
	if err != nil {
		log.Fatalf("could not decode monitoring stream: %v", err)
	}
	log.Printf("code: %v", mon.Code)
	if mon.Error != "" {
		log.Fatalf("err: %q", mon.Error)
	}
	for _, info := range mon.Infos {
		log.Printf("--- motor %v ---", info.Motor)
		log.Printf(" online:  %v", info.Online)
		log.Printf(" status:  %v", info.Status)
		log.Printf(" mode:    %v", info.Mode)
		log.Printf(" RPMs:    %v", info.RPMs)
		log.Printf(" angle:   %v", info.Angle)
		for i, temp := range info.Temps {
			log.Printf(" temp[%d]: %v", i, temp)
		}
	}
}
