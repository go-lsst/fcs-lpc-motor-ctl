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

	"github.com/go-lsst/ncs/drivers/m702"
)

type scriptCmd func(args []string) error

type Script struct {
	cmds  map[string]scriptCmd
	motor m702.Motor
}

func (sc *Script) cmdGet(args []string) error {
	param, err := sc.parseParam(args[0])
	if err != nil {
		return err
	}

	return sc.motor.ReadParam(&param)
}

func (sc *Script) cmdSet(args []string) error {
	param, err := sc.parseParam(args[0])
	if err != nil {
		return err
	}
	vtype := "u32"
	if len(args) > 2 {
		vtype = args[2]
	}

	switch vtype {
	case "u32", "uint32":
		vv, err := strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			return err
		}
		codec.PutUint32(param.Data[:], uint32(vv))

	default:
		return fcsError{200, fmt.Sprintf("invalid value-type (%v)", vtype)}
	}

	return sc.motor.WriteParam(param)
}

func (sc *Script) run(motor m702.Motor, r io.Reader) error {
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
		err = sc.dispatch(toks)
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

func (sc *Script) dispatch(toks []string) error {
	fct, ok := sc.cmds[toks[0]]
	if !ok {
		return fcsError{200, fmt.Sprintf("invalid command verb [%s]", toks[0])}
	}
	return fct(toks[1:])
}

func (sc *Script) parseParam(arg string) (m702.Parameter, error) {
	if !strings.Contains(arg, ".") {
		var p m702.Parameter
		return p, fcsError{200, fmt.Sprintf("invalid parameter (%s)", arg)}
	}
	return m702.NewParameter(arg)
}

func newScripter(motor m702.Motor) Script {
	var script Script
	script = Script{
		cmds: map[string]scriptCmd{
			"get": script.cmdGet,
			"set": script.cmdSet,
		},
		motor: motor,
	}
	return script
}
