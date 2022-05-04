package calculations_lab_4

import (
	"math"
)

type Sums struct {
	x1 float64
	x2 float64
	x3 float64
	x4 float64
	x5 float64
	x6 float64

	y   float64
	xy  float64
	x2y float64
	x3y float64
	x4y float64

	lnX    float64
	lnY    float64
	ln2X   float64
	lnXLnY float64
	xLnY   float64
	yLnX   float64
}

var GlobalSums Sums

func FillConstants(inputSeries [][]float64) {
	//for i := 0; i < 11; i++ {
	//	_, _ = fmt.Printf("%f %f\n", inputSeries[i][0], inputSeries[i][1])
	//}
	GlobalSums.x1 = 0
	GlobalSums.x2 = 0
	GlobalSums.x3 = 0
	GlobalSums.x4 = 0
	GlobalSums.x5 = 0
	GlobalSums.x6 = 0
	GlobalSums.y = 0
	GlobalSums.xy = 0
	GlobalSums.x2y = 0
	GlobalSums.x3y = 0
	GlobalSums.x4y = 0
	GlobalSums.lnX = 0
	GlobalSums.lnY = 0
	GlobalSums.ln2X = 0
	GlobalSums.lnXLnY = 0
	GlobalSums.xLnY = 0
	GlobalSums.yLnX = 0

	for i := 0; i < len(inputSeries); i++ {
		GlobalSums.x1 += inputSeries[i][0]
		GlobalSums.x2 += math.Pow(inputSeries[i][0], 2)
		GlobalSums.x3 += math.Pow(inputSeries[i][0], 3)
		GlobalSums.x4 += math.Pow(inputSeries[i][0], 4)
		GlobalSums.x5 += math.Pow(inputSeries[i][0], 5)
		GlobalSums.x6 += math.Pow(inputSeries[i][0], 6)

		GlobalSums.y += inputSeries[i][1]
		GlobalSums.xy += inputSeries[i][0] * inputSeries[i][1]
		GlobalSums.x2y += math.Pow(inputSeries[i][0], 2) * inputSeries[i][1]
		GlobalSums.x3y += math.Pow(inputSeries[i][0], 3) * inputSeries[i][1]
		GlobalSums.x4y += math.Pow(inputSeries[i][0], 4) * inputSeries[i][1]

		if inputSeries[i][0] >= 0.0000000000001 && inputSeries[i][1] >= 0.0000000000001 {
			//_, _ = fmt.Printf("----%f ---- %f----\n", inputSeries[i][0], inputSeries[i][1])
			GlobalSums.lnX += math.Log(inputSeries[i][0])
			GlobalSums.lnY += math.Log(inputSeries[i][1])
			GlobalSums.ln2X += math.Pow(math.Log(inputSeries[i][0]), 2)
			GlobalSums.lnXLnY += math.Log(inputSeries[i][0]) * math.Log(inputSeries[i][1])
			GlobalSums.xLnY += inputSeries[i][0] * math.Log(inputSeries[i][1])
			GlobalSums.yLnX += inputSeries[i][1] * math.Log(inputSeries[i][0])
		}
	}
}
