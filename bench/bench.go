// Copyright ©2017 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bench

import (
	"fmt"

	"github.com/go-lsst/ncs/drivers/m702"
)

const (
	ParamManualOverride = "0.08.005" // 0:sw, 1:manual-override
	ParamCmdReady       = "0.08.015"
	ParamHome           = "2.02.017"
	ParamModePos        = "2.02.011"
	ParamRPMs           = "0.20.022"
	ParamWritePos       = "3.70.000"
	ParamReadPos        = "0.18.002"
	ParamTemp0          = "0.07.004"
	ParamTemp1          = "0.07.005"
	ParamTemp2          = "0.07.006"
	ParamTemp3          = "0.07.034"

	ParamHWSafety = "0.08.040" // 0:OK, 1:HW-Safety ON
)

var (
	ErrMotorOffline     = FcsError{1, "fcs: motor OFFLINE"}
	ErrMotorHWLock      = FcsError{2, "fcs: motor HW-safety enabled"}
	ErrMotorManual      = FcsError{3, "fcs: motor manual-mode enabled"}
	ErrOpNotSupported   = FcsError{20, "fcs: operation not supported"}
	ErrInvalidReq       = FcsError{102, "fcs: invalid request"}
	ErrInvalidMotorName = FcsError{200, "fcs: invalid motor name"}
)

type FcsError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e FcsError) Error() string {
	return fmt.Sprintf("[%03d]: %s", e.Code, e.Msg)
}

func (e FcsError) String() string {
	return e.Error()
}

type Bench struct {
	Motor struct {
		X Motor
		Z Motor
	}
}

// Motor is the interface to the FCS testbench motors
type Motor interface {
	//State() m702.FSM
	ReadParam(p *m702.Parameter) error
	WriteParam(p m702.Parameter) error
}

func NewMotorFrom(m m702.Motor) Motor {
	return &m
}
