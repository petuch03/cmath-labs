package lab_1

import (
	"fmt"
	"math"
	"os"
)

var matrixA [][]float64
var matrixB [][]float64
var MatrixX1 [][]float64
var MatrixX2 [][]float64
var isDefinitelyBigger bool
var count = 0

func PrepareMatrixForCalculation(allMatrix [][]float64) {
	matrixA = make([][]float64, Size)
	for i := range matrixA {
		matrixA[i] = make([]float64, Size)
	}
	matrixB = make([][]float64, Size)
	for i := range matrixB {
		matrixB[i] = make([]float64, 1)
	}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			matrixA[i][j] = allMatrix[i][j]
		}
		matrixB[i][0] = allMatrix[i][Size]
	}
	//for i := 0; i < Size; i++ {
	//	fmt.Print(matrixA[i], matrixB[i])
	//	fmt.Println()
	//}
	setDiagonalDominance()
}

func shufflingRows(indexOfVariable int) int {
	var currentIndex = indexOfVariable
	var currentCoefficient float64
	var currentSum float64

	for i := 0; i < Size; i++ {
		currentSum = 0
		currentCoefficient = math.Abs(matrixA[i][currentIndex])
		for j := 0; j < Size; j++ {
			if j != currentIndex {
				currentSum += math.Abs(matrixA[i][j])
			}
		}

		if currentCoefficient >= currentSum {
			if currentCoefficient > currentSum {
				isDefinitelyBigger = true
			}
			swapLines(indexOfVariable, i)
			return 1
		}
	}
	fmt.Println("unable to shuffle matrix in proper way")
	os.Exit(1)
	return 0
}

// переставить все возможные варианты и определяет выполняется ли хоть когда нибудь диагональное соотношение
func setDiagonalDominance() {
	for i := 0; i < Size; i++ {
		shufflingRows(i)
	}
	if !isDefinitelyBigger {
		fmt.Println("no row with '>' strict rule")
		os.Exit(1)
	}
}

func SetResultMatrices() {
	MatrixX1 = make([][]float64, Size)
	MatrixX2 = make([][]float64, Size)
	for i := range MatrixX1 {
		MatrixX1[i] = make([]float64, 1)
	}
	for i := range MatrixX2 {
		MatrixX2[i] = make([]float64, 1)
	}

	for i := 0; i < Size; i++ {
		MatrixX2[i][0] = matrixB[i][0] / matrixA[i][i]
	}
}

func EntryPoint() {

	for true {
		//iteration
		for i := 0; i < Size; i++ {
			MatrixX1[i][0] = MatrixX2[i][0] // перезапись матрицы
		}
		for i := 0; i < Size; i++ {
			sum := 0.0
			for j := 0; j < Size; j++ {
				if j < i {
					sum += matrixA[i][j] * MatrixX2[j][0] / matrixA[i][i] // расчет суммы до
				} else if j != i {
					sum += matrixA[i][j] * MatrixX1[j][0] / matrixA[i][i] // расчет суммы после
				}
			}
			MatrixX2[i][0] = matrixB[i][0]/matrixA[i][i] - sum // вычитание сумм
		}
		count++                             // + итерация
		if checkPrecision() || count >= M { // проверка точности и номера итерации
			break
		}
	}
}

func checkPrecision() bool {
	for i := 0; i < Size; i++ {
		if math.Abs(MatrixX2[i][0]-MatrixX1[i][0]) > Precision {
			return false
		}
	}
	return true
}

// перестановка по индексам
func swapLines(i int, j int) {
	tmp := matrixA[i]
	matrixA[i] = matrixA[j]
	matrixA[j] = tmp

	tmpB := matrixB[i][0]
	matrixB[i][0] = matrixB[j][0]
	matrixB[j][0] = tmpB
}
