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
	ParamRPMs           = "3.70.064"
	ParamWritePos       = "3.70.000"
	ParamReadPos        = "0.18.002"
	ParamTemp0          = "0.07.004"
	ParamTemp1          = "0.07.005"
	ParamTemp2          = "0.07.006"
	ParamTemp3          = "0.07.034"

	ParamHWSafety = "0.08.040" // 0:HW-Safety On, 1:HW-Safety Off

	// Number of messages exchanged b/w master and slave.
	// This number needs to be >0 on the slave to indicate the connection
	// is correctly established.
	ParamMasterSlaveExchange = "4.10.004"

	// Ratio load (in %) on the master/slave axis.
	ParamMasterSlaveLoadRatio = "4.020"

	// Encoder position value (master/slave synchro monitoring)
	ParamMotorPosMonitor = "3.58"

	ParamMotorStatusReady  = "10.001" // Motor ready:  0=off, 1=on
	ParamMotorStatusActive = "10.002" // Motor active: 0=off, 1=on

	ParamPositionReached = "0.19.033" // position: 0=no, 1=reached

	ParamMotorStatus = "0.60"  // overall motor status (copied from 10.101)
	ParamMotorReset  = "10.33" // 0 -> 1

	ParamFwdLimitSwitch = "6.35" // forward limit marker
	ParamBwdLimitSwitch = "6.36" // backward limit marker
)

var (
	ErrMotorOffline     = FcsError{1, "fcs: motor OFFLINE"}
	ErrMotorHWLock      = FcsError{2, "fcs: motor HW-safety enabled"}
	ErrMotorManual      = FcsError{3, "fcs: motor manual-mode enabled"}
	ErrOpNotSupported   = FcsError{20, "fcs: operation not supported"}
	ErrInvalidReq       = FcsError{102, "fcs: invalid request"}
	ErrInvalidMotorName = FcsError{200, "fcs: invalid motor name"}
	ErrInvalidQuadrant  = FcsError{300, "fcs: invalid quadrant for find-home"}
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

type MotorInfos struct {
	Motor  string            `json:"motor"` // x,z
	Online bool              `json:"online"`
	Status string            `json:"status"` // N/A,manual,hw-safety,ready
	FSM    string            `json:"fsm"`    // HW-safety, Ready, Stop, Scan, Run, Trip
	Sync   bool              `json:"sync"`   // master/slave synchronization
	Mode   string            `json:"mode"`   // N/A,ready,home,position
	RPMs   int               `json:"rpms"`
	Angle  float64           `json:"angle"`
	Temps  [4]float64        `json:"temps"`
	Histos map[string]string `json:"histos"`
	Webcam string            `json:"webcam"`
}
