// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
	"github.com/go-lsst/fcs-lpc-motor-ctl/mock"
	"github.com/go-lsst/ncs/drivers/m702"
	"github.com/pkg/errors"
)

func (srv *server) getMotor(name string) (*motor, error) {
	var m *motor
	switch strings.ToLower(name) {
	case "x":
		m = &srv.motor.x
	case "z":
		m = &srv.motor.z
	default:
		return nil, bench.ErrInvalidMotorName
	}
	return m, nil
}

type motor struct {
	name   string
	addr   string
	mock   bool
	params motorParams
	histos motorHistos
	online bool // whether motors are online/connected

	slave  *motor // nil if no master/slave
	master *motor

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
		&m.params.CmdReady,
		&m.params.MotorStatus,
		&m.params.MotorReady,
		&m.params.MotorActive,
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
		err := m.retry(func() error { return mm.ReadParam(p) })
		if err != nil {
			errs = append(errs, fmt.Errorf("error reading %v (motor-%s) Pr-%v: %v\n", m.addr, m.name, *p, err))
		}
	}
	if m.slave != nil {
		slave := m.slave.pollSlave()
		if len(slave) > 0 {
			errs = append(errs, slave...)
		}
	}
	return errs
}

func (m *motor) pollSlave() []error {
	var errs []error
	mm := m.Motor()
	for _, p := range []*m702.Parameter{
		&m.params.Manual,
		&m.params.CmdReady,
		&m.params.MotorStatus,
		&m.params.MotorReady,
		&m.params.MotorActive,
		&m.params.HWSafety,
		&m.params.Temps[0],
		&m.params.Temps[1],
		&m.params.Temps[2],
		&m.params.Temps[3],
	} {
		err := m.retry(func() error { return mm.ReadParam(p) })
		if err != nil {
			errs = append(errs, fmt.Errorf("error reading %v (motor-%s) Pr-%v: %v\n", m.addr, m.name, *p, err))
		}
	}
	return errs
}

func (m *motor) isSlave() bool { return m.master != nil }

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

func (m *motor) fsm() string {
	v := codec.Uint32(m.params.MotorStatus.Data[:])
	switch v {
	case 0:
		return "inhibit"
	case 1:
		return "ready"
	case 2:
		return "stop"
	case 3:
		return "scan"
	case 4:
		return "run"
	case 5:
		return "supply loss"
	case 6:
		return "deceleration"
	case 7:
		return "dc injection"
	case 8:
		return "position"
	case 9:
		return "trip"
	case 10:
		return "active"
	case 11:
		return "off"
	case 12:
		return "hand"
	case 13:
		return "auto"
	case 14:
		return "heat"
	case 15:
		return "under voltage"
	case 16:
		return "phasing"
	}
	return fmt.Sprintf("fsm=%d", v)
}

func (m *motor) isHWLocked() bool {
	return codec.Uint32(m.params.HWSafety.Data[:]) == 0
}

func (m *motor) isManual() bool {
	return codec.Uint32(m.params.Manual.Data[:]) == 1
}

func (m *motor) isSyncOK() bool {
	switch m.name {
	case "x":
		master, err1 := m.sync()
		slave, err2 := m.slave.sync()
		if err1 != nil || err2 != nil {
			return false
		}
		if master <= 10 || slave <= 10 {
			return false
		}
		return math.Abs(float64(master-slave)) < 20
	case "z":
		return true
	default:
		return false
	}
}

func (m *motor) rpms() uint32 {
	return codec.Uint32(m.params.RPMs.Data[:])
}

func (m *motor) angle() float64 {
	return float64(int32(codec.Uint32(m.params.ReadAngle.Data[:]))) * 0.1
}

