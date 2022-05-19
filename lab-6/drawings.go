package lab_6

import (
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

	graph1, _ := os.Create("./lab-6/resources/graphs/graph1.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}

func drawEuler() {
	fmt.Println("---euler---")
	euler := EulerMethod()
	rawSeries := getRawSeries()

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, rawSeries.MinY(), euler.MaxY()),
		m.WithAutorange(m.XAxis, rawSeries),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("square"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 1, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Line(euler, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))

	diagram.Frame()
	diagram.Title("euler")

	graph1, _ := os.Create("./lab-6/resources/graphs/graph2.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}

func drawAdams() {
	fmt.Println("---adams---")
	adams := AdamsMethod()
	rawSeries := getRawSeries()

	diagram := m.New(400, 400,
		m.WithRange(m.YAxis, rawSeries.MinY(), adams.MaxY()),
		m.WithAutorange(m.XAxis, rawSeries),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Line(rawSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("square"))
	diagram.Axis(rawSeries, m.XAxis, diagram.ValueTicker('f', 1, 10), false, "X")
	diagram.Axis(rawSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Line(adams, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))

	diagram.Frame()
	diagram.Title("adams")

	graph1, _ := os.Create("./lab-6/resources/graphs/graph3.svg")
	err := diagram.Render(graph1)
	if err != nil {
		return
	}
}
