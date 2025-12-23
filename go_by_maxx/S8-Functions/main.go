package main

import "fmt"

func main() {
	arrOfInts := make([]int, 6)

	for i := 0; i < len(arrOfInts); i++ {
		arrOfInts[i] = i
	}

	fmt.Println(arrOfInts) // [0 1 2 3 4 5]

	doubledArr := transformItFunc(&arrOfInts, double)
	fmt.Println(doubledArr) // [0 2 4 6 8 10]

	tripleArr := transformItFunc(&arrOfInts, triple)
	fmt.Println(tripleArr) // [0 3 6 9 12 15]
}

type transformType func(*int) int

// transformType -> this denotes a function that takes a pointer to an int and returns an int

func transformItFunc(arr *[]int, transform transformType) []int {
	newArr := []int{}

	// here we don't care about the index, thats why we use _
	for _, value := range *arr {
		newArr = append(newArr, transform(&value))
	}

	return newArr
}

func double(x *int) int {
	return *x * 2
}

func triple(x *int) int {
	return *x * 3
}