func (m *motor) sync() (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	mm := m702.New(m.addr)
	p := newParameter(bench.ParamMasterSlaveExchange)
	err := m.retry(func() error { return mm.ReadParam(&p) })
	if err != nil {
		return 0, errors.Wrapf(err, "motor %q: could not read %v", m.name, p)
	}

	return int(codec.Uint32(p.Data[:])), nil
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

func (m *motor) reset() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.slave != nil {
		err := m.slave.reset()
		if err != nil {
			return err
		}
	}

	mm := m702.New(m.addr)
	ps := append([]m702.Parameter{},
		newParameter(bench.ParamMotorReset),
		newParameter(bench.ParamMotorReset),
		newParameter(bench.ParamMotorReset),
		newParameter(bench.ParamCmdReady),
		newParameter(bench.ParamCmdReady),
		newParameter(bench.ParamCmdReady),
	)
	codec.PutUint32(ps[0].Data[:], 0)
	codec.PutUint32(ps[1].Data[:], 1)
	codec.PutUint32(ps[2].Data[:], 0)
	codec.PutUint32(ps[3].Data[:], 1)
	codec.PutUint32(ps[4].Data[:], 0)
	codec.PutUint32(ps[5].Data[:], 1)

	for _, p := range ps {
		err := m.retry(func() error { return mm.WriteParam(p) })
		if err != nil {
			return errors.Wrapf(err, "motor %q: could not send reset", m.name)
		}
	}

	pos := newParameter(bench.ParamReadPos)
	err := m.retry(func() error { return mm.ReadParam(&pos) })
	if err != nil {
		return errors.Wrapf(err, "motor %q: could not read position during reset", m.name)
	}

	if !m.isSlave() {
		wpos := newParameter(bench.ParamWritePos)
		codec.PutUint32(wpos.Data[:], codec.Uint32(pos.Data[:]))
		err = m.retry(func() error { return mm.WriteParam(wpos) })
		if err != nil {
			return errors.Wrapf(err, "motor %q: could not update position during reset", m.name)
		}
	}

	err = m.retry(func() error { return mm.ReadParam(&pos) })
	if err != nil {
		return errors.Wrapf(err, "motor %q: could not read position during reset", m.name)
	}

	return nil
}

func (m *motor) stop() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	mm := m702.New(m.addr)

	ps := append([]m702.Parameter{},
		newParameter(bench.ParamModePos),
		newParameter(bench.ParamHome),
		newParameter(bench.ParamCmdReady),
	)
	codec.PutUint32(ps[0].Data[:], 0)
	codec.PutUint32(ps[1].Data[:], 0)
	codec.PutUint32(ps[2].Data[:], 0)

	for _, p := range ps {
		err := m.retry(func() error { return mm.WriteParam(p) })
		if err != nil {
			return errors.Wrapf(err, "motor %q: could not send reset", m.name)
		}
	}

	go func() {
		m.mu.Lock()
		defer m.mu.Unlock()

		time.Sleep(2 * time.Second)

		pos := newParameter(bench.ParamReadPos)
		err := m.retry(func() error { return mm.ReadParam(&pos) })
		if err != nil {
			log.Printf("motor %q: could not read position during reset: %v", m.name, err)
			return
		}

		wpos := newParameter(bench.ParamWritePos)
		codec.PutUint32(wpos.Data[:], codec.Uint32(pos.Data[:]))
		err = m.retry(func() error { return mm.WriteParam(wpos) })
		if err != nil {
			log.Printf("motor %q: could not update position during reset: %v", m.name, err)
			return
		}

		err = m.retry(func() error { return mm.ReadParam(&pos) })
		if err != nil {
			log.Printf("motor %q: could not read position during reset: %v", m.name, err)
		}
	}()

	return nil
}

func (m *motor) findHome() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	mm := m702.New(m.addr)

	err := m.retry(func() error { return mm.ReadParam(&m.params.ReadAngle) })
	if err != nil {
		return errors.Wrapf(err, "motor %q: could not read position before find-home", m.name)
	}

	switch m.name {
	case "x":
		if m.angle() > 0 {
			return bench.ErrInvalidQuadrant
		}
	case "z":
		if m.angle() < 0 {
			return bench.ErrInvalidQuadrant
		}
	default:
		return bench.ErrInvalidMotorName
	}

	ps := append([]m702.Parameter{},
		newParameter(bench.ParamCmdReady),
		newParameter(bench.ParamModePos),
		newParameter(bench.ParamHome),
		newParameter(bench.ParamCmdReady),
	)
	codec.PutUint32(ps[0].Data[:], 0)
	codec.PutUint32(ps[1].Data[:], 0)
	codec.PutUint32(ps[2].Data[:], 1)
	codec.PutUint32(ps[3].Data[:], 1)

	for _, p := range ps {
		err := m.retry(func() error { return mm.WriteParam(p) })
		if err != nil {
			return errors.Wrapf(err, "motor %q: could not send reset", m.name)
		}
	}

	return nil
}

