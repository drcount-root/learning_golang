package main

import "fmt"

func main() {
	arrOfInts := make([]int, 6)

	for i := 0; i < len(arrOfInts); i++ {
		arrOfInts[i] = i
	}

	fmt.Println(arrOfInts) // [0 1 2 3 4 5]

	doubledArr := doubleItFunc(&arrOfInts)

	fmt.Println(doubledArr) // [0 2 4 6 8 10]
}

func doubleItFunc(arr *[]int) []int {
	newArr := []int{}

	// here we don't care about the index, thats why we use _
	for _, value := range *arr {
		newArr = append(newArr, double(&value))
	}

	return newArr
}

func double(x *int) int {
	return *x * 2
}
