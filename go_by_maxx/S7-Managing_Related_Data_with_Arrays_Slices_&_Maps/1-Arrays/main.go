package main

import "fmt"

func main() {
	// Declare an array of 4 float64 values
	prices := [4]float64{12.3, 45.6, 78.9, 23.4}

	fmt.Printf("Price array: %v, which is of length %v\n", prices, len(prices))

	// Declare an array of 6 string without initial values
	var productNames [6]string
	fmt.Printf("1 Product names array: %v, which is of length %v\n", productNames, len(productNames))
	productNames = [6]string{"Laptop", "Smartphone", "Tablet", "Monitor", "Keyboard", "Mouse"}
	fmt.Printf("2 Product names array after adding values: %v, which is of length %v\n", productNames, len(productNames))

	productNames[2] = "Smartwatch"
	fmt.Printf("3 Product names array after updating value of index 2: %v, which is of length %v\n", productNames, len(productNames))

}
