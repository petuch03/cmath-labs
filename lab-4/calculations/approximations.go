package calculations

import (
	"fmt"
	m "github.com/erkkah/margaid"
	"math"
)

func LinearApproximation(inputSeries [][]float64, size int) *m.Series {
	a := matrix{
		{GlobalSums.x2, GlobalSums.x1, GlobalSums.xy},
		{GlobalSums.x1, float64(size), GlobalSums.y},
	}

	linearAnswers := CalculateMatrix(a, 2)

	linearSeries := m.NewSeries()

	sumx := 0.0
	sumOldY := 0.0
	for i := 0; i < size; i++ {
		sumx += inputSeries[i][0]
		sumOldY += inputSeries[i][1]
	}
	avgX := sumx / float64(size)
	avgY := sumOldY / float64(size)
	sumMult := 0.0
	sumXForSquare := 0.0
	sumYForSquare := 0.0

	eps := 0.0
	y2 := 0.0
	sumy := 0.0
	for i := 0; i < size; i++ {
		y := linearAnswers[0]*inputSeries[i][0] + linearAnswers[1]
		linearSeries.Add(m.MakeValue(inputSeries[i][0], y))

		sumMult += (inputSeries[i][0] - avgX) * (inputSeries[i][1] - avgY)
		sumXForSquare += (inputSeries[i][0] - avgX) * (inputSeries[i][0] - avgX)
		sumYForSquare += (inputSeries[i][1] - avgY) * (inputSeries[i][1] - avgY)

		y2 += y * y
		sumy += y
		eps += math.Pow(math.Abs(y-inputSeries[i][1]), 2)
	}

	fmt.Printf("---linear---\neps = %f\n", eps)
	fmt.Printf("R^2 = %f\n", 1-eps/(y2-(sumy/float64(size))))
	fmt.Printf("r = %f\n", sumMult/math.Sqrt(sumXForSquare*sumYForSquare))
	fmt.Printf("P(x) = %f * x + %f\n\n", linearAnswers[0], linearAnswers[1])

	return linearSeries
}

func ExponentApproximation(inputSeries [][]float64, size int) *m.Series {

	exponentAnswers := make([]float64, 2)

	exponentAnswers[1] =
		(float64(len(inputSeries))*GlobalSums.xLnY - GlobalSums.x1*GlobalSums.lnY) / (float64(len(inputSeries))*GlobalSums.x2 - math.Pow(GlobalSums.x1, 2))

	exponentAnswers[0] =
		GlobalSums.lnY/float64(len(inputSeries)) - exponentAnswers[1]*GlobalSums.x1/float64(len(inputSeries))
	expSeries := m.NewSeries()
	eps := 0.0
	y2 := 0.0
	sumy := 0.0
	for i := 0; i < size; i++ {
		if inputSeries[i][1] >= 0.0000000000001 {
			y := math.Exp(exponentAnswers[1]*inputSeries[i][0] + exponentAnswers[0])
			expSeries.Add(m.MakeValue(inputSeries[i][0], y))
			y2 += y * y
			sumy += y
			eps += math.Pow(math.Abs(y-inputSeries[i][1]), 2)
		}
	}
	//fmt.Printf("%f %f", exponentAnswers[0], exponentAnswers[1])
	fmt.Printf("---exponent---\neps = %f\n", eps)
	fmt.Printf("R^2 = %f\n", 1-eps/(y2-(sumy/float64(size))))
	fmt.Printf("P(x) = e ^ (%f * x + %f)\n\n", exponentAnswers[1], exponentAnswers[0])

	return expSeries
}

func LogApproximation(inputSeries [][]float64, size int) *m.Series {
	logAnswers := make([]float64, 2)

	logAnswers[1] =
		(float64(len(inputSeries))*GlobalSums.yLnX - GlobalSums.lnX*GlobalSums.y) /
			(float64(len(inputSeries))*GlobalSums.ln2X - math.Pow(GlobalSums.lnX, 2))

	logAnswers[0] =
		(GlobalSums.y / float64(len(inputSeries))) -
			(logAnswers[1] * GlobalSums.lnX / float64(len(inputSeries)))

	logSeries := m.NewSeries()
	eps := 0.0
	y2 := 0.0
	sumy := 0.0
	for i := 0; i < size; i++ {
		if inputSeries[i][0] != math.NaN() && inputSeries[i][0] >= 0.0000000000001 {
			y := logAnswers[0] + logAnswers[1]*math.Log(inputSeries[i][0])
			logSeries.Add(m.MakeValue(inputSeries[i][0], y))
			y2 += y * y
			sumy += y
			eps += math.Pow(math.Abs(y-inputSeries[i][1]), 2)
		}
	}
	fmt.Printf("---log---\neps = %f\n", eps)
	fmt.Printf("R^2 = %f\n", 1-eps/(y2-(sumy/float64(size))))
	fmt.Printf("P(x) = %f * ln(x) + %f\n\n", logAnswers[1], logAnswers[0])
	return logSeries
}

