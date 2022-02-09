package main

import (
	"fmt"
	"math"
	"os"
)

var matrixA [][]float64
var matrixB [][]float64
var matrixX1 [][]float64
var matrixX2 [][]float64
var isDefinitelyBigger bool

func prepareMatrixForCalculation(allMatrix [][]float64) {
	matrixA = make([][]float64, size)
	for i := range matrixA {
		matrixA[i] = make([]float64, size)
	}
	matrixB = make([][]float64, size)
	for i := range matrixB {
		matrixB[i] = make([]float64, 1)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrixA[i][j] = allMatrix[i][j]
		}
		matrixB[i][0] = allMatrix[i][size]
	}
	fmt.Println("before swaps: ")
	for i := 0; i < size; i++ {
		fmt.Print(matrixA[i], matrixB[i])
		fmt.Println()
	}

	setDiagonalDominance()

	fmt.Println("after swaps: ")
	for i := 0; i < size; i++ {
		fmt.Print(matrixA[i], matrixB[i])
		fmt.Println()
	}

}

// перестановка по индексам
func swapLines(i int, j int) {
	var tmp = matrixA[i]
	matrixA[i] = matrixA[j]
	matrixA[j] = tmp

	var tmpB = matrixB[i][0]
	matrixB[i][0] = matrixB[j][0]
	matrixB[j][0] = tmpB
}

func shufflingRows(indexOfVariable int) int {
	var currentIndex = indexOfVariable
	var currentCoefficient float64
	var currentSum float64

	for i := 0; i < size; i++ {
		currentSum = 0
		currentCoefficient = math.Abs(matrixA[i][currentIndex])
		for j := 0; j < size; j++ {
			if j != currentIndex {
				currentSum += math.Abs(matrixA[i][j])
			}
		}

		if currentCoefficient >= currentSum {
			if currentCoefficient > currentSum {
				isDefinitelyBigger = true
			}
			swapLines(indexOfVariable, i)
			return i
		}
	}
	fmt.Println("unable to shuffle matrix in proper way")
	os.Exit(0)
	return 0
}

// переставить все возможные варианты и определяет выполняется ли хоть когда нибудь диагональное соотношение
func setDiagonalDominance() {
	for i := 0; i < size; i++ {
		shufflingRows(i)
	}
	if !isDefinitelyBigger {
		fmt.Println("no row with '>' strict rule")
		os.Exit(0)
	}
}

func initMatrixX1andX2() {
	matrixX1 = make([][]float64, size)
	matrixX2 = make([][]float64, size)
	for i := range matrixX1 {
		matrixX1[i] = make([]float64, 1)
	}
	for i := range matrixX2 {
		matrixX2[i] = make([]float64, 1)
	}

	for i := 0; i < size; i++ {
		matrixX2[i][0] = matrixB[i][0] / matrixA[i][i]
	}
}

func iteration() {
	for i := 0; i < size; i++ {
		matrixX1[i][0] = matrixX2[i][0]
	}
	var sumOther float64
	for i := 0; i < size; i++ {
		sumOther = 0
		for j := 0; j < size; j++ {
			if j < i {
				sumOther += matrixA[i][j] * matrixX2[j][0] / matrixA[i][i]
			} else if j == i {
			} else {
				sumOther += matrixA[i][j] * matrixX1[j][0] / matrixA[i][i]
			}
		}
		matrixX2[i][0] = matrixB[i][0]/matrixA[i][i] - sumOther
	}
}

func checkAllNewX() bool {
	for i := 0; i < size; i++ {
		if math.Abs(matrixX2[i][0]-matrixX1[i][0]) > precision {
			return false
		}
	}
	return true
}

func startComputed() {
	count := 0

	for true {
		iteration()
		count++
		if checkAllNewX() || count >= M {
			break
		}
	}

	fmt.Println("after all iterations")
	for i := 0; i < size; i++ {
		fmt.Printf("x%d=%f \n", i+1, matrixX2[i][0])
	}

	if count >= M {
		fmt.Println("does not converge")
	} else {
		fmt.Println("converges at", count)
	}

	for i := 0; i < size; i++ {
		fmt.Printf("vector Х_%d=%f \n", i+1, math.Abs(matrixX2[i][0]-matrixX1[i][0]))
	}
}
