// Copyright ©2019 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net"
	"sync"
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
	"github.com/go-lsst/fcs-lpc-motor-ctl/mock"
	"github.com/go-lsst/ncs/drivers/m702"
)

func newMotor(name, addr string) motor {
	return motor{
		name:   name,
		addr:   addr,
		params: newMotorParams(),
	}
}

type motor struct {
	name   string
	addr   string
	mock   bool
	params motorParams
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
	m.mu.Lock()
	defer m.mu.Unlock()

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

func (m *motor) infos(timeout time.Duration) (infos bench.MotorInfos, err error) {
	online, err := m.isOnline(timeout)
	if err != nil {
		return infos, err
	}
	if !online {
		infos = bench.MotorInfos{
			Motor:  m.name,
			Online: online,
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
		Mode:   mon.Mode(),
		RPMs:   int(mon.rpms),
		Angle:  mon.angle,
		Temps:  mon.temps,
	}

	return infos, err
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

type motorMode byte

const (
	motorModeDefault motorMode = iota
	motorModeHome
	motorModePos
)

type monData struct {
	id    time.Time
	mode  motorMode
	rpms  uint32
	angle float64
	temps [4]float64
}

const (
	monDataLen = 54
)

func (mon *monData) x() float64 {
	mon.buflen()
	return float64(mon.id.Unix())
}

func (mon *monData) write(buf []byte) {
	i := 0
	binary.LittleEndian.PutUint64(buf[i:i+8], uint64(mon.id.Unix()))
	i += 8
	binary.LittleEndian.PutUint16(buf[i:i+2], uint16(mon.mode))
	i += 2
	binary.LittleEndian.PutUint32(buf[i:i+4], mon.rpms)
	i += 4
	binary.LittleEndian.PutUint64(buf[i:i+8], math.Float64bits(mon.angle))
	i += 8
	for _, temp := range mon.temps {
		binary.LittleEndian.PutUint64(buf[i:i+8], math.Float64bits(temp))
		i += 8
	}
}

func (mon *monData) Mode() string {
	switch mon.mode {
	case motorModeDefault:
		return "N/A"
	case motorModeHome:
		return "home"
	case motorModePos:
		return "pos"
	default:
		panic(fmt.Errorf("invalid monData.mode=%v", mon.mode))
	}
	panic("unreachable")
}

func (mon *monData) buflen() int {
	sz := 0
	sz += 8     // mon.id
	sz += 2     // mon.mode
	sz += 4     // mon.rpms
	sz += 8     // mon.angle
	sz += 4 * 8 // mon.temps
	return sz
}

func init() {
	blen := ((*monData)(nil)).buflen()
	if blen != monDataLen {
		panic(fmt.Errorf("fcs: monData buffer sanity check: blen=%d want=%d", blen, monDataLen))
	}
}

func newParameter(name string) m702.Parameter {
	p, err := m702.NewParameter(name)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
