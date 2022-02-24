package lab_2

import (
	lab1 "cmath-labs/lab-1"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/jedib0t/go-pretty/table"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

var lowerBound float64
var upperBound float64
var currentA float64
var currentB float64
var precision float64 = 0.001
var inputType string
var outputType string
var index = 1
var listOfExpressions = []string{
	"x*x*x - x + 4",
	"1.8*x*x*x - 2.47*x*x - 5.53*x + 1.539",
	"sin_x",
	"sin_x + cos_x",
	"system: {x^2-3y=0, y^2-2x=0}",
}
var currentExpression = listOfExpressions[index-1]
var expressions = map[string]string{
	"x*x*x - x + 4":                         "3*x*x - 1",
	"1.8*x*x*x - 2.47*x*x - 5.53*x + 1.539": "1.8*3*x*x - 2.47*2*x - 5.53",
	"sin_x":                                 "cos_x",
	"sin_x + cos_x":                         "cos_x - sin_x",
}

func MainLab2() {
	fmt.Print("\nlist of expressions: \n")
	for i := range listOfExpressions {
		fmt.Println(i+1, ") ", listOfExpressions[i])
	}
	fmt.Print("\nprint 'file' or 'console' to set input type: ")
	_, _ = fmt.Scanf("%s", &inputType)
	if inputType == "file" {
		f, _ := os.Open("lab-2/resources/input.txt")
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = f
	} else if inputType == "console" {
		fmt.Printf("\ninput the name of method, index of expression you want (1-5), lower bound (initial for x), upper bound (initial for y), precision and output type in proper order: ")
	} else {
		fmt.Printf("invalid type")
		os.Exit(1)
	}
	var method string
	_, _ = fmt.Scanf("%s %d %g %g %g %s", &method, &index, &lowerBound, &upperBound, &precision, &outputType)
	currentExpression = listOfExpressions[index-1]

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	var F *os.File
	if outputType == "file" {
		F, _ = os.Create("lab-2/resources/output.txt")
		os.Stdout = w
	}
	t := table.NewWriter()
	start := time.Now()
	if method == "iteration" {
		entryPointIteration()
		tableTemplateIteration(t)
		fmt.Printf("iterations: %d\n", len(xk)-1)
		fmt.Printf("x= %f\n", currentXk)
		fmt.Printf("f(x) = %f\n", currentFxk)
		drawPlot()
	} else if method == "half" {
		entryPointHalf()
		tableTemplateHalf(t)
		fmt.Printf("iterations: %d\n", len(x)-1)
		fmt.Printf("x = %f\n", currentX)
		fmt.Printf("f(x) = %f\n", currentFx)
		drawPlot()
	} else if method == "newton" {
		system()
	}
	duration := time.Since(start)
	fmt.Println("----", duration, "----")
	PrintMemUsage()

	if outputType == "file" {
		_ = w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		_, _ = F.Write(out)
	}

	debug.FreeOSMemory()
}

func function(x float64) float64 {
	expression, _ := govaluate.NewEvaluableExpression(currentExpression)

	parameters := make(map[string]interface{}, 8)
	parameters["x"] = x
	parameters["sin_x"] = math.Sin(x)
	parameters["cos_x"] = math.Cos(x)
	result, _ := expression.Evaluate(parameters)

	switch i := result.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	default:
		os.Exit(1)
		return -1
	}
}

func derivative(x float64) float64 {
	derivative, _ := govaluate.NewEvaluableExpression(expressions[currentExpression])
	//if currentExpression == listOfExpressions[2] {
	//	return math.Cos(x)
	//}

	parameters := make(map[string]interface{}, 8)
	parameters["x"] = x
	parameters["sin_x"] = math.Sin(x)
	parameters["cos_x"] = math.Cos(x)
	result, _ := derivative.Evaluate(parameters)

	switch i := result.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	default:
		os.Exit(1)
		return -1
	}
}

func tableTemplateHalf(t table.Writer) {
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "a", "b", "x", "f(a)", "f(b)", "f(x)", "|a-b|"})
	for i := 0; i < len(x); i++ {
		t.AppendRows([]table.Row{{
			i,
			math.Round(a[i].value*1000) / 1000,
			math.Round(b[i].value*1000) / 1000,
			math.Round(x[i].value*1000) / 1000,
			math.Round(fa[i].value*1000) / 1000,
			math.Round(fb[i].value*1000) / 1000,
			math.Round(fx[i].value*1000) / 1000,
			math.Round(mod[i].value*1000) / 1000}})
	}
	t.Render()
}

