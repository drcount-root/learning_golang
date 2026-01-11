package main

import "fmt"

func main() {
	/*
		In Go, arrays and slices are different types:

		names := []string{}   // names is a slice as it is dynamic
		ages  := [3]int{}     // ages is an array of predefined length

		names is a slice. A slice is NOT a dynamically growing array by itself.
		It is a lightweight descriptor over an underlying array, consisting of:
			- a pointer to the underlying array
			- length
			- capacity

		When we append to a slice, Go may allocate a new underlying array
		and update the slice header accordingly.
		names := []string{}

		ages is an array. Arrays have a fixed length that is part of their type.
		The length (3) is known at compile time and cannot change.
		Arrays are value types and are copied when passed around.
		ages := [3]int{}

		valll := [3]string {}
		valll.append("abc")
		valll.append undefined (type [3]string has no field or method append

		Why is that? because append is a method of slice, and as while declaring valll we have declared it as an array with fixed size
	*/

	// Declare a slice of int values with dynamic size
	prices := []int{}

	// initialize the slice with some values
	prices = []int{1, 3, 4, 7, 8}

	// fmt.Printf("\nprices[6] => %v\n", prices[6])
	// panic: runtime error: index out of range [6] with length 5

	updatedPrices := append(prices, 10, 12, 15)
	// append doesn't change the original slice instead it returns a new slice
	fmt.Printf("\nupdatedPrices => %v\n", updatedPrices) // [1 3 4 7 8 10 12 15]
	fmt.Printf("\nprices => %v\n", prices)               // [1 3 4 7 8]

	// To make it append the original arr we can reassign like this
	// prices = append(prices, 10, 12, 15) // now prices = [1 3 4 7 8 10 12 15]

	// to remove an element from the slice go don't have any built-in method
	afterRemovalUpdatedPrices := append(updatedPrices[:1], updatedPrices[3:]...)
	// by doing updatedPrices[3:]... we are unpacking all the values from updatedPrices[3:]
	fmt.Printf("\nafterRemovalUpdatedPrices => %v\n", afterRemovalUpdatedPrices) // [1 7 8 10 12 15]

	// array
	// v := [2]float64{0, 3.2}
	// v[3] = 4.5 // invalid argument: index 3 out of bounds [0:2]

	// slice
	// w := []float64{}
	// println(len(w)) // 0
	// w[1] = 3.14 // panic: runtime error: index out of range [1] with length 0
}

/*
	Slice Summary (Mental Model)
	Operation	Supported
	Append	✅
	Delete	✅ (manual)
	Resize	✅
	Copy	✅
	Sort	✅
	Compare	✅
	Pass to function	Reference-like
*/
