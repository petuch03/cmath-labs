package lab_4

import (
	"fmt"
	m "github.com/erkkah/margaid"
	"math"
)

func getRawSeries() *m.Series {
	size := Size

	rawSeries := m.NewSeries()
	for i := 0; i < size; i++ {
		rawSeries.Add(m.MakeValue(InputSeries[i][0], InputSeries[i][1]))
	}

	return rawSeries
}

func f(x float64) float64 {
	return 4 * x / (math.Pow(x, 4) + 4)
}

func input() {

	InputSeries = make([][]float64, Size)

	for i := range InputSeries {
		InputSeries[i] = make([]float64, 2)
	}
	for i := 0; i < Size; i++ {
		InputSeries[i][0] = 0 + (float64(i) * 0.2)
		InputSeries[i][1] = f(InputSeries[i][0])
	}

	//for i := 0; i < Size; i++ {
	//	_, _ = fmt.Printf("%15.14f %15.14f\n", InputSeries[i][0], InputSeries[i][1])
	//}

}

//{x1:-11 x2:15.399999999999999 x3:-24.200000000000003 x4:40.53280000000001 x5:-70.664 x6:126.61792000000001 y:-5.711742999999999 xy:6.5406088         x2y:-8.8445568         x3y:13.229669536000001 x4y:-21.1530269568     lnX:0                 lnY:0                 ln2X:0                  lnXLnY:0                  xLnY:0                       yLnX:0}
//{x1:-11 x2:15.399999999999999 x3:-24.200000000000003 x4:40.5328           x5:-70.664 x6:126.61792000000001 y:-5.711743425985285 xy:6.540608844868527 x2y:-8.844556596402272 x3y:13.229669117458394 x4y:-21.15302629605886 lnX:-36.7368005696771 lnY:-36.7368005696771 ln2X:1349.5925160962277 lnXLnY:1349.5925160962277 xLnY:-4.0786041843518085e-15 yLnX:-4.0786041843518085e-15}

func customInput() {
	_, _ = fmt.Scanf("%d\n", &Size)

	InputSeries = make([][]float64, Size)

	for i := range InputSeries {
		InputSeries[i] = make([]float64, 2)
	}
	for i := 0; i < Size; i++ {
		_, _ = fmt.Scanf("%f %f\n", &InputSeries[i][0], &InputSeries[i][1])
	}
}
