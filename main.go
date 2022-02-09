package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

var inputType string
var size int
var precision float64
var matrix [][]float64
var M int

func main() {
	fmt.Print("print file to file input and print console to console input: ")
	_, _ = fmt.Scanf("%s", &inputType)
	if inputType == "file" {
		f, _ := os.Open("resources/input.txt")
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = f
	} else if inputType == "console" {

	} else {
		os.Exit(1)
	}
	_, _ = fmt.Scanf("%d %d %f", &size, &M, &precision)
	matrix = make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size+1)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size+1; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Println()
	prepareMatrixForCalculation(matrix)
	initMatrixX1andX2()
	startComputed()
	matrix = nil
	debug.FreeOSMemory()
}
