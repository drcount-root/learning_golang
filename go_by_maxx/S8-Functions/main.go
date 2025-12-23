package main

import "fmt"

func main() {
	sliceOfInts := make([]int, 6)

	for i := 0; i < len(sliceOfInts); i++ {
		sliceOfInts[i] = i
	}

	fmt.Println(sliceOfInts) // [0 1 2 3 4 5]

	doubledSlice := transformItFunc(&sliceOfInts, double)
	fmt.Println(doubledSlice) // [0 2 4 6 8 10]

	tripleSlice := transformItFunc(&sliceOfInts, triple)
	fmt.Println(tripleSlice) // [0 3 6 9 12 15]

	check1Slice := []int{1, 2, 3}
	check1 := transformItFunc(&check1Slice, getTransformarFunc(&check1Slice))
	fmt.Println(check1) // [2 4 6]

	check2Slice := []int{0, 9, 5}
	check2 := transformItFunc(&check2Slice, getTransformarFunc(&check2Slice))
	fmt.Println(check2) // [0 27 15]
}

type transformType func(*int) int

// transformType -> this denotes a function that takes a pointer to an int and returns an int

func transformItFunc(slice *[]int, transform transformType) []int {
	newSlice := []int{}

	// here we don't care about the index, thats why we use _
	for _, value := range *slice {
		newSlice = append(newSlice, transform(&value))
	}

	return newSlice
}

func getTransformarFunc(numbers *[]int) transformType {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(x *int) int {
	return *x * 2
}

func triple(x *int) int {
	return *x * 3
}
