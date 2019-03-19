// Copyright ©2017 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"encoding/binary"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
	"github.com/go-lsst/ncs/drivers/m702"
)

var (
	codec = binary.BigEndian
	freq  = 50 * time.Millisecond
)

type Motor struct {
	addr string

	mu    sync.RWMutex
	state struct {
		Manual     uint32
		CmdReady   uint32
		FSM        uint32
		HWSafety   uint32
		Home       uint32
		ModePos    uint32
		RPMs       uint32
		WriteAngle uint32
		ReadAngle  uint32
		Temps      [4]float64
	}
}

func (m *Motor) run() {
	manual := time.After(5 * time.Second)
	hwsafety := time.After(6 * time.Second)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-manual:
			m.setManual(0)
		case <-hwsafety:
			m.setHWSafety(1)
		case <-ticker.C:
			m.state.FSM = 4
			m.genTemp()
		}
	}
}

func (m *Motor) setManual(v uint32) {
	m.mu.Lock()
	m.state.Manual = v
	m.mu.Unlock()
}

func (m *Motor) setHWSafety(v uint32) {
	m.mu.Lock()
	m.state.HWSafety = v
	m.mu.Unlock()
}

func (m *Motor) genTemp() {
	m.mu.Lock()
	for i := range m.state.Temps {
		m.state.Temps[i] = 2.5*rand.Float64() + (35 + 5*float64(i))
	}
	m.mu.Unlock()
}

func New(addr string) *Motor {
	return get(addr)
}

func (m *Motor) ReadParam(p *m702.Parameter) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	switch p.Index {
	case newParameter(bench.ParamManualOverride).Index:
		codec.PutUint32(p.Data[:], m.state.Manual)
	case newParameter(bench.ParamCmdReady).Index:
		codec.PutUint32(p.Data[:], m.state.CmdReady)
	case newParameter(bench.ParamHWSafety).Index:
		codec.PutUint32(p.Data[:], m.state.HWSafety)
	case newParameter(bench.ParamHome).Index:
		codec.PutUint32(p.Data[:], m.state.Home)
	case newParameter(bench.ParamModePos).Index:
		codec.PutUint32(p.Data[:], m.state.ModePos)
	case newParameter(bench.ParamRPMs).Index:
		codec.PutUint32(p.Data[:], m.state.RPMs)
	case newParameter(bench.ParamWritePos).Index:
		codec.PutUint32(p.Data[:], m.state.WriteAngle)
	case newParameter(bench.ParamReadPos).Index:
		codec.PutUint32(p.Data[:], m.state.ReadAngle)
	case newParameter(bench.ParamTemp0).Index:
		codec.PutUint32(p.Data[:], uint32(m.state.Temps[0]))
	case newParameter(bench.ParamTemp1).Index:
		codec.PutUint32(p.Data[:], uint32(m.state.Temps[1]))
	case newParameter(bench.ParamTemp2).Index:
		codec.PutUint32(p.Data[:], uint32(m.state.Temps[2]))
	case newParameter(bench.ParamTemp3).Index:
		codec.PutUint32(p.Data[:], uint32(m.state.Temps[3]))
	case newParameter(bench.ParamMotorStatus).Index:
		codec.PutUint32(p.Data[:], m.state.FSM)
	case newParameter(bench.ParamMotorStatusReady).Index:
		codec.PutUint32(p.Data[:], 1)
	case newParameter(bench.ParamMotorStatusActive).Index:
		codec.PutUint32(p.Data[:], 1)
	default:
		panic("invalid parameter: " + p.String())
	}
	return nil
}

func (m *Motor) WriteParam(p m702.Parameter) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	switch p.Index {
	case newParameter(bench.ParamManualOverride).Index:
		m.state.Manual = codec.Uint32(p.Data[:])
	case newParameter(bench.ParamCmdReady).Index:
		m.state.CmdReady = codec.Uint32(p.Data[:])
	case newParameter(bench.ParamHWSafety).Index:
		m.state.HWSafety = codec.Uint32(p.Data[:])
	case newParameter(bench.ParamHome).Index:
		m.state.Home = codec.Uint32(p.Data[:])
		if m.state.Home == 1 {
			go m.cmdGoHome()
		}
	case newParameter(bench.ParamModePos).Index:
		m.state.ModePos = codec.Uint32(p.Data[:])
	case newParameter(bench.ParamRPMs).Index:
		m.state.RPMs = codec.Uint32(p.Data[:])
	case newParameter(bench.ParamWritePos).Index:
		m.state.WriteAngle = codec.Uint32(p.Data[:])
		go m.cmdGoTo(m.state.WriteAngle)
	case newParameter(bench.ParamReadPos).Index:
		m.state.ReadAngle = codec.Uint32(p.Data[:])
	case newParameter(bench.ParamTemp0).Index:
		m.state.Temps[0] = float64(codec.Uint32(p.Data[:]))
	case newParameter(bench.ParamTemp1).Index:
		m.state.Temps[1] = float64(codec.Uint32(p.Data[:]))
	case newParameter(bench.ParamTemp2).Index:
		m.state.Temps[2] = float64(codec.Uint32(p.Data[:]))
	case newParameter(bench.ParamTemp3).Index:
		m.state.Temps[3] = float64(codec.Uint32(p.Data[:]))
	default:
		panic("invalid parameter: " + p.String())
	}
	return nil
}

func (m *Motor) cmdGoHome() {
	m.cmdGoTo(0)
}

func (m *Motor) cmdGoTo(v uint32) {
	tick := time.NewTicker(freq)
	defer tick.Stop()

	target := int32(v)
	for range tick.C {
		m.mu.Lock()
		angle := int32(m.state.ReadAngle)
		if angle > target {
			angle--
		}
		if angle < target {
			angle++
		}
		m.state.ReadAngle = uint32(angle)
		exit := m.state.ReadAngle == v
		m.mu.Unlock()
		if exit {
			return
		}
	}

}

var (
	mu     sync.RWMutex
	motors = make(map[string]*Motor)
)

func get(addr string) *Motor {
	mu.RLock()
	m, ok := motors[addr]
	mu.RUnlock()
	if ok {
		return m
	}
	mu.Lock()
	m = &Motor{addr: addr}
	m.state.Manual = 1
	m.state.HWSafety = 0
	m.state.Home = 1
	m.state.ModePos = 0
	m.state.RPMs = 1300
	go m.run()
	motors[addr] = m
	mu.Unlock()
	return m
}

func newParameter(name string) m702.Parameter {
	p, err := m702.NewParameter(name)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
