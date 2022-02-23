package lab_2

import (
	"math"
	"os"
)

type myType struct {
	value float64
}

var a []myType
var b []myType
var x []myType
var fa []myType
var fb []myType
var fx []myType
var mod []myType

var currentX float64
var currentFa float64
var currentFb float64
var currentFx float64
var currentMod float64

func entryPointHalf() {

	currentA = lowerBound
	currentB = upperBound
	currentX = (currentA + currentB) / 2

	currentFa = function(currentA)
	currentFb = function(currentB)
	currentFx = function(currentX)
	currentMod = math.Abs(currentA - currentB)

	a = append(a, myType{currentA})
	b = append(b, myType{currentB})
	x = append(x, myType{currentX})

	fa = append(fa, myType{currentFa})
	fb = append(fb, myType{currentFb})
	fx = append(fx, myType{currentFx})
	mod = append(mod, myType{currentMod})

	for precision < math.Abs(currentFx) && precision < currentMod {
		if checkHalf() == 0 {
			currentB = currentX
			currentX = (currentA + currentB) / 2
			currentFa = function(currentA)
			currentFb = function(currentB)
			currentFx = function(currentX)
			currentMod = math.Abs(currentA - currentB)
		} else if checkHalf() == 1 {
			currentA = currentX
			currentX = (currentA + currentB) / 2
			currentFa = function(currentA)
			currentFb = function(currentB)
			currentFx = function(currentX)
			currentMod = math.Abs(currentA - currentB)
		} else {
			os.Exit(1)
		}

		a = append(a, myType{currentA})
		b = append(b, myType{currentB})
		x = append(x, myType{currentX})

		fa = append(fa, myType{currentFa})
		fb = append(fb, myType{currentFb})
		fx = append(fx, myType{currentFx})
		mod = append(mod, myType{currentMod})
	}
}

func checkHalf() int {
	if currentFa*currentFx < 0 {
		return 0
	} else {
		return 1
	}
}
