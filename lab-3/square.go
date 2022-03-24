package lab_3

import (
	"fmt"
	"math"
)

var a float64
var b float64
var n int
var precision float64
var yI1 float64
var yI float64
var h float64
var xI float64
var xI1 float64
var result float64

func startSquare() {
	var res float64 = 0
	var prevRes float64 = 0
	xI1 = a
	yI1 = f(a)
	//fmt.Printf("y%d = %f \n", 0, yI1)
	//fmt.Printf("x%d = %f \n", 0, xI1)
	n = 4
	//counter := 6
	for true {
		entryPointSquare(calcType, n)
		res = result
		if prevRes != 0 && math.Abs(res-prevRes) <= precision {
			//if prevRes != 0 && counter == 0 {
			break
		}
		prevRes = res
		fmt.Printf("n = %d; res = %f\n", n, res)
		n *= 2
	}
}

func entryPointSquare(calcType string, innerN int) {
	result = 0
	if calcType == "left" {
		xI1 = a
		h = (b - a) / float64(innerN)
		xI = xI1 + h
		yI1 = f(a)
		result += h * yI1

		for i := 0; i < innerN; i++ {
			xI1 = xI
			yI1 = f(xI)
			xI = xI1 + h
			result += h * yI1
		}
	} else if calcType == "right" {
		xI1 = a
		h = (b - a) / float64(innerN)
		xI = xI1 + h
		yI = f(xI)
		result += h * yI

		for i := 0; i < innerN; i++ {
			xI1 = xI
			xI = xI1 + h
			yI = f(xI)
			result += h * yI
		}
	} else if calcType == "center" {
		h = (b - a) / float64(innerN)
		xI1 = a
		xI = xI1 + h
		yI = f((xI1 + xI) / 2)
		result += h * yI
		//fmt.Printf("y%d = %f \n", 0, yI)
		//fmt.Printf("x%d = %f \n", 0, xI1)
		//fmt.Printf("h = %f \n\n", h)

		for i := 0; i < innerN; i++ {
			result += h * yI
			//fmt.Printf("y%d = %f \n", i+1, yI)
			//fmt.Printf("x%d = %f \n", i+1, xI)

			xI1 = xI
			xI = xI1 + h
			yI = f((xI1 + xI) / 2)
		}
	}
}
