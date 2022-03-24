package lab_3

import (
	"fmt"
	"math"
)

var sumOdd float64
var sumEven float64
var counter = 0
var yI0 float64

func startSimpson() {
	var res float64 = 0
	var prevRes float64 = 0
	n = 4
	for true {
		entryPointSimpson(n)
		res = result
		if prevRes != 0 && math.Abs(res-prevRes) <= precision {
			break
		}
		prevRes = res
		fmt.Printf("res = %f\n", res)
		n *= 2
	}
}

func entryPointSimpson(innerN int) {
	sumEven = 0
	sumOdd = 0
	result = 0
	counter = 0
	xI1 = a
	h = (b - a) / float64(innerN)
	xI = xI1 + h
	yI0 = f(a)
	yI = f(a)
	counter++

	//fmt.Printf("y%d = %f \n", 0, yI0)
	//fmt.Printf("x%d = %f \n", 0, xI1)
	//fmt.Printf("h = %f \n\n", h)

	for i := 0; i < innerN; i++ {
		//fmt.Printf("y%d = %f \n", i, yI)
		//fmt.Printf("x%d = %f \n", i+1, xI)
		yI = f(xI)
		xI1 = xI
		xI = xI1 + h
		if counter != innerN {
			if counter%2 == 0 {
				sumEven += yI
			} else {
				sumOdd += yI
			}
		}
		counter++
	}
	yI = f(b)
	//fmt.Printf("y%d = %f \n", 6, yI)
	//fmt.Printf("sumEven = %f \n", sumEven)
	//fmt.Printf("sumOdd = %f \n", sumOdd)
	result = h / 3 * (yI0 + 4*sumOdd + 2*sumEven + yI)
}
