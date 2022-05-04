package calculations_lab_5

import (
	"fmt"
	m "github.com/erkkah/margaid"
	"math"
)

var table [][]float64

func LagrangeInterpolation(inputSeries [][]float64, size int, x float64) *m.Series {
	var res_y float64 = 0
	var up float64 = 1
	var down float64 = 1

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j != i {
				up *= x - inputSeries[j][0]
				down *= inputSeries[i][0] - inputSeries[j][0]
				//fmt.Printf("up *= (%f - %f)\n", x, inputSeries[j][0])
				//fmt.Printf("down *= (%f - %f)\n\n", inputSeries[i][0], inputSeries[j][0])
			}
		}
		res_y += (up / down) * inputSeries[i][1]
		up = 1
		down = 1
		//fmt.Printf("res_y += (%f) * %f\n", up/down, inputSeries[i][1])
		//fmt.Printf("res_y = %f\n==========\n", res_y)
	}

	fmt.Printf("L(%f) = %f\n", x, res_y)

	interpolatedSeries := m.NewSeries()
	interpolatedSeries.Add(m.MakeValue(x, res_y))
	interpolatedSeries.Add(m.MakeValue(x, res_y))
	return interpolatedSeries
}

func NewtonInterpolation(inputSeries [][]float64, size int, x float64) *m.Series {
	counter := 0
	for i := 0; i < size; i++ {
		if inputSeries[i][0] > x {
			counter = i
			break
		}
	}
	betweenRight := 0
	betweenLeft := 0
	if counter > 0 {
		betweenRight = counter
		betweenLeft = counter - 1
	} else {
		betweenRight = counter + 1
		betweenLeft = counter
	}

	tableSize := size - 1
	t := 0.0
	coef := 1.0
	res_y := 0.0
	h := inputSeries[1][0] - inputSeries[0][0]
	table = make([][]float64, tableSize)
	for i := 0; i < tableSize; i++ {
		table[i] = make([]float64, tableSize-i)
	}

	for columns := 0; columns < tableSize; columns++ {
		for raws := 0; raws < tableSize-columns; raws++ {
			if columns == 0 {
				table[columns][raws] = inputSeries[raws+1][1] - inputSeries[raws][1]
			} else {
				table[columns][raws] = table[columns-1][raws+1] - table[columns-1][raws]
			}
		}
	}
	// if in left part
	if math.Abs(inputSeries[betweenLeft][0]-x) <= math.Abs(inputSeries[betweenRight][0]-x) {
		t = (x - inputSeries[betweenLeft][0]) / h

		res_y += inputSeries[betweenLeft][1] * coef
		coef = t
		res_y += table[0][betweenLeft] * coef
		coef = 1
		for i := 1; i < tableSize-betweenLeft; i++ {
			for j := 0; j <= i; j++ {
				coef *= t - float64(j)
			}
			coef /= float64(Factorial(i + 1))
			res_y += table[i][betweenLeft] * coef
			coef = 1
		}
	} else {
		t = (x - inputSeries[betweenRight][0]) / h

		res_y += inputSeries[betweenRight][1] * coef
		coef = t
		res_y += table[0][betweenRight-1] * coef
		coef = 1
		for i := 1; i < betweenRight; i++ {
			for j := 0; j <= i; j++ {
				coef *= t + float64(j)
			}
			coef /= float64(Factorial(i + 1))
			res_y += table[i][betweenRight-i-1] * coef
			coef = 1
		}
	}

	fmt.Printf("N(%f) = %f\n", x, res_y)

	interpolatedSeries := m.NewSeries()
	interpolatedSeries.Add(m.MakeValue(x, res_y))
	interpolatedSeries.Add(m.MakeValue(x, res_y))
	return interpolatedSeries

}

func NewtonAllPoints(inputSeries [][]float64, size int, n int) *m.Series {
	outputSeries := make([][]float64, size)
	for i := 0; i < size; i++ {
		outputSeries[i] = make([]float64, size)
	}
	h := inputSeries[1][0] - inputSeries[0][0]

	for i := 0; i < size; i++ {
		x := inputSeries[i][0]
		res_y := inputSeries[0][1]
		tmp := 1.0

		for j := 1; j < n; j++ {
			for k := 0; k < j; k++ {
				tmp *= x - inputSeries[k][0]
				tmp /= h
			}
			tmp *= table[j-1][0]
			tmp /= float64(Factorial(j))
			res_y += tmp
			tmp = 1
		}

		outputSeries[i][0] = x
		outputSeries[i][1] = res_y
	}

	interpolatedSeries := m.NewSeries()
	for i := 0; i < size; i++ {
		interpolatedSeries.Add(m.MakeValue(outputSeries[i][0], outputSeries[i][1]))
	}
	return interpolatedSeries
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
