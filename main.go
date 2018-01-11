// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"sort"
	"strconv"

	"gonum.org/v1/plot/vg"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

// Example_timeSeries draws a time series.
func Draw(values [][]float64) {

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	computeUseFullValue(values[0])
	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.Add(plotter.NewGrid())
	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	points := func(i int) plotter.XYs {
		const (
			month = 1
			day   = 1
			hour  = 1
			min   = 1
			sec   = 1
			nsec  = 1
		)
		values := readData()
		pts := make(plotter.XYs, len(values[i]))
		xs := values[i]
		for j := range pts {
			pts[j].X = float64(j)
			pts[j].Y = xs[j]
		}
		return pts
	}

	l, err := plotter.NewLine(points(0))

	if err != nil {
		log.Printf("error is %s", err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(l)

	l1, err := plotter.NewLine(points(1))

	if err != nil {
		log.Printf("error is %s", err)
	}
	l1.LineStyle.Width = vg.Points(1)
	l1.LineStyle.Color = color.RGBA{B: 0, A: 255}

	p.Add(l1)

	if err := p.Save(10*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func readData() [][]float64 {

	file, err := os.Open("vitesse.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()
	//
	reader := csv.NewReader(file)

	reader.Comma = ';'
	result, err := reader.ReadAll()
	if err != nil {
		log.Printf("error is %s", err)
	}

	var points [][]float64
	f := func(result [][]string, i int) []float64 {
		var pts []float64
		for j := 0; j < len(result); j++ {

			v, err := strconv.ParseFloat(result[j][i], 64)
			if err != nil {
				log.Printf("error is %s", err)
				break
			}
			if v != 0 {
				pts = append(pts, v)
			}
		}
		return pts
	}
	points = append(points, f(result, 1))
	points = append(points, f(result, 2))
	return points
}

func computeUseFullValue(serie []float64){
	sort.Float64s(serie)

	mean := stat.Mean(serie, nil)
	median := stat.Quantile(0.5, stat.Empirical, serie, nil)
	variance := stat.Variance(serie, nil)
	stddev := math.Sqrt(variance)

	fmt.Printf("mean=     %v\n", mean)
	fmt.Printf("median=   %v\n", median)
	fmt.Printf("variance= %v\n", variance)
	fmt.Printf("std-devpoint=   %v\n", stddev)
}
func main() {


	
	Draw(readData())
}
