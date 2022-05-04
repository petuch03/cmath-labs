package calculations_lab_4

import (
	"fmt"
	"math"
)

type matrix [][]float64
type Vector []float64

func CalculateMatrix(custom matrix, lines int) (Vector, error, error) {

	columns := len(custom[0]) - 1
	a := make(matrix, lines)
	for i := range a {
		a[i] = make([]float64, columns)
	}

	for i := 0; i < lines; i++ {
		for j := 0; j < columns; j++ {
			a[i][j] = custom[i][j]
		}
	}

	b := make(Vector, lines)
	for i := range b {
		b[i] = custom[i][columns]
	}

	index := make([]int, len(a))
	for i := range index {
		index[i] = i
	}

	var errMult error
	var errNone error
	for i := 0; i < len(a); i++ {

		r := a[i][index[i]]
		if r == 0 {
			var kk int

			for k := i; k < len(a); k++ {
				if math.Abs(a[i][index[k]]) > r {
					kk = k
				}
			}

			if kk > 0 {
				index[i], index[kk] = index[kk], index[i]
			}
			r = a[i][index[i]]
		}

		if r == 0 {
			if b[i] == 0 {
				errMult = fmt.Errorf("система имеет множество решений")
				//fmt.Println("система имеет множество решений")
			} else {
				errNone = fmt.Errorf("система не имеет решений")
				//fmt.Println("система не имеет решений")
			}
		}

		for j := 0; j < len(a[i]); j++ {
			a[i][index[j]] /= r
		}
		b[i] /= r

		for k := i + 1; k < len(a); k++ {
			r = a[k][index[i]]
			for j := 0; j < len(a[i]); j++ {
				a[k][index[j]] = a[k][index[j]] - a[i][index[j]]*r
			}
			b[k] = b[k] - b[i]*r
		}
	}

	var x Vector = make(Vector, len(b))

	for i := len(a) - 1; i >= 0; i-- {
		x[i] = b[i]

		for j := i + 1; j < len(a); j++ {
			x[i] = x[i] - (x[j] * a[i][index[j]])
		}
	}

	return x, errMult, errNone
}
