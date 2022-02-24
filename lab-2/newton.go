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

func entryPointNewton(a float64, b float64) {
	startX = a
	startY = b
	for true {
		counter++
		//"system: {x^2-3y=0, y^2-2x=0}",
		matrix = make([][]float64, lab1.Size)
		for i := range matrix {
			matrix[i] = make([]float64, lab1.Size+1)
		}

		matrix[0][0] = 2 * startX
		matrix[0][1] = -3
		matrix[0][2] = 3*startY - startX*startX
		matrix[1][0] = -2
		matrix[1][1] = 2 * startY
		matrix[1][2] = 2*startX - startY*startY

		lab1.Precision = precision
		lab1.PrepareMatrixForCalculation(matrix)
		lab1.SetResultMatrices()
		lab1.EntryPoint()
		secondX = startX + lab1.MatrixX2[0][0]
		secondY = startY + lab1.MatrixX2[1][0]
		if checkNewton() || counter > 20 {
			break
		} else {
			startX = secondX
			startY = secondY
		}

	}
}

func checkNewton() bool {
	return math.Abs(secondX-startX) <= precision || math.Abs(secondY-startY) <= precision
}
