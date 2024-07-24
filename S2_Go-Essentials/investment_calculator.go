package main

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount float64 = 1000
	// var years float64 = 10

	// var investmentAmount, years float64 = 1000, 10
	// expectedReturnRate := 5.5

	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	const inflationRate = 2.5

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