func PowApproximation(inputSeries [][]float64, size int) *m.Series {
	powAnswers := make([]float64, 2)

	powAnswers[1] =
		(float64(len(inputSeries))*GlobalSums.lnXLnY - GlobalSums.lnX*GlobalSums.lnY) /
			(float64(len(inputSeries))*GlobalSums.ln2X - math.Pow(GlobalSums.lnX, 2))

	powAnswers[0] = math.Exp(
		GlobalSums.lnY/float64(len(inputSeries)) -
			powAnswers[1]*GlobalSums.lnX/float64(len(inputSeries)))

	powSeries := m.NewSeries()
	eps := 0.0
	y2 := 0.0
	sumy := 0.0
	for i := 0; i < size; i++ {
		if inputSeries[i][0] >= 0.0000000000001 {
			y := math.Pow(inputSeries[i][0], powAnswers[1]) * powAnswers[0]
			powSeries.Add(m.MakeValue(inputSeries[i][0], y))
			y2 += y * y
			sumy += y
			eps += math.Pow(math.Abs(y-inputSeries[i][1]), 2)
		}
	}
	fmt.Printf("---power---\neps = %f\n", eps)
	fmt.Printf("R^2 = %f\n", 1-eps/(y2-(sumy/float64(size))))
	fmt.Printf("P(x) = x ^ (%f) + %f\n\n", powAnswers[1], powAnswers[0])

	return powSeries
}

func QuadraticApproximation(inputSeries [][]float64, size int) *m.Series {
	a := matrix{
		{float64(len(inputSeries)), GlobalSums.x1, GlobalSums.x2, GlobalSums.y},
		{GlobalSums.x1, GlobalSums.x2, GlobalSums.x3, GlobalSums.xy},
		{GlobalSums.x2, GlobalSums.x3, GlobalSums.x4, GlobalSums.x2y},
	}

	quadraticAnswers := CalculateMatrix(a, 3)

	quadraticSeries := m.NewSeries()
	eps := 0.0
	y2 := 0.0
	sumy := 0.0
	for i := 0; i < size; i++ {
		y := quadraticAnswers[2]*inputSeries[i][0]*inputSeries[i][0] + quadraticAnswers[1]*inputSeries[i][0] + quadraticAnswers[0]
		quadraticSeries.Add(m.MakeValue(inputSeries[i][0], y))
		y2 += y * y
		sumy += y
		eps += math.Pow(math.Abs(y-inputSeries[i][1]), 2)
	}

	fmt.Printf("---quadratic---\neps = %f\n", eps)
	fmt.Printf("R^2 = %f\n", 1-eps/(y2-(sumy/float64(size))))
	fmt.Printf("P(x) = %f * x^2 + %f * x + %f\n\n", quadraticAnswers[2], quadraticAnswers[1], quadraticAnswers[0])

	return quadraticSeries
}

func CubicApproximation(inputSeries [][]float64, size int) *m.Series {
	a := matrix{
		{float64(len(inputSeries)), GlobalSums.x1, GlobalSums.x2, GlobalSums.x3, GlobalSums.y},
		{GlobalSums.x1, GlobalSums.x2, GlobalSums.x3, GlobalSums.x4, GlobalSums.xy},
		{GlobalSums.x2, GlobalSums.x3, GlobalSums.x4, GlobalSums.x5, GlobalSums.x2y},
		{GlobalSums.x3, GlobalSums.x4, GlobalSums.x5, GlobalSums.x6, GlobalSums.x3y},
	}

	cubicAnswers := CalculateMatrix(a, 4)

	cubicSeries := m.NewSeries()
	eps := 0.0
	y2 := 0.0
	sumy := 0.0
	for i := 0; i < size; i++ {
		y := cubicAnswers[3]*math.Pow(inputSeries[i][0], 3) + cubicAnswers[2]*math.Pow(inputSeries[i][0], 2) + cubicAnswers[1]*math.Pow(inputSeries[i][0], 1) + cubicAnswers[0]
		cubicSeries.Add(m.MakeValue(inputSeries[i][0], y))
		y2 += y * y
		sumy += y
		eps += math.Pow(math.Abs(y-inputSeries[i][1]), 2)
	}

	fmt.Printf("---cubic---\neps = %f\n", eps)
	fmt.Printf("R^2 = %f\n", 1-eps/(y2-(sumy/float64(size))))
	fmt.Printf("P(x) = %f * x^3 + %f * x^2 + %f * x + %f\n\n", cubicAnswers[3], cubicAnswers[2], cubicAnswers[1], cubicAnswers[0])

	return cubicSeries
}
