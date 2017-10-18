// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
	"github.com/go-lsst/fcs-lpc-motor-ctl/mock"
	"github.com/go-lsst/ncs/drivers/m702"
)

type motor struct {
	name   string
	addr   string
	mock   bool
	params motorParams
	histos motorHistos
	online bool // whether motors are online/connected

	mu sync.RWMutex // see motor.updateAnglePos
}

func (m *motor) Motor() bench.Motor {
	if m.mock {
		return mock.New(m.addr)
	}
	return bench.NewMotorFrom(m702.New(m.addr))
}

func (m *motor) poll() []error {
	var errs []error
	mm := m.Motor()
	for _, p := range []*m702.Parameter{
		&m.params.Manual,
		&m.params.HWSafety,
		&m.params.Home,
		&m.params.ModePos,
		&m.params.RPMs,
		&m.params.ReadAngle,
		&m.params.Temps[0],
		&m.params.Temps[1],
		&m.params.Temps[2],
		&m.params.Temps[3],
	} {
		err := mm.ReadParam(p)
		if err != nil {
			errs = append(errs, fmt.Errorf("error reading %v (motor-%s) Pr-%v: %v\n", m.addr, m.name, *p, err))
		}
	}
	return errs
}

func (m *motor) isOnline(timeout time.Duration) (bool, error) {
	if m.mock {
		return true, nil
	}
	online := false
	c, err := net.DialTimeout("tcp", m.addr, timeout)
	if c != nil {
		defer c.Close()
	}
	if err == nil && c != nil {
		online = true
	}
	return online, err
}

func (m *motor) isHWLocked() bool {
	return codec.Uint32(m.params.HWSafety.Data[:]) == 0
}

func (m *motor) isManual() bool {
	return codec.Uint32(m.params.Manual.Data[:]) == 1
}

func (m *motor) rpms() uint32 {
	return codec.Uint32(m.params.RPMs.Data[:])
}

func (m *motor) angle() float64 {
	return float64(int32(codec.Uint32(m.params.ReadAngle.Data[:]))) * 0.1
}

// updateAnglePos is used to track the current motor position when
// in manual-mode (and the operator is moving the motor using
// a physical controller.)
// This prevents the motor from going back to the last position (before
// being moved via the manual-mode)
func (m *motor) updateAnglePos() error {
	m.mu.Lock()
	pos := codec.Uint32(m.params.ReadAngle.Data[:])
	mm := m702.New(m.addr)
	param := newParameter(bench.ParamWritePos)
	codec.PutUint32(param.Data[:], pos)
	err := mm.WriteParam(param)
	m.mu.Unlock()
	return err
}

func newMotor(name, addr string) motor {
	return motor{
		name:   name,
		addr:   addr,
		params: newMotorParams(),
		histos: motorHistos{
			rows: make([]monData, 0, 128),
		},
	}
}

func newMotorMock(name, addr string) motor {
	m := newMotor(name, addr)
	m.mock = true
	return m
}

type motorParams struct {
	Manual     m702.Parameter
	CmdReady   m702.Parameter
	HWSafety   m702.Parameter
	Home       m702.Parameter
	ModePos    m702.Parameter
	RPMs       m702.Parameter
	WriteAngle m702.Parameter
	ReadAngle  m702.Parameter
	Temps      [4]m702.Parameter
}

func newMotorParams() motorParams {
	return motorParams{
		Manual:     newParameter(bench.ParamManualOverride),
		CmdReady:   newParameter(bench.ParamCmdReady),
		HWSafety:   newParameter(bench.ParamHWSafety),
		Home:       newParameter(bench.ParamHome),
		ModePos:    newParameter(bench.ParamModePos),
		RPMs:       newParameter(bench.ParamRPMs),
		WriteAngle: newParameter(bench.ParamWritePos),
		ReadAngle:  newParameter(bench.ParamReadPos),
		Temps: [4]m702.Parameter{
			newParameter(bench.ParamTemp0),
			newParameter(bench.ParamTemp1),
			newParameter(bench.ParamTemp2),
			newParameter(bench.ParamTemp3),
		},
	}
}

type motorHistos struct {
	rows []monData
	Temp [4][]float64
	Pos  []float64
	RPMs []float64
}

type motorInfos struct {
	Motor  string            `json:"motor"` // x,z
	Online bool              `json:"online"`
	Status string            `json:"status"` // N/A,manual,hw-safety,ready
	Mode   string            `json:"mode"`   // N/A,ready,home,position
	RPMs   int               `json:"rpms"`
	Angle  int               `json:"angle"`
	Temps  [4]float64        `json:"temps"`
	Histos map[string]string `json:"histos"`
	Webcam string            `json:"webcam"`
}
