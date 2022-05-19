package lab_6

import (
	m "github.com/erkkah/margaid"
	"math"
)

var table [][]float64
var a float64
var b float64
var h float64
var y0 float64
var n int
var precision float64
var euler [][]float64
var euler_half [][]float64
var adams [][]float64
var adams_half [][]float64
var functionResults []float64

func EulerMethod() *m.Series {
	n = int((b-a)/h) + 1
	euler = make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		euler[i] = make([]float64, 2)
	}
	euler[0][0] = a
	euler[0][1] = y0
	for i := 1; i < n; i++ {
		previousX := euler[i-1][0]
		previousY := euler[i-1][1]
		currentX := previousX + h
		currentY := previousY
		if index == 1 {
			currentY += h * function_1(currentX, previousY)
		} else if index == 2 {
			currentY += h * function_2(currentX, previousY)
		} else if index == 3 {
			currentY += h * function_3(currentX, previousY)
		}
		euler[i][0] = currentX
		euler[i][1] = currentY
	}
	EulerMethodHalf()
	n = int((b-a)/h) + 1
	finalSeries := m.NewSeries()

	for i := 0; i < n; i++ {
		finalSeries.Add(m.MakeValue(euler[i][0], euler[i][1]))
	}

	return finalSeries
}

func EulerMethodHalf() {
	local_h := h / 2
	n = int((b-a)/local_h) + 1
	euler_half = make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		euler_half[i] = make([]float64, 2)
	}
	euler_half[0][0] = a
	euler_half[0][1] = y0
	for i := 1; i < n; i++ {
		previousX := euler_half[i-1][0]
		previousY := euler_half[i-1][1]
		currentX := previousX + local_h
		currentY := previousY
		if index == 1 {
			currentY += local_h * function_1(currentX, previousY)
		} else if index == 2 {
			currentY += local_h * function_2(currentX, previousY)
		} else if index == 3 {
			currentY += local_h * function_3(currentX, previousY)
		}
		euler_half[i][0] = currentX
		euler_half[i][1] = currentY
	}
}

func AdamsMethod() *m.Series {
	n = int((b-a)/h) + 1
	adams = make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		adams[i] = make([]float64, 2)
	}
	functionResults = make([]float64, n+1)

	adams[0][0] = a
	adams[0][1] = y0

	var k1, k2, k3, k4 float64
	for i := 1; i < 4; i++ {
		previousX := adams[i-1][0]
		previousY := adams[i-1][1]
		if index == 1 {
			k1 = h * function_1(previousX, previousY)
			k2 = h * function_1(previousX+h/2, previousY+k1/2)
			k3 = h * function_1(previousX+h/2, previousY+k2/2)
			k4 = h * function_1(previousX+h, previousY+k3)
		} else if index == 2 {
			k1 = h * function_2(previousX, previousY)
			k2 = h * function_2(previousX+h/2, previousY+k1/2)
			k3 = h * function_2(previousX+h/2, previousY+k2/2)
			k4 = h * function_2(previousX+h, previousY+k3)
		} else if index == 3 {
			k1 = h * function_3(previousX, previousY)
			k2 = h * function_3(previousX+h/2, previousY+k1/2)
			k3 = h * function_3(previousX+h/2, previousY+k2/2)
			k4 = h * function_3(previousX+h, previousY+k3)
		}
		currentX := previousX + h
		currentY := previousY + (k1+2*k2+2*k3+k4)/6

		adams[i][0] = currentX
		adams[i][1] = currentY
		if index == 1 {
			functionResults[i] = function_1(currentX, currentY)
		} else if index == 2 {
			functionResults[i] = function_2(currentX, currentY)
		} else if index == 3 {
			functionResults[i] = function_3(currentX, currentY)
		}
	}
	for i := 4; i < n; i++ {
		currentX := adams[i-1][0] + h
		y_predictor := adams[i-1][1] + h*(55*functionResults[i-1]-59*functionResults[i-2]+37*functionResults[i-3]-9*functionResults[i-4])/24

		if index == 1 {
			functionResults[i] = function_1(currentX, y_predictor)
		} else if index == 2 {
			functionResults[i] = function_2(currentX, y_predictor)
		} else if index == 3 {
			functionResults[i] = function_3(currentX, y_predictor)
		}

		y_corrector := adams[i-1][1] + h*(9*functionResults[i]+19*functionResults[i-1]-5*functionResults[i-2]+functionResults[i-3])/24
		for precision < math.Abs(y_corrector-y_predictor) {
			y_predictor = y_corrector
			if index == 1 {
				functionResults[i] = function_1(currentX, y_predictor)
			} else if index == 3 {
				functionResults[i] = function_3(currentX, y_predictor)
			}
			y_corrector = adams[i-1][1] + h*(9*functionResults[i]+19*functionResults[i-1]-5*functionResults[i-2]+functionResults[i-3])/24
		}

		adams[i][0] = currentX
		adams[i][1] = y_corrector
	}
	AdamsMethodHalf()
	n = int((b-a)/h) + 1
	finalSeries := m.NewSeries()

	for i := 0; i < n; i++ {
		finalSeries.Add(m.MakeValue(adams[i][0], adams[i][1]))
	}

	return finalSeries
}

