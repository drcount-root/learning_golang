package main

import "fmt"

func main() {
	dataSetArr := [5]string{}

	fmt.Println(dataSetArr)

	dataSetIntArr := [5]int{}

	fmt.Println(dataSetIntArr) // [0 0 0 0 0] // for floats as well

	dataSetArr[0] = "a"
	dataSetArr[1] = "b"
	dataSetArr[2] = "c"
	dataSetArr[3] = "d"
	dataSetArr[4] = "e"

	fmt.Println(dataSetArr)

	dataSetSlice := []string{"s", "d", "e", "f", "s"}

	// dataSet2Slice := make([]int, 3, 5) // this slice has 3 elements
	// allocates an underlying array of size 5 and returns a slice of length 3 and capacity 5 that is backed by this underlying array

	fmt.Println(dataSetSlice)

	x := append(dataSetArr[:], "k") // arrays cant be appended, only possible after transforming them into slices
	fmt.Println(x)

	m := make(map[string]int, 6)

	for i := 0; i <= 5; i++ {
		m[fmt.Sprintf("%v", i)] = i
	}

	fmt.Println(m)

	delete(m, "4") // only maps are having delete mothod

	fmt.Println(m)

	for i := 0; i <= 10; i++ {
		switch {
		case i > 3:
			fmt.Println("Greater than 3")
		case i == 3:
			continue // skip when i == 3
		case i < 3:
			fmt.Println("Less than 3")
		default:
			fmt.Println("Default")
		}
	}

	arrX := []int{1, 4, 5, 6, 8}

	for i, v := range arrX {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	strX := "Hey there!"

	for i, v := range strX {
		fmt.Printf("Index: %v, Value: %c\n", i, v) // used %c to format the character
	}

}
