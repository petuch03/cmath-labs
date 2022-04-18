package lab_4

import (
	"bufio"
	"bytes"
	"cmath-labs/lab-4/calculations"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
)

var InputSeries [][]float64
var Size = 11
var inputType string
var outputType string

func MainLab4() {
	// all calculations here
	//fmt.Print("print 'file' to set stdout to file (otherwise press enter): ")
	//_, _ = fmt.Scanf("%s", &outputType)
	fmt.Print("result will be at localhost:9999/main\n")
	fmt.Print("print 'input?' or 'console' to set input type: ")
	_, _ = fmt.Scanf("%s", &inputType)
	if inputType == "input1" {
		f, _ := os.Open("lab-4/resources/inputs/input1.txt")
		oldStdin := os.Stdin
		os.Stdin = f

		customInput()

		_ = f.Close()
		os.Stdin = oldStdin
	} else if inputType == "input2" {
		f, _ := os.Open("lab-4/resources/inputs/input2.txt")
		oldStdin := os.Stdin
		os.Stdin = f

		customInput()

		_ = f.Close()
		os.Stdin = oldStdin
	} else if inputType == "input0" {
		f, _ := os.Open("lab-4/resources/inputs/input0.txt")
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
	sort.Slice(InputSeries, func(i, j int) bool {
		return InputSeries[i][0] < InputSeries[j][0]
	})
	calculations.FillConstants(InputSeries)

	// HTML builder
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

	http.Handle("/info", http.HandlerFunc(infoHandler))

	//fmt.Printf("%+v\n", calculations.GlobalConstants)

	ba, err := readFile("lab-4/resources/outputs/output1.txt")
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

	F, _ = os.Create("lab-4/resources/outputs/output1.txt")
	os.Stdout = w

	drawLinear()
	drawQuad()
	drawCub()
	drawExp()
	drawPow()
	drawLog()

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	_, _ = F.Write(out)
}

func infoHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	b, err := ioutil.ReadFile("lab-4/resources/outputs/output1.txt") // just pass the file name
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
		log.Println("includeHTML - error reading file: %v", err)
		return ""
	}

	return template.HTML(string(b))
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

	defer file.Close()
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
