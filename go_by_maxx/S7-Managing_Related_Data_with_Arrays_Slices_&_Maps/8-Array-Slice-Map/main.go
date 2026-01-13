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
	fmt.Println(dataSetSlice)

	dataSet2Slice := make([]int, 3, 5) // this slice has 3 elements
	// allocates an underlying array of size 5 and returns a slice of length 3 and capacity 5 that is backed by this underlying array

	for i := 0; i <= 2; i++ {
		dataSet2Slice[i] = i
	}

	fmt.Printf("dataSet2Slice => %v with length %v and capacity %v\n", dataSet2Slice, len(dataSet2Slice), cap(dataSet2Slice))

	for i := 0; i <= 6; i++ {
		dataSet2Slice = append(dataSet2Slice, i) // now as we append more elements, the capacity is increased to 10
	}

	fmt.Printf("dataSet2Slice => %v with length %v and capacity %v\n", dataSet2Slice, len(dataSet2Slice), cap(dataSet2Slice))

	for i := 7; i <= 12; i++ {
		dataSet2Slice = append(dataSet2Slice, i) // now as we append more elements, the capacity is increased to 20 so doubled
	}

	fmt.Printf("dataSet2Slice => %v with length %v and capacity %v\n", dataSet2Slice, len(dataSet2Slice), cap(dataSet2Slice))

	x := append(dataSetArr[:], "k") // arrays cant be appended, only possible after transforming them into slices
	fmt.Println(x)

	m := make(map[string]int, 6)

	for i := 0; i <= 5; i++ {
		m[fmt.Sprintf("%v", i)] = i
	}

	fmt.Println(m)

	delete(m, "4") // only maps are having delete mothod

	fmt.Println(m)
}
