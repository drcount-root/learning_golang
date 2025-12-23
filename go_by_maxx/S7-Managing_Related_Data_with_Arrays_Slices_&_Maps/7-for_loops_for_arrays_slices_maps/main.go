package main

import (
	"fmt"
)

func main() {
	// Array
	arrX := [6]int{9, 4, 1, 7, 3, 2}

	for i, v := range arrX {
		fmt.Printf("Index %v : Value %v\n", i, v)
	}

	userNames2 := make([]string, 3)
	// userNames2[0] = "Chris"
	userNames2[1] = "David"
	userNames2[2] = "Kate"

	for i, v := range userNames2 {
		fmt.Printf("Index %v : Value %v\n", i, v)
	}

	/*
		Index 0 : Value
		Index 1 : Value David
		Index 2 : Value Kate
	*/

	// Slice
	sliceX := []int{}
	sliceX = append(sliceX, arrX[:3]...)

	for k, v := range sliceX {
		fmt.Printf("Index %v : Value %v\n", k, v)
	}

	/*
		Index 0 : Value 9
		Index 1 : Value 4
		Index 2 : Value 1
	*/

	// Map

	mapX := map [string]float64 {
		"USD": 1.0,
		"EUR": 0.9,
		"GBP": 0.8,
	}

	for i, v := range mapX {
		mapX[i] = v + 0.1
	}

	fmt.Printf("mapX => %v\n", mapX) // map[EUR:1 GBP:0.9 USD:1.1]
}
