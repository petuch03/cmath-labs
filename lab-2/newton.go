package lab_2

import (
	lab1 "cmath-labs/lab-1"
	"math"
)

var startX float64
var secondX float64
var startY float64
var secondY float64
var deltaX float64
var deltaY float64
var matrix [][]float64
var counter = 0

func initSolve(a float64, b float64) {
	startX = a
	startY = b
	for true {
		counter++

		matrix = make([][]float64, lab1.Size)
		for i := range matrix {
			matrix[i] = make([]float64, lab1.Size+1)
		}

		matrix[0][0] = 2 * startX
		matrix[0][1] = -2
		matrix[0][2] = 2*startY - startX*startX
		matrix[1][0] = 3
		matrix[1][1] = -2 * startY
		matrix[1][2] = -3 + startY*startY - 3*startX
		lab1.Precision = precision
		lab1.PrepareMatrixForCalculation(matrix)
		lab1.SetResultMatrices()
		lab1.EntryPoint()
		deltaX = lab1.MatrixX2[0][0]
		deltaY = lab1.MatrixX2[1][0]

		secondX = startX + deltaX
		secondY = startY + deltaY
		if check() || counter > 20 {
			break
		}

		startX = secondX
		startY = secondY
	}
}

func check() bool {
	return math.Abs(secondX-startX) <= precision || math.Abs(secondY-startY) <= precision
}
