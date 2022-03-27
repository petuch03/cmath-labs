package lab_4

import (
	svg "github.com/ajstarks/svgo"
	m "github.com/erkkah/margaid"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func MainLab4() {
	forHTML()

	tmpl := template.New("sample")
	tmpl.Funcs(template.FuncMap{
		"IncludeHTML": IncludeHTML,
	})

	tmpl, err2 := tmpl.Parse(`
<!DOCTYPE>
<html>
<body> <h1> Summary of all svgs and "margaids" </h1> <br>
<a href="/svgo">
   <input type="button" value="show pure svgo" />
</a>
<a href="/margaid">
   <input type="button" value="show pure margaid" />
</a> <br> 
    {{ IncludeHTML "./lab-4/resources/example1.svg" }}
	{{ IncludeHTML "./lab-4/resources/example2.svg" }}
</body>
</html>
    `)
	if err2 != nil {
		log.Fatal(err2)
	}

	http.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("Error executing template: %v\n", err)
		}
	})

	http.Handle("/svgo", http.HandlerFunc(svgo))
	http.Handle("/margaid", http.HandlerFunc(margaid))
	_ = http.ListenAndServe(":9999", nil)

}

func svgo(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "image/svg+xml")
	//s := svg.New(w)
	//s.Start(1350, 500)
	//s.Line(25, 125, 25+200, 125+250, "fill:none;stroke:black")
	//s.Line(25, 125+250, 25+200, 125, "fill:none;stroke:black")
	//
	//s.Line(25+200+50, 125, 25+200+50+100, 125+125, "fill:none;stroke:black")
	//s.Line(25+200+50, 125+250, 25+200+50+200, 125, "fill:none;stroke:black")
	//
	//s.Line(25+200+50+200+50, 125, 25+200+50+200+50, 125+250, "fill:none;stroke:black")
	//s.Line(25+200+50+200+50+200, 125, 25+200+50+200+50+200, 125+250, "fill:none;stroke:black")
	//s.Line(25+200+50+200+50, 125+250, 25+200+50+200+50+200, 125, "fill:none;stroke:black")
	//s.Line(25+200+50+200+50+80, 100, 25+200+50+200+50+200-80, 100, "fill:none;stroke:black")
	//
	//s.Line(250, 550, 375, 675, "fill:none;stroke:black")
	//s.Line(500, 550, 250, 800, "fill:none;stroke:black")
	//s.Circle(250, 250, 125, "fill:none;stroke:black")
	//s.End()

	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}

func margaid(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	randomSeries := m.NewSeries()
	rand.Seed(time.Now().Unix())
	for i := float64(0); i < 10; i++ {
		randomSeries.Add(m.MakeValue(i+1, rand.Float64()-0.5))
	}

	diagram := m.New(800, 800,
		m.WithRange(m.YAxis, -1, 1.5),
		m.WithAutorange(m.XAxis, randomSeries),
		//m.WithInset(70),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Smooth(randomSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Axis(randomSeries, m.XAxis, diagram.ValueTicker('f', 0, 10), false, "X")
	diagram.Axis(randomSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Frame()
	diagram.Title("unix pure margaid sample")

	err2 := diagram.Render(w)
	if err2 != nil {
		return
	}
}

func forHTML() {
	//func forHTML(w http.ResponseWriter, req *http.Request) {
	//	w.Header().Set("Content-Type", "image/svg+xml")

	randomSeries := m.NewSeries()
	rand.Seed(time.Now().UnixNano())
	for i := float64(0); i < 10; i++ {
		randomSeries.Add(m.MakeValue(i+1, rand.Float64()-0.5))
	}

	diagram := m.New(500, 500,
		m.WithRange(m.YAxis, -1, 1.5),
		m.WithAutorange(m.XAxis, randomSeries),
		//m.WithInset(70),
		m.WithPadding(2),
		m.WithColorScheme(90),
	)

	diagram.Smooth(randomSeries, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram.Axis(randomSeries, m.XAxis, diagram.ValueTicker('f', 0, 10), false, "X")
	diagram.Axis(randomSeries, m.YAxis, diagram.ValueTicker('f', 1, 10), true, "Y")

	diagram.Frame()
	diagram.Title("UnixNano")

	randomSeries2 := m.NewSeries()
	rand.Seed(time.Now().Unix())
	for i := float64(0); i < 10; i++ {
		randomSeries2.Add(m.MakeValue(i+1, rand.Float64()))
	}

	diagram2 := m.New(500, 500,
		m.WithRange(m.YAxis, -1, 1.5),
		m.WithAutorange(m.XAxis, randomSeries2),
		//m.WithInset(70),
		m.WithPadding(2),
		m.WithColorScheme(80),
	)

	diagram2.Smooth(randomSeries2, m.UsingAxes(m.XAxis, m.YAxis), m.UsingMarker("filled-circle"))
	diagram2.Axis(randomSeries2, m.XAxis, diagram2.ValueTicker('f', 0, 10), false, "X")
	diagram2.Axis(randomSeries2, m.YAxis, diagram2.ValueTicker('f', 1, 10), true, "Y")

	diagram2.Frame()
	diagram2.Title("unix")

	outfile1, _ := os.Create("./lab-4/resources/example1.svg")
	outfile2, _ := os.Create("./lab-4/resources/example2.svg")

	err := diagram.Render(outfile1)
	if err != nil {
		return
	}
	err2 := diagram2.Render(outfile2)
	if err2 != nil {
		return
	}

}

func IncludeHTML(path string) template.HTML {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("includeHTML - error reading file: %v", err)
		return ""
	}

	return template.HTML(string(b))
}
