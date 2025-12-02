package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount, yrs, returnRate := 1000.0, 10.0, 5.5

	investmentAmount = 2000.0

	futureValue := investmentAmount * math.Pow(1+returnRate/100, yrs)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, yrs)

	fmt.Println("Future Value =", futureValue)
	fmt.Println("Future Real Value =", futureRealValue)

	fmt.Println("\n--------------------- User Input -----------------------")

	const inflationRate2 = 2.5
	var investmentAmount2, yrs2, returnRate2 float64

	newFutureValue, newFutureRealValue := userInputCalc(inflationRate2, investmentAmount2, yrs2, returnRate2)

	fmt.Printf("\nFuture Value = %v\nFuture Real Value = %v", newFutureValue, newFutureRealValue)

	x := 7.0897079
	fmt.Printf("\nHi there %.2f", x)
	fmt.Printf("\nHi there %v\n", x)

	v := fmt.Sprintf("\nMy name is %v", "Alex")
	w := fmt.Sprintf("\nI am %v yrs old\n", 19)

	fmt.Print(v, w)

	outputText("My name is Gon")
}

func outputText(x string) {
	fmt.Printf("%v\n", x)
}

func userInputCalc(inflationRate2, investmentAmount2, yrs2, returnRate2 float64) (float64, float64) {
	fmt.Print("Investment Amount = ")
	fmt.Scan(&investmentAmount2)
	fmt.Print("Investment Period = ")
	fmt.Scan(&yrs2)
	fmt.Print("Return Rate = ")
	fmt.Scan(&returnRate2)

	futureValue2 := investmentAmount2 * math.Pow(1+returnRate2/100, yrs2)
	futureRealValue2 := futureValue2 / math.Pow(1+inflationRate2/100, yrs2)

	return futureValue2, futureRealValue2
}
