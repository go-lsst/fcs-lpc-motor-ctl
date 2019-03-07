// Copyright ©2019 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
	"github.com/go-lsst/ncs/drivers/m702"
)

type scriptCmd func(args []string, w io.Writer) (m702.Parameter, error)

type Script struct {
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
		// fmt.Fprintf(w, "<<< ")
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
		// fmt.Fprintf(w, "<<< ")
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

func (sc *Script) cmdFindHomeX(args []string, w io.Writer) (m702.Parameter, error) {
	return sc.cmdFindHome(&motors[0], args, w)
}

func (sc *Script) cmdPosX(args []string, w io.Writer) (m702.Parameter, error) {
	return sc.cmdPos(&motors[0], args, w)
}

func (sc *Script) cmdRPMX(args []string, w io.Writer) (m702.Parameter, error) {
	return sc.cmdRPM(&motors[0], args, w)
}

func (sc *Script) cmdAnglePosX(args []string, w io.Writer) (m702.Parameter, error) {
	return sc.cmdAnglePos(&motors[0], args, w)
}

func (sc *Script) cmdFindHomeZ(args []string, w io.Writer) (m702.Parameter, error) {
	return sc.cmdFindHome(&motors[1], args, w)
}

func (sc *Script) cmdPosZ(args []string, w io.Writer) (m702.Parameter, error) {
	return sc.cmdPos(&motors[1], args, w)
}

func (sc *Script) cmdRPMZ(args []string, w io.Writer) (m702.Parameter, error) {
	return sc.cmdRPM(&motors[1], args, w)
}

func (sc *Script) cmdAnglePosZ(args []string, w io.Writer) (m702.Parameter, error) {
	return sc.cmdAnglePos(&motors[1], args, w)
}

func (sc *Script) cmdFindHome(m *motor, args []string, w io.Writer) (m702.Parameter, error) {
	var p m702.Parameter
	err := sc.check(m)
	if err != nil {
		return p, err
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
		err := m.Motor().WriteParam(p)
		if err != nil {
			return p, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err)
		}
	}

	return p, nil
}

func (sc *Script) cmdPos(m *motor, args []string, w io.Writer) (m702.Parameter, error) {
	var p m702.Parameter
	err := sc.check(m)
	if err != nil {
		return p, err
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
			return p, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err)
		}
	}

	return p, nil
}

func (sc *Script) cmdRPM(m *motor, args []string, w io.Writer) (m702.Parameter, error) {
	var p m702.Parameter
	err := sc.check(m)
	if err != nil {
		return p, err
	}

	switch len(args) {
	case 0:
		// TODO(sbinet): only retrieve the needed infos.
		info, err := m.infos(1 * time.Second)
		if err != nil {
			return p, err
		}
		fmt.Fprintf(w, "get-%v-rpm=%d\n", m.name, info.RPMs)
	case 1:
		rpm, err := strconv.Atoi(args[0])
		if err != nil {
			return p, err
		}
		switch {
		case rpm > 3000:
			return p, fmt.Errorf("invalid RPM value (%v > 3000)", rpm)
		case rpm < 0:
			return p, fmt.Errorf("invalid RPM value (%v < 0)", rpm)
		}
		p = newParameter(bench.ParamRPMs)
		codec.PutUint32(p.Data[:], uint32(rpm))
		err = m.Motor().WriteParam(p)
		if err != nil {
			return p, err
		}
		fmt.Fprintf(w, "set-%v-rpm=%d\n", m.name, rpm)
		return p, nil

	default:
		return p, fmt.Errorf("fcs: invalid number of parameters (got=%d. want=0,1)", len(args))
	}
	return p, nil
}

func (sc *Script) cmdAnglePos(m *motor, args []string, w io.Writer) (m702.Parameter, error) {
	var p m702.Parameter
	err := sc.check(m)
	if err != nil {
		return p, err
	}

	switch len(args) {
	case 0:
		// TODO(sbinet): only retrieve the needed infos.
		info, err := m.infos(1 * time.Second)
		if err != nil {
			return p, err
		}
		fmt.Fprintf(w, "get-%v-angle-pos=%d\n", m.name, info.Angle)
	case 1:
		angle, err := strconv.Atoi(args[0])
		if err != nil {
			return p, err
		}
		switch {
		case angle > +90:
			return p, fmt.Errorf("invalid angle position (%v > +90.0)", angle)

		case angle < -90:
			return p, fmt.Errorf("invalid angle position (%v < -90.0)", angle)
		}

		p = newParameter(bench.ParamWritePos)
		codec.PutUint32(p.Data[:], uint32(angle*10))
		err = m.Motor().WriteParam(p)
		if err != nil {
			return p, err
		}
		fmt.Fprintf(w, "set-%v-angle-pos=%d\n", m.name, angle)
		return p, nil

	default:
		return p, fmt.Errorf("fcs: invalid number of parameters (got=%d. want=0,1)", len(args))
	}
	return p, nil
}

func (sc *Script) cmdSleep(args []string, w io.Writer) (m702.Parameter, error) {
	var p m702.Parameter
	if len(args) != 1 {
		return p, fmt.Errorf("fcs: invalid number of parameters (got=%d. want=1)", len(args))
	}

	t, err := time.ParseDuration(args[0])
	if err != nil {
		return p, fmt.Errorf("fcs: could not parse 'sleep %s': %v", args[0], err)
	}

	time.Sleep(t)
	return p, nil
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
	// fmt.Fprintf(w, ">>> %s\n", strings.Join(toks, " "))
	return fct(toks[1:], w)
}

func (sc *Script) check(m *motor) error {
	online, err := m.isOnline(motorTimeout)
	if err != nil {
		return fmt.Errorf("fcs: error checking whether motor-%v is online: %v", m.name, err)
	}
	if !online {
		return bench.ErrMotorOffline
		// return fmt.Sprintf("fcs: motor-%v is NOT ONLINE", m.name)
	}
	if m.isHWLocked() {
		return bench.ErrMotorHWLock
	}
	if m.isManual() {
		return bench.ErrMotorManual
	}
	return nil
}

func (sc *Script) parseParam(arg string) (m702.Parameter, error) {
	if !strings.Contains(arg, ".") {
		var p m702.Parameter
		return p, bench.FcsError{200, fmt.Sprintf("invalid parameter (%s)", arg)}
	}
	return m702.NewParameter(arg)
}

func newScripter(motor bench.Motor) Script {
	var script Script
	script = Script{
		cmds: map[string]scriptCmd{
			"get":         script.cmdGet,
			"set":         script.cmdSet,
			"sleep":       script.cmdSleep,
			"motor":       script.cmdMotor,
			"x-find-home": script.cmdFindHomeX,
			"x-pos":       script.cmdPosX,
			"x-rpm":       script.cmdRPMX,
			"x-angle-pos": script.cmdAnglePosX,
			"z-find-home": script.cmdFindHomeZ,
			"z-pos":       script.cmdPosZ,
			"z-rpm":       script.cmdRPMZ,
			"z-angle-pos": script.cmdAnglePosZ,
		},
		motor: motor,
	}
	return script
}
