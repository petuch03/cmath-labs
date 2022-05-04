package lab_5

import (
	"cmath-labs/lab-5/calculations_lab_5"
	"fmt"
	m "github.com/erkkah/margaid"
	"os"
)

func drawRaw() {
	rawSeries := getRawSeries()

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, rawSeries.MinY(), rawSeries.MaxY()),
		m.WithAutorange(m.XAxis, rawSeries),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("square"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 1, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Frame()
	diagram.Title("raw")

	graph1, _ := os.Create("./lab-5/resources/graphs/graph1.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}

func drawLagrange() {
	fmt.Println("---lagrange---")
	rawSeries := getRawSeries()
	interpolated_1 := calculations_lab_5.LagrangeInterpolation(InputSeries, Size, X_1)
	interpolated_2 := calculations_lab_5.LagrangeInterpolation(InputSeries, Size, X_2)

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, rawSeries.MinY(), rawSeries.MaxY()),
		m.WithAutorange(m.XAxis, rawSeries),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("square"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 1, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Line(interpolated_1, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Line(interpolated_2, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))

	diagram.Frame()
	diagram.Title("lagrange")

	graph1, _ := os.Create("./lab-5/resources/graphs/graph2.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}

func drawNewton() {
	fmt.Println("---newton---")
	rawSeries := getRawSeries()
	interpolated_1 := calculations_lab_5.NewtonInterpolation(InputSeries, Size, X_1)
	interpolated_2 := calculations_lab_5.NewtonInterpolation(InputSeries, Size, X_2)
	all_interpolated := calculations_lab_5.NewtonAllPoints(InputSeries, Size, Size)

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, rawSeries.MinY(), rawSeries.MaxY()),
		m.WithAutorange(m.XAxis, rawSeries),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("square"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 1, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Smooth(all_interpolated, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-square"))
	diagram.Line(interpolated_1, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Line(interpolated_2, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))

	diagram.Frame()
	diagram.Title("newton")

	graph1, _ := os.Create("./lab-5/resources/graphs/graph3.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}
