package lab_4

import (
	"cmath-labs/lab-4/calculations"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var InputSeries [][]float64
var Size = 11
var inputType string

func MainLab4() {
	// all calculations here
	fmt.Print("print 'input?' or 'console' to set input type: ")
	_, _ = fmt.Scanf("%s", &inputType)
	if inputType == "input1" {
		f, _ := os.Open("lab-4/resources/inputs/input1.txt")
		oldStdin := os.Stdin
		os.Stdin = f

		customInput()

		_ = f.Close()
		os.Stdin = oldStdin
	} else if inputType == "console" {
		customInput()
	} else {
		fmt.Print("unsupported input type: ", inputType, "\n")
		fmt.Print("initiated default function approximation...\n")
		input()
	}
	calculations.FillConstants(InputSeries)

	// HTML builder
	forHTML()

	tmpl := template.New("sample")
	tmpl.Funcs(template.FuncMap{
		"IncludeHTML": IncludeHTML,
	})

	tmpl, err2 := tmpl.Parse(`
<!DOCTYPE>
<html>
<body> <h1 align="center"> Summary of all charts </h1> <br>
<br> 
    <div align="center">
		{{ IncludeHTML "./lab-4/resources/graphs/graph1.svg" }}
		{{ IncludeHTML "./lab-4/resources/graphs/graph2.svg" }}
		{{ IncludeHTML "./lab-4/resources/graphs/graph3.svg" }} 
	</div> <br>
	<div align="center"> 
		{{ IncludeHTML "./lab-4/resources/graphs/graph5.svg" }}
		{{ IncludeHTML "./lab-4/resources/graphs/graph4.svg" }} 
		{{ IncludeHTML "./lab-4/resources/graphs/graph6.svg" }} 
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

	//fmt.Printf("%+v\n", calculations.GlobalConstants)

	_ = http.ListenAndServe(":9999", nil)
}

func forHTML() {
	drawLinear()
	drawPow()
	drawExp()
	drawCub()
	drawLog()
	drawQuad()
}

func IncludeHTML(path string) template.HTML {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("includeHTML - error reading file: %v", err)
		return ""
	}

	return template.HTML(string(b))
}
