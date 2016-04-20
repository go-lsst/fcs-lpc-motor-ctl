// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
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

type monData struct {
	id       time.Time
	rotation int
	rpms     uint32
	angle    float64
	temps    [4]float64
}

func (mon monData) x() float64 {
	return float64(mon.id.Unix())
}

type monRotation []monData

func (mon monRotation) Len() int { return len(mon) }
func (mon monRotation) XY(i int) (float64, float64) {
	v := mon[i]
	return v.x(), float64(v.rotation)
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

func (srv *server) makeMonPlots() map[string]string {
	plots := make(map[string]string, 3)

	// temperature
	{
		p, err := newPlot("", "T (°C)",
			monTemps{0, srv.histos.rows}, monTemps{1, srv.histos.rows},
			monTemps{2, srv.histos.rows}, monTemps{3, srv.histos.rows},
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
			monAngle(srv.histos.rows),
		)
		if err != nil {
			panic(err)
		}
		plots["position"] = renderSVG(p)
	}

	// RPMs
	{
		p, err := newPlot("", "RPMs", monRPMs(srv.histos.rows))
		if err != nil {
			panic(err)
		}
		plots["rpms"] = renderSVG(p)
	}
	return plots
}
