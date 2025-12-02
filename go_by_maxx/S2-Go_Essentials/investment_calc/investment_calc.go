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

	fmt.Print("Investment Amount = ")
	fmt.Scan(&investmentAmount2)
	fmt.Print("Investment Period = ")
	fmt.Scan(&yrs2)
	fmt.Print("Return Rate = ")
	fmt.Scan(&returnRate2)

	newFutureValue, newFutureRealValue := userInputCalc(inflationRate2, investmentAmount2, yrs2, returnRate2)

	fmt.Printf("\nFuture Value = %v\nFuture Real Value = %v", newFutureValue, newFutureRealValue)

	newFutureValue1, newFutureRealValue1 := userInputCalcReturnStyle(inflationRate2, investmentAmount2, yrs2, returnRate2)

	fmt.Printf("\nNew Return Style Future Value = %v\nFuture Real Value = %v", newFutureValue1, newFutureRealValue1)

	x := 7.0897079
	fmt.Printf("\nHi there %.2f", x)
	fmt.Printf("\nHi there %v\n", x)

	v := fmt.Sprintf("\nMy name is %v", "Alex")
	w := fmt.Sprintf("\nI am %v yrs old\n", 19)

	fmt.Print(v, w)

	outputText("My name is Gon")

	// Revenue calc
	revenue := getUserInput("Revenue:")
	expense := getUserInput("Expense:")
	taxRate := getUserInput("Tax Rate:")

	ebt, profit, ratio := revCalc(revenue, expense, taxRate)

	fmt.Printf("\nEBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)
}

func outputText(x string) {
	fmt.Printf("%v\n", x)
}

func userInputCalc(inflationRate2, investmentAmount2, yrs2, returnRate2 float64) (float64, float64) {
	futureValue2 := investmentAmount2 * math.Pow(1+returnRate2/100, yrs2)
	futureRealValue2 := futureValue2 / math.Pow(1+inflationRate2/100, yrs2)
	return futureValue2, futureRealValue2
}

func userInputCalcReturnStyle(inflationRate2, investmentAmount2, yrs2, returnRate2 float64) (futureValue2 float64, futureRealValue2 float64) {
	futureValue2 = investmentAmount2 * math.Pow(1+returnRate2/100, yrs2)
	futureRealValue2 = futureValue2 / math.Pow(1+inflationRate2/100, yrs2)
	return
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func revCalc(revenue, expense, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