func AdamsMethodHalf() {
	local_h := h / 2
	n = int((b-a)/local_h) + 1
	adams_half = make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		adams_half[i] = make([]float64, 2)
	}
	functionResults = make([]float64, n+1)

	adams_half[0][0] = a
	adams_half[0][1] = y0

	var k1, k2, k3, k4 float64
	for i := 1; i < 4; i++ {
		previousX := adams_half[i-1][0]
		previousY := adams_half[i-1][1]
		if index == 1 {
			k1 = local_h * function_1(previousX, previousY)
			k2 = local_h * function_1(previousX+local_h/2, previousY+k1/2)
			k3 = local_h * function_1(previousX+local_h/2, previousY+k2/2)
			k4 = local_h * function_1(previousX+local_h, previousY+k3)
		} else if index == 2 {
			k1 = local_h * function_2(previousX, previousY)
			k2 = local_h * function_2(previousX+local_h/2, previousY+k1/2)
			k3 = local_h * function_2(previousX+local_h/2, previousY+k2/2)
			k4 = local_h * function_2(previousX+local_h, previousY+k3)
		} else if index == 3 {
			k1 = h * function_3(previousX, previousY)
			k2 = h * function_3(previousX+h/2, previousY+k1/2)
			k3 = h * function_3(previousX+h/2, previousY+k2/2)
			k4 = h * function_3(previousX+h, previousY+k3)
		}
		currentX := previousX + local_h
		currentY := previousY + (k1+2*k2+2*k3+k4)/6

		adams_half[i][0] = currentX
		adams_half[i][1] = currentY
		if index == 1 {
			functionResults[i] = function_1(currentX, currentY)
		} else if index == 2 {
			functionResults[i] = function_2(currentX, currentY)
		} else if index == 3 {
			functionResults[i] = function_3(currentX, currentY)
		}
	}
	for i := 4; i < n; i++ {
		currentX := adams_half[i-1][0] + local_h

		y_pred := adams_half[i-1][1] + local_h*(55*functionResults[i-1]-59*functionResults[i-2]+37*functionResults[i-3]-9*functionResults[i-4])/24

		if index == 1 {
			functionResults[i] = function_1(currentX, y_pred)
		} else if index == 2 {
			functionResults[i] = function_2(currentX, y_pred)
		} else if index == 3 {
			functionResults[i] = function_3(currentX, y_pred)
		}

		y_cor := adams_half[i-1][1] + local_h*(9*functionResults[i]+19*functionResults[i-1]-5*functionResults[i-2]+functionResults[i-3])/24

		for precision < math.Abs(y_cor-y_pred) {
			y_pred = y_cor
			if index == 1 {
				functionResults[i] = function_1(currentX, y_pred)
			} else if index == 2 {
				functionResults[i] = function_2(currentX, y_pred)
			} else if index == 3 {
				functionResults[i] = function_3(currentX, y_pred)
			}
			y_cor = adams_half[i-1][1] + local_h*(9*functionResults[i]+19*functionResults[i-1]-5*functionResults[i-2]+functionResults[i-3])/24
		}

		adams_half[i][0] = currentX
		adams_half[i][1] = y_cor
	}
}

func function_1(x float64, y float64) float64 {
	return y + (1+x)*y*y
}

func function_2(x float64, y float64) float64 {
	return math.Pow(x+1, 3) - y
}

func function_3(x float64, y float64) float64 {
	return 6*x*x + 5*y
}

func precise_1(x float64) float64 {
	return -math.Pow(math.E, x) / (x*math.Pow(math.E, x) + (math.E - math.E))
}

func precise_2(x float64) float64 {
	return const_2()*math.Pow(math.E, -x) + x*x*x + 3*x - 2
}

func precise_3(x float64) float64 {
	return const_3()*math.Pow(math.E, 5*x) - (6*x*x)/5 - (12*x)/25 - 12/125
}

func const_1() float64 {
	return math.Pow(-math.E, a)/y0 - a*math.Pow(math.E, a)
}

func const_2() float64 {
	return (y0 - a*a*a - 3*a + 2) * math.Pow(math.E, a)
}

func const_3() float64 {
	return 1 / 125 * (347*math.Pow(math.E, (5*a-5)) - 6*(25*a*a+10*a+2))
}
