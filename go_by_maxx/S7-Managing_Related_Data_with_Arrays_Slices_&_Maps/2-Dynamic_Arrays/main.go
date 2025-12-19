package main

import "fmt"

func main() {
	// Declare an array of int values with dynamic size
	prices := []int{}

	// initialize the array with some values
	prices = []int{1, 3, 4, 7, 8}

	// fmt.Printf("\nprices[6] => %v\n", prices[6])
	// panic: runtime error: index out of range [6] with length 5

	updatedPrices := append(prices, 10, 12, 15)
	// append doesn't change the original array instead it returns a new array
	fmt.Printf("\nupdatedPrices => %v\n", updatedPrices) // [1 3 4 7 8 10 12 15]
	fmt.Printf("\nprices => %v\n", prices) // [1 3 4 7 8]

	// to remove an element from the array go don't have any built-in method
	afterRemovalUpdatedPrices := append(updatedPrices[:1], updatedPrices[3:]...)
	fmt.Printf("\nafterRemovalUpdatedPrices => %v\n", afterRemovalUpdatedPrices) // [1 7 8 10 12 15]
}
