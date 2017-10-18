// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
	"github.com/go-lsst/ncs/drivers/m702"
)

type scriptCmd func(args []string, w io.Writer) (m702.Parameter, error)

type Script struct {
	srv   *server
	cmds  map[string]scriptCmd
	motor bench.Motor
}

func (sc *Script) displayParam(w io.Writer, p m702.Parameter) {
	fmt.Fprintf(
		w,
		"Pr-%v: %s (%v)\n",
		p,
		sc.displayBytes(p.Data[:]),
		codec.Uint32(p.Data[:]),
	)
}

func (sc *Script) displayBytes(o []byte) string {
	hex := make([]string, len(o))
	dec := make([]string, len(o))
	for i, v := range o {
		hex[i] = fmt.Sprintf("0x%02x", v)
		dec[i] = fmt.Sprintf("%3d", v)
	}
	return fmt.Sprintf("hex=%s dec=%s", hex, dec)
}

func (sc *Script) cmdGet(args []string, w io.Writer) (m702.Parameter, error) {
	param, err := sc.parseParam(args[0])
	if err != nil {
		return param, err
	}
	err = sc.motor.ReadParam(&param)
	if err == nil {
		fmt.Fprintf(w, "<<< ")
		sc.displayParam(w, param)
	}

	return param, err
}

func (sc *Script) cmdSet(args []string, w io.Writer) (m702.Parameter, error) {
	param, err := sc.parseParam(args[0])
	if err != nil {
		return param, err
	}
	vtype := "u32"
	if len(args) > 1 && len(args[1]) > 0 && string(args[1][0]) == "-" {
		vtype = "i32"
	}
	if len(args) > 2 {
		vtype = args[2]
	}

	switch vtype {
	case "u32", "uint32":
		vv, err := strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			return param, err
		}
		codec.PutUint32(param.Data[:], uint32(vv))

	case "i32", "int32":
		vv, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			return param, err
		}
		codec.PutUint32(param.Data[:], uint32(vv))

	default:
		return param, bench.FcsError{200, fmt.Sprintf("invalid value-type (%v)", vtype)}
	}

	err = sc.motor.WriteParam(param)
	if err == nil {
		fmt.Fprintf(w, "<<< ")
		sc.displayParam(w, param)
	}

	return param, err
}

func (sc *Script) cmdMotor(args []string, w io.Writer) (m702.Parameter, error) {
	var p m702.Parameter
	switch len(args) {
	case 0:
		// get
		return p, bench.FcsError{200, fmt.Sprintf("invalid number of arguments (got=%d, want=1|2)", len(args))}
	case 1:
		// set
		motors := sc.srv.motors()
		switch strings.ToLower(args[0]) {
		case "x":
			sc.motor = motors[0].Motor()
		case "z":
			sc.motor = motors[1].Motor()
		default:
			return p, bench.FcsError{200, fmt.Sprintf("invalid motor name (got=%v, want=x|z)", args[0])}
		}
		return p, nil
	default:
		return p, bench.FcsError{200, fmt.Sprintf("invalid number of arguments (got=%d, want=1|2)", len(args))}
	}
}

func (sc *Script) run(motor bench.Motor, r io.Reader, w io.Writer) error {
	var (
		err      error
		oldMotor = sc.motor
	)
	defer func() {
		sc.motor = oldMotor
	}()

	sc.motor = motor
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		txt := strings.TrimSpace(scan.Text())
		if txt == "" {
			continue
		}
		toks := strings.Split(txt, " ")
		_, err = sc.dispatch(toks, w)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return err
		}
	}

	err = scan.Err()
	if err == io.EOF {
		err = nil
	}
	return err
}

func (sc *Script) dispatch(toks []string, w io.Writer) (m702.Parameter, error) {
	fct, ok := sc.cmds[toks[0]]
	if !ok {
		var p m702.Parameter
		return p, bench.FcsError{200, fmt.Sprintf("invalid command verb [%s]", toks[0])}
	}
	fmt.Fprintf(w, ">>> %s\n", strings.Join(toks, " "))
	return fct(toks[1:], w)
}

func (sc *Script) parseParam(arg string) (m702.Parameter, error) {
	if !strings.Contains(arg, ".") {
		var p m702.Parameter
		return p, bench.FcsError{200, fmt.Sprintf("invalid parameter (%s)", arg)}
	}
	return m702.NewParameter(arg)
}

func newScripter(srv *server, motor bench.Motor) Script {
	var script Script
	script = Script{
		srv: srv,
		cmds: map[string]scriptCmd{
			"get":   script.cmdGet,
			"set":   script.cmdSet,
			"motor": script.cmdMotor,
		},
		motor: motor,
	}
	return script
}
