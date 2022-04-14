package lab_4

import (
	"cmath-labs/lab-4/calculations"
	m "github.com/erkkah/margaid"
	"math"
	"os"
)

func drawLinear() {
	rawSeries := getRawSeries()
	approximated := calculations.LinearApproximation(InputSeries, Size)

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, math.Min(rawSeries.MinY(), approximated.MinY()), math.Max(rawSeries.MaxY(), approximated.MaxY())),
		m.WithAutorange(m.XAxis, rawSeries),
		//m.WithInset(70),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 0, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Line(approximated, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-square"))

	diagram.Frame()
	diagram.Title("linear")

	graph1, _ := os.Create("./lab-4/resources/graphs/graph1.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}

func drawExp() {
	rawSeries := getRawSeries()
	approximated := calculations.ExponentApproximation(InputSeries, Size)

	diagram := m.New(400, 400,
		//m.WithRange(m.YAxis, math.Min(rawSeries.MinY(), approximated.MinY()), math.Max(rawSeries.MaxY(), rawSeries.MaxY())),
		m.WithRange(m.YAxis, math.Min(rawSeries.MinY(), approximated.MinY()), math.Max(rawSeries.MaxY(), approximated.MaxY())),
		m.WithAutorange(m.XAxis, rawSeries),
		//m.WithRange(m.XAxis, math.Min(rawSeries.MinX(), approximated.MinX()), math.Max(rawSeries.MaxX(), approximated.MaxX())),

		//m.WithInset(70),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 0, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Line(approximated, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-square"))

	diagram.Frame()
	diagram.Title("exp")

	graph1, _ := os.Create("./lab-4/resources/graphs/graph5.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}

func drawLog() {
	rawSeries := getRawSeries()
	approximated := calculations.LogApproximation(InputSeries, Size)

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, math.Min(rawSeries.MinY(), approximated.MinY()), math.Max(rawSeries.MaxY(), approximated.MaxY())),
		//m.WithRange(m.YAxis, rawSeries.MinY()-1, rawSeries.MaxY()+1),
		m.WithAutorange(m.XAxis, rawSeries),
		//m.WithInset(70),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 0, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Line(approximated, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-square"))

	diagram.Frame()
	diagram.Title("log")

	graph1, _ := os.Create("./lab-4/resources/graphs/graph6.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}

func drawQuad() {
	rawSeries := getRawSeries()
	approximated := calculations.QuadraticApproximation(InputSeries, Size)

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, math.Min(rawSeries.MinY(), approximated.MinY()), math.Max(rawSeries.MaxY(), approximated.MaxY())),
		//m.WithRange(m.YAxis, rawSeries.MinY()-1, rawSeries.MaxY()+1),
		m.WithAutorange(m.XAxis, rawSeries),
		//m.WithInset(70),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 0, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Line(approximated, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-square"))

	diagram.Frame()
	diagram.Title("quadratic")

	graph1, _ := os.Create("./lab-4/resources/graphs/graph2.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}

func drawCub() {
	rawSeries := getRawSeries()
	approximated := calculations.CubicApproximation(InputSeries, Size)

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, math.Min(rawSeries.MinY(), approximated.MinY()), math.Max(rawSeries.MaxY(), approximated.MaxY())),
		//m.WithRange(m.YAxis, rawSeries.MinY()-1, rawSeries.MaxY()+1),
		m.WithAutorange(m.XAxis, rawSeries),
		//m.WithInset(70),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 0, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Line(approximated, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-square"))

	diagram.Frame()
	diagram.Title("cubic")

	graph1, _ := os.Create("./lab-4/resources/graphs/graph3.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}

func drawPow() {
	rawSeries := getRawSeries()
	approximated := calculations.PowApproximation(InputSeries, Size)

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, math.Min(rawSeries.MinY(), approximated.MinY()), math.Max(rawSeries.MaxY(), approximated.MaxY())),
		//m.WithRange(m.YAxis, rawSeries.MinY()-1, rawSeries.MaxY()+1),
		m.WithAutorange(m.XAxis, rawSeries),
		//m.WithInset(70),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 0, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Line(approximated, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-square"))

	diagram.Frame()
	diagram.Title("pow")

	graph1, _ := os.Create("./lab-4/resources/graphs/graph4.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}