func tableTemplateIteration(t table.Writer) {
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "x_k", "f(x_k)", "x_k+1", "phi(x_k)", "|x_k-x_k+1|"})
	for i := 0; i < len(xk); i++ {
		t.AppendRows([]table.Row{{
			i,
			math.Round(xk[i].value*1000) / 1000,
			math.Round(fxk[i].value*1000) / 1000,
			math.Round(xk1[i].value*1000) / 1000,
			math.Round(phixk[i].value*1000) / 1000,
			math.Round(modxk[i].value*1000) / 1000}})
	}
	t.Render()
}

func drawPlot() {
	p := plot.New()

	p.Title.Text = "Function"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	f := plotter.NewFunction(func(x float64) float64 { return function(x) })
	f.Color = color.RGBA{B: 255, A: 255}

	der := plotter.NewFunction(func(x float64) float64 { return derivative(x) })
	der.Color = color.RGBA{G: 255, A: 255}

	xAxis := plotter.NewFunction(func(x float64) float64 { return 0 })
	xAxis.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	xAxis.Width = vg.Points(1.5)
	xAxis.Color = color.RGBA{A: 255}

	p.Add(f, xAxis, der)
	p.Legend.Add("function", f)
	p.Legend.Add("derivative", der)
	p.Legend.ThumbnailWidth = 1 * vg.Inch

	p.X.Min = lowerBound - 2
	p.X.Max = upperBound + 2
	p.Y.Min = -5
	p.Y.Max = 5

	if err := p.Save(7*vg.Inch, 7*vg.Inch, "lab-2/resources/function.png"); err != nil {
		log.Fatal(err)
	}
}

func system() {
	lab1.Size = 2
	entryPointNewton(lowerBound, upperBound)
	fmt.Printf("iterations: %d\n", counter)
	if counter >= 50 {
		println("iteration limit stopped the program")
	}
	fmt.Printf("x_1 = %e\nx_2 = %e\n", secondX, secondY)
	fmt.Println("---error vector---")
	for i := 0; i < lab1.Size; i++ {
		fmt.Printf("x_%d = %e \n", i+1, math.Abs(lab1.MatrixX2[i][0]-lab1.MatrixX1[i][0]))
	}

	p := plot.New()

	p.Title.Text = "Function"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	g := plotter.NewFunction(func(x float64) float64 {
		if x < 0 {
			return 0
		}
		return math.Sqrt(2 * x)
	})
	g.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	g.Width = vg.Points(1)
	g.Color = color.RGBA{R: 189, G: 155, B: 25, A: 255}
	g1 := plotter.NewFunction(func(x float64) float64 {
		if x < 0 {
			return 0
		}
		return -math.Sqrt(2 * x)
	})
	g1.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	g1.Width = vg.Points(1)
	g1.Color = color.RGBA{R: 189, G: 155, B: 25, A: 255}

	f := plotter.NewFunction(func(x float64) float64 { return x * x / 3 })
	f.Color = color.RGBA{G: 255, A: 255}

	xAxis := plotter.NewFunction(func(x float64) float64 { return 0 })
	xAxis.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	xAxis.Width = vg.Points(1.5)
	xAxis.Color = color.RGBA{A: 255}

	p.Add(f, g1, xAxis, g)
	p.Legend.Add("f", f)
	p.Legend.Add("g", g)
	p.Legend.ThumbnailWidth = 1 * vg.Inch

	p.X.Min = -8
	p.X.Max = 8
	p.Y.Min = -8
	p.Y.Max = 8

	if err := p.Save(7*vg.Inch, 7*vg.Inch, "lab-2/resources/function.png"); err != nil {
		log.Fatal(err)
	}
}

// 1.8*x*x*x - 2.47*x*x - 5.53*x + 1.539
// (-2, -1), (0, 1), (2, 3)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v KB\n", btoKB(float64(m.Alloc)))
	fmt.Printf("TotalAlloc = %v KB\n", btoKB(float64(m.TotalAlloc)))
	fmt.Printf("Sys = %v KB\n", btoKB(float64(m.Sys)))
}

func btoKB(b float64) float64 {
	return b / 1024 / 1024 / 1024
}
