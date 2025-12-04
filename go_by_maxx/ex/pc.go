package main

import (
	"errors"
	"fmt"
	"os"
)

const fileToOperateFrom = "balance.txt"

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		panic("Can't continue sorry!")
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		panic("Can't continue sorry!")
	}
	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		panic("Can't continue sorry!")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeToFile(ebt, profit, ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput < 0 {
		return 0, errors.New("⚠️ no negative value is allowed ⚠️")
	}

	if userInput == 0 {
		return 0, errors.New("⚠️ no zero value is allowed ⚠️")
	}

	return userInput, nil
}

func writeToFile(ebt, profit, ratio float64) {
	resultsString := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(fileToOperateFrom, []byte(resultsString), 0644)
}
