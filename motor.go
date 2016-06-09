// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/go-lsst/ncs/drivers/m702"
)

type motor struct {
	name   string
	addr   string
	params motorParams
	histos motorHistos
	online bool // whether motors are online/connected
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

type motorParams struct {
	Ready     m702.Parameter
	Rotation0 m702.Parameter
	Rotation1 m702.Parameter
	RPMs      m702.Parameter
	Angle     m702.Parameter
	Temps     [4]m702.Parameter
}

func newMotorParams() motorParams {
	return motorParams{
		Ready:     newParameter(paramReady),
		Rotation0: newParameter(paramRotation0),
		Rotation1: newParameter(paramRotation1),
		RPMs:      newParameter(paramRPMs),
		// Angle: newParameter(""), // FIXME(sbinet)
		Temps: [4]m702.Parameter{
			newParameter(paramTemp0),
			newParameter(paramTemp1),
			newParameter(paramTemp2),
			newParameter(paramTemp3),
		},
	}
}

type motorHistos struct {
	rows []monData
	Temp [4][]float64
	Pos  []float64
	RPMs []float64
}

type motorStatus struct {
	Motor    string            `json:"motor"`
	Online   bool              `json:"online"`
	Ready    bool              `json:"ready"`
	Rotation int               `json:"rotation_direction"`
	RPMs     int               `json:"rpms"`
	Angle    int               `json:"angle"`
	Temps    [4]float64        `json:"temps"`
	Histos   map[string]string `json:"histos"`
	Webcam   string            `json:"webcam"`
}
