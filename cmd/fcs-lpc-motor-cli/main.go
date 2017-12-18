// Copyright ©2017 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command fcs-lpc-motor-cli is a simple REST client for the fcs-lpc-motor-ctl server.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("fcs-lpc-motor-cli: ")

	usr := flag.String("u", "faux-fcs", "user name for the authentication")
	pwd := flag.String("p", "faux-fcs", "user password for the authentication")
	addr := flag.String("addr", "http://clrbinetsrv.in2p3.fr:5555", "address:port of the fcs-lpc-motor-cli")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "expect a path to a command script file or a list of inline commands\n")
		flag.Usage()
		os.Exit(1)
	}

	// TODO(sbinet): use something else than a script file.
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

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

	req, err := http.NewRequest(http.MethodPost, *addr+"/api/cmd/req-upload-script", body)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.SetBasicAuth(*usr, *pwd)

	var cli http.Client
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatalf("could not POST request: %v", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatalf("could not stream out response: %v", err)
	}
}