func (m *motor) infos(timeout time.Duration) (infos bench.MotorInfos, err error) {
	online, err := m.isOnline(timeout)
	if err != nil {
		return infos, err
	}
	if !online {
		infos = bench.MotorInfos{
			Motor:  m.name,
			Online: online,
			FSM:    "N/A",
			Sync:   false,
			Mode:   "N/A",
		}
		return infos, err
	}

	errs := m.poll()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("%v", err)
		}
		return infos, errs[0]
	}

	if m.isManual() {
		// make sure we won't override what manual-mode did
		// when we go back to sw-mode/ready-mode
		err = m.updateAnglePos()
		if err != nil {
			log.Printf("-- motor-%v: standby: %v\n", m.name, err)
		}
		return infos, err
	}

	mon := monData{
		id:    time.Now(),
		rpms:  m.rpms(),
		angle: m.angle(),
		temps: [4]float64{
			float64(codec.Uint32(m.params.Temps[0].Data[:])),
			float64(codec.Uint32(m.params.Temps[1].Data[:])),
			float64(codec.Uint32(m.params.Temps[2].Data[:])),
			float64(codec.Uint32(m.params.Temps[3].Data[:])),
		},
	}

	status := "N/A"

	manual := m.isManual()
	ready := !manual
	hwsafetyON := m.isHWLocked()
	fsm := m.fsm()
	sync := m.isSyncOK()

	switch {
	case hwsafetyON:
		status = "h/w safety"
	case manual:
		status = "manual"
	case ready:
		status = "ready"
	}

	if online {
		switch {
		case codec.Uint32(m.params.Home.Data[:]) == 1:
			mon.mode = motorModeHome
		case codec.Uint32(m.params.ModePos.Data[:]) == 1:
			mon.mode = motorModePos
		}
	}
	infos = bench.MotorInfos{
		Motor:  m.name,
		Online: online,
		Status: status,
		FSM:    fsm,
		Sync:   sync,
		Mode:   mon.Mode(),
		RPMs:   int(mon.rpms),
		Angle:  int(mon.angle),
		Temps:  mon.temps,
	}

	return infos, err
}

func (m *motor) retry(f func() error) error {
	const retries = 10
	var err error
	for i := 0; i < retries; i++ {
		err = f()
		if err == nil {
			return err
		}
	}
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

func newMotorSlave(name, addr string, master *motor) motor {
	return motor{
		name:   name + "-slave",
		addr:   addr,
		params: newMotorSlaveParams(),
		histos: motorHistos{
			rows: make([]monData, 0, 128),
		},
		master: master,
	}
}

func newMotorMock(name, addr string) motor {
	m := newMotor(name, addr)
	m.mock = true
	return m
}

type motorParams struct {
	Manual      m702.Parameter
	CmdReady    m702.Parameter
	MotorReady  m702.Parameter
	MotorActive m702.Parameter
	MotorStatus m702.Parameter
	HWSafety    m702.Parameter
	Home        m702.Parameter
	ModePos     m702.Parameter
	RPMs        m702.Parameter
	WriteAngle  m702.Parameter
	ReadAngle   m702.Parameter
	Temps       [4]m702.Parameter
}

func newMotorParams() motorParams {
	return motorParams{
		Manual:      newParameter(bench.ParamManualOverride),
		CmdReady:    newParameter(bench.ParamCmdReady),
		MotorStatus: newParameter(bench.ParamMotorStatus),
		MotorReady:  newParameter(bench.ParamMotorStatusReady),
		MotorActive: newParameter(bench.ParamMotorStatusActive),
		HWSafety:    newParameter(bench.ParamHWSafety),
		Home:        newParameter(bench.ParamHome),
		ModePos:     newParameter(bench.ParamModePos),
		RPMs:        newParameter(bench.ParamRPMs),
		WriteAngle:  newParameter(bench.ParamWritePos),
		ReadAngle:   newParameter(bench.ParamReadPos),
		Temps: [4]m702.Parameter{
			newParameter(bench.ParamTemp0),
			newParameter(bench.ParamTemp1),
			newParameter(bench.ParamTemp2),
			newParameter(bench.ParamTemp3),
		},
	}
}

func newMotorSlaveParams() motorParams {
	return motorParams{
		Manual:      newParameter(bench.ParamManualOverride),
		CmdReady:    newParameter(bench.ParamCmdReady),
		MotorStatus: newParameter(bench.ParamMotorStatus),
		MotorReady:  newParameter(bench.ParamMotorStatusReady),
		MotorActive: newParameter(bench.ParamMotorStatusActive),
		HWSafety:    newParameter(bench.ParamHWSafety),
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
