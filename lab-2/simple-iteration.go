package lab_2

import (
	"math"
)

var xk []myType
var fxk []myType
var xk1 []myType
var phixk []myType
var modxk []myType

var x0 float64
var currentXk float64
var currentFxk float64
var currentXk1 float64
var currentPhixk float64
var currentModxk float64
var lambda float64

func entryPointIteration() {
	currentA = lowerBound
	currentB = upperBound
	lambda = -1 / math.Max(derivative(currentA), derivative(currentB))
	if derivative(currentA) > derivative(currentB) {
		x0 = currentA
	} else {
		x0 = currentB
	}

	currentXk = x0
	currentFxk = function(currentXk) // k = 0
	xk = append(xk, myType{currentXk})
	fxk = append(fxk, myType{currentFxk})

	currentPhixk = currentXk + lambda*currentFxk
	phixk = append(phixk, myType{currentPhixk})
	currentXk1 = currentPhixk
	xk1 = append(xk1, myType{currentXk1})

	currentModxk = math.Abs(currentXk - currentXk1)
	modxk = append(modxk, myType{currentXk})
	for precision < math.Abs(currentModxk) && counter <= 50 {
		currentXk = currentXk1
		currentFxk = function(currentXk)
		currentPhixk = currentXk + lambda*currentFxk
		currentXk1 = currentPhixk
		currentModxk = math.Abs(currentXk - currentXk1)

		xk = append(xk, myType{currentXk})
		fxk = append(fxk, myType{currentFxk})
		phixk = append(phixk, myType{currentPhixk})
		xk1 = append(xk1, myType{currentXk1})
		modxk = append(modxk, myType{currentModxk})
	}
}
