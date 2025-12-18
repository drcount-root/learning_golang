package main

import "fmt"

func main() {
	result1 := add(10, 20)
	fmt.Printf("Addition of integers: %v\n", result1) // Addition of integers: 30

	result2 := add(5.5, 4.5)
	fmt.Printf("Addition of floats: %v\n", result2) // Addition of floats: 10

	result3 := add("Hello, ", "World!")
	fmt.Printf("Addition of strings: %v\n", result3) // Addition of strings: Hello, World!

	result4 := add(10, "20")
	fmt.Printf("Addition of mixed types: %v\n", result4) // Addition of mixed types: <nil>

	result5 := add(10.6, 30)
	fmt.Printf("Addition of mixed types: %v\n", result5) // Addition of mixed types: <nil>

	// result := result1 + 1 // go doesn't understand what type result is & we can't perform the operations on it like how we are trying to add 1 to it.
	// This problem can be solved using Generics in Go.
	resultA := addWithGenerics(10, 20) // now go will understand this is int type

	fmt.Printf("Addition with resultA: %v\n", resultA+1) // 31

	resultB := addWithGenerics(5.5, 4.5) // now go will understand this is float64 type
	fmt.Printf("Addition with resultB: %v\n", resultB+78) // 88

	resultC := addWithGenerics("Hello, ", "Generics!") // now go will understand this is string type
	fmt.Printf("Addition with resultC: %v\n", resultC) // Hello, Generics!

	// Even though we can't add mixed types using generics
	// resultD := addWithGenerics(10, "20") // This will throw a compile-time error
}

// Suppose we want to add some data but we are not sure about its type
func add(a, b interface{}) interface{} {
	aInt, aIntOk := a.(int)
	bInt, bIntOk := b.(int)

	if aIntOk && bIntOk {
		return aInt + bInt
	}

	aFloat, aFloatOk := a.(float64)
	bFloat, bFloatOk := b.(float64)

	if aFloatOk && bFloatOk {
		return aFloat + bFloat
	}

	aStr, aStrOk := a.(string)
	bStr, bStrOk := b.(string)

	if aStrOk && bStrOk {
		return aStr + bStr
	}

	return nil
}

func addWithGenerics[T int | float64 | string](a, b T) T {
	return a + b
}
