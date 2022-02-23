package lab_1

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime/debug"
)

var inputType string
var Size int
var Precision float64
var matrix [][]float64
var M int

func MainLab1() {
	fmt.Print("print 'input?' or 'console' to set input type: ")
	_, _ = fmt.Scanf("%s", &inputType)
	if inputType == "input1" {
		f, _ := os.Open("lab-1/resources/input1.txt")
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = f
	} else if inputType == "input2" {
		f, _ := os.Open("lab-1/resources/input2.txt")
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = f
	} else if inputType == "input3" {
		f, _ := os.Open("lab-1/resources/input3.txt")
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = f
	} else if inputType == "f4" {
		f, _ := os.Open("lab-1/resources/f4.txt")
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = f
	} else if inputType == "input4" {
		f, _ := os.Open("lab-1/resources/input4.txt")
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = f
	} else if inputType == "console" {
	} else {
		fmt.Print("unsupported input type: ", inputType)
		os.Exit(1)
	}
	_, _ = fmt.Scanf("%d %d %f", &Size, &M, &Precision)
	matrix = make([][]float64, Size)
	for i := range matrix {
		matrix[i] = make([]float64, Size+1)
	}

	for i := 0; i < Size; i++ {
		for j := 0; j < Size+1; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Println()
	PrepareMatrixForCalculation(matrix)
	SetResultMatrices()
	EntryPoint()
	fmt.Println("-----result vector-----")
	for i := 0; i < Size; i++ {
		fmt.Printf("x%d=%e \n", i+1, MatrixX2[i][0])
	}

	fmt.Println("\n-----converges?-----")
	if count >= M {
		fmt.Println("no")
	} else {
		fmt.Println("yes, converges at", count)
	}

	fmt.Println("\n-----error vector-----")
	for i := 0; i < Size; i++ {
		fmt.Printf("x%d=%e \n", i+1, math.Abs(MatrixX2[i][0]-matrixX1[i][0]))
	}
	matrix = nil
	debug.FreeOSMemory()
}

func generateN(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i < j {
				fmt.Print(rand.Intn(2), " ")
			} else if i == j {
				fmt.Print(rand.Intn(2)+100, " ")
			} else if i > j {
				fmt.Print(rand.Intn(2), " ")
			}
		}
		fmt.Print(rand.Intn(5), " ")
		fmt.Println()
	}
}
