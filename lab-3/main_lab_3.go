package lab_3

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"math"
	"os"
)

var index = 1
var calcType string
var listOfExpressions = []string{
	"3*x*x*x - 4*x*x + 5*x - 16",
	"x",
}
var currentExpression = listOfExpressions[index-1]

func MainLab3() {

	println("\nlist of expressions:")
	for i := range listOfExpressions {
		println(i+1, ") ", listOfExpressions[i])
	}
	println()
	print("index of expression: ")
	_, _ = fmt.Scanf("%d", &index)
	print("        lower bound: ")
	_, _ = fmt.Scanf("%g", &a)
	print("        upper bound: ")
	_, _ = fmt.Scanf("%g", &b)
	print("          precision: ")
	_, _ = fmt.Scanf("%g", &precision)
	print("          calc type: ")
	_, _ = fmt.Scanf("%s", &calcType)
	currentExpression = listOfExpressions[index-1]
	println()

	if calcType == "left" || calcType == "right" || calcType == "center" {
		println("===calculation initiated===")
		//entryPointSquare(calcType, 6)
		startSquare()
		fmt.Printf("      type: %s\n", calcType)
		printf64("square res:", result)
		fmt.Printf("square   n: %d", n)
	} else if calcType == "simpson" {
		println("===calculation initiated===")
		entryPointSimpson(6)
		//startSimpson()
		printf64("simpson res:", result)
		fmt.Printf("simpson   n: %d", n)
	} else {
		println("INVALID")
	}

}

func f(x float64) float64 {
	expression, _ := govaluate.NewEvaluableExpression(currentExpression)

	parameters := make(map[string]interface{}, 8)
	parameters["x"] = x
	parameters["sin_x"] = math.Sin(x)
	parameters["cos_x"] = math.Cos(x)
	result, _ := expression.Evaluate(parameters)

	switch i := result.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	default:
		os.Exit(1)
		return -1
	}
}

func printf64(s string, i float64) {
	fmt.Printf("%s %f\n", s, i)
}
