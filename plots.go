// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image/color"
	"image/color/palette"
	"math"
	"time"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
	"github.com/gonum/plot/vg/vgsvg"
)

var (
	plotColors []color.Color
)

func init() {
	plotColors = []color.Color{
		color.NRGBA{255, 0, 0, 128},
		color.NRGBA{0, 255, 0, 128},
		color.NRGBA{0, 0, 255, 128},
	}
	plotColors = append(plotColors, palette.Plan9...)
}

func newPlot(title, yaxis string, data ...plotter.XYer) (*plot.Plot, error) {
	p, err := plot.New()
	if err != nil {
		return nil, err
	}

	p.Title.Text = title
	p.Y.Label.Text = yaxis
	p.X.Tick.Marker = plot.UnixTimeTicks{Format: "2006-01-02\n15:04:05"}

	for i, v := range data {
		lines, points, err := plotter.NewLinePoints(v)
		if err != nil {
			return nil, err
		}
		c := plotColors[i]
		points.Color = c
		lines.Color = c
		p.Add(points, lines)
	}
	p.Add(plotter.NewGrid())

	return p, nil
}

type motorMode byte

const (
	motorModeDefault motorMode = iota
	motorModeReady
	motorModeHWSafety
	motorModeSTO
	motorModeHome
	motorModeRandom
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
	case motorModeReady:
		return "ready"
	case motorModeHWSafety:
		return "h/w safety"
	case motorModeSTO:
		return "sto"
	case motorModeHome:
		return "home"
	case motorModeRandom:
		return "random"
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

type monRPMs []monData

func (mon monRPMs) Len() int { return len(mon) }
func (mon monRPMs) XY(i int) (float64, float64) {
	v := mon[i]
	return v.x(), float64(v.rpms)
}

type monTemps struct {
	t    int
	data []monData
}

func (mon monTemps) Len() int { return len(mon.data) }
func (mon monTemps) XY(i int) (float64, float64) {
	v := mon.data[i]
	return v.x(), float64(v.temps[mon.t])
}

type monAngle []monData

func (mon monAngle) Len() int { return len(mon) }
func (mon monAngle) XY(i int) (float64, float64) {
	v := mon[i]
	return v.x(), v.angle
}

func renderSVG(p *plot.Plot) string {
	size := 10 * vg.Centimeter
	canvas := vgsvg.New(size, size/vg.Length(math.Phi))
	p.Draw(draw.New(canvas))
	out := new(bytes.Buffer)
	_, err := canvas.WriteTo(out)
	if err != nil {
		panic(err)
	}
	return string(out.Bytes())
}

func (srv *server) makeMonPlots(i int) map[string]string {
	plots := make(map[string]string, 3)
	motor := srv.motors()[i]

	// temperature
	{
		p, err := newPlot("", "T (°C)",
			monTemps{0, motor.histos.rows}, monTemps{1, motor.histos.rows},
			monTemps{2, motor.histos.rows}, monTemps{3, motor.histos.rows},
		)
		if err != nil {
			panic(err)
		}

		plots["temperature"] = renderSVG(p)
	}

	// angular position
	{
		p, err := newPlot(
			"", "Angular Position",
			monAngle(motor.histos.rows),
		)
		if err != nil {
			panic(err)
		}
		plots["position"] = renderSVG(p)
	}

	// RPMs
	{
		p, err := newPlot("", "RPMs", monRPMs(motor.histos.rows))
		if err != nil {
			panic(err)
		}
		plots["rpms"] = renderSVG(p)
	}
	return plots
}
