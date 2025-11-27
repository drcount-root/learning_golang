package main

import (
	"fmt" 
	"math"
)

const str string = "Hi there"

func main() {
	fmt.Println(str)

	const pi float64 = 3.14
	fmt.Println("Value of pi =>", pi)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
