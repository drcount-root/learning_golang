package main

import "fmt"

func main() {
	dataSetArr := [5]string{}

	dataSetArr[0] = "a"
	dataSetArr[1] = "b"
	dataSetArr[2] = "c"
	dataSetArr[3] = "d"
	dataSetArr[4] = "e"

	fmt.Println(dataSetArr)

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
