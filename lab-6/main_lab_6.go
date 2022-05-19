package lab_6

import (
	"bufio"
	"bytes"
	"fmt"
	m "github.com/erkkah/margaid"
	table_pac "github.com/jedib0t/go-pretty/table"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
)

var index = 1
var listOfExpressions = []string{
	"y' = y + (1+x)*y^2",
	"y' = -y + (x + 1)^3",
	"y' = 6x^2 + 5y",
}

func MainLab6() {
	// all input here

	println("\nlist of expressions:")
	for i := range listOfExpressions {
		println(i+1, ") ", listOfExpressions[i])
	}
	println()
	print("index of expression: ")
	_, _ = fmt.Scanf("%d", &index)
	print("        lower bound: ")
	_, _ = fmt.Scanf("%g", &a)
	print("        upper bound: ")
	_, _ = fmt.Scanf("%g", &b)
	print("          precision: ")
	_, _ = fmt.Scanf("%g", &precision)
	print("                y_0: ")
	_, _ = fmt.Scanf("%g", &y0)
	print("                  h: ")
	_, _ = fmt.Scanf("%g", &h)

	println()

	//HTML builder
	forHTML()

	tmpl := template.New("tmpl")
	tmpl.Funcs(template.FuncMap{
		"IncludeHTML": IncludeHTML,
	})

	tmpl, err2 := tmpl.Parse(`
	<!DOCTYPE>
	<html>
	<body> <h1 align="center" width="100%"> Summary of all charts </h1>
	
	<div align="center" width="100%">
	<a href="/info" align="center">
	   <input type="button" value="info here"></input>
	</a>
	</div>
	   <div align="center">
			{{ IncludeHTML "./lab-6/resources/graphs/graph1.svg" }}
			{{ IncludeHTML "./lab-6/resources/graphs/graph2.svg" }}
			{{ IncludeHTML "./lab-6/resources/graphs/graph3.svg" }}
		</div> <br>
	
	</body>
	</html>
	   `)
	if err2 != nil {
		log.Fatal(err2)
	}

	// output
	http.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("Error executing template: %v\n", err)
		}
	})

	http.Handle("/info", http.HandlerFunc(infoHandler))

	//fmt.Printf("%+v\n", calculations.GlobalConstants)

	ba, err := readFile("lab-6/resources/outputs/output1.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("%s \n", ba)
	_ = http.ListenAndServe(":9999", nil)
}

func forHTML() {

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	var F *os.File

	F, _ = os.Create("lab-6/resources/outputs/output1.txt")
	os.Stdout = w

	// input calculation here
	drawEuler()
	drawRaw()

	t := table_pac.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table_pac.Row{"#", "x", "y", "Runge", "precise", "<=e"})
	for i := 0; i < n; i++ {
		precise := 0.0
		if index == 1 {
			precise = precise_1(euler[i][0])
		} else if index == 2 {
			precise = precise_2(euler[i][0])
		} else if index == 3 {
			precise = precise_3(euler[i][0])
		}
		t.AppendRows([]table_pac.Row{{
			i,
			math.Round(euler[i][0]*100000) / 100000,
			math.Round(euler[i][1]*100000) / 100000,
			math.Round(((euler[i][1]-euler_half[2*i][1])/(math.Pow(2, 2)-1))*100000) / 100000,
			math.Round((precise)*100000) / 100000,
			math.Abs(((euler[i][1] - euler_half[2*i][1]) / (math.Pow(2, 2) - 1))) < precision}})
	}
	t.Render()

	drawAdams()
	t = table_pac.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table_pac.Row{"#", "x", "y", "runge", "precise", "<=e"})
	for i := 0; i < n; i++ {
		precise := 0.0
		if index == 1 {
			precise = precise_1(adams[i][0])
		} else if index == 2 {
			precise = precise_2(adams[i][0])
		} else if index == 3 {
			precise = precise_3(euler[i][0])
		}
		t.AppendRows([]table_pac.Row{{
			i,
			math.Round(adams[i][0]*100000) / 100000,
			math.Round(adams[i][1]*100000) / 100000,
			math.Round(((adams[i][1]-adams_half[2*i][1])/(math.Pow(2, 4)-1))*100000) / 100000,
			math.Round((precise)*100000) / 100000,
			math.Abs(((adams[i][1] - adams_half[2*i][1]) / (math.Pow(2, 4) - 1))) < precision,
		}})
	}
	t.Render()

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	_, _ = F.Write(out)
}

func infoHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	b, err := ioutil.ReadFile("lab-6/resources/outputs/output1.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	_, err = w.Write(b)
	if err != nil {
		return
	}

}

func IncludeHTML(path string) template.HTML {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("includeHTML - error reading file: %v\n", err)
		return ""
	}

	return template.HTML(b)
}

func readFile(path string) ([]byte, error) {
	parentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	pullPath := filepath.Join(parentPath, path)
	file, err := os.Open(pullPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return read(file)
}

func read(fd_r io.Reader) ([]byte, error) {
	br := bufio.NewReader(fd_r)
	var buf bytes.Buffer

	for {
		ba, isPrefix, err := br.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}

	}
	return buf.Bytes(), nil
}

func getRawSeries() *m.Series {

	rawSeries := m.NewSeries()
	if index == 1 {
		for i := 0; i < n; i++ {
			rawSeries.Add(m.MakeValue(euler[i][0], precise_1(euler[i][0])))
		}
	} else if index == 2 {
		for i := 0; i < n; i++ {
			rawSeries.Add(m.MakeValue(euler[i][0], precise_2(euler[i][0])))
		}
	} else if index == 3 {
		for i := 0; i < n; i++ {
			rawSeries.Add(m.MakeValue(euler[i][0], precise_3(euler[i][0])))
		}
	}

	return rawSeries
}

func f(x float64) float64 {
	return 2 * math.Sin(x)
}
