package main

import "fmt"

func main() {
	age := 3
	// the address of age from memory (RAM) is stored in the variable agePointer
	// agePointer will be of type *int
	agePointer := &age
	fmt.Printf("Address of age is %v\n", agePointer)
	fmt.Printf("Value of age from pointer is %v\n", *agePointer)

	amount := 7878.988
	// amountPointer will be of type *float64
	amountPointer := &amount
	fmt.Printf("Address of amount is %v\n", amountPointer)
	fmt.Printf("Value of amount from pointer is %v\n", *amountPointer)

	someString := "hello"
	// someStringPointer will be of type *string
	someStringPointer := &someString
	fmt.Printf("Address of someString is %v\n", someStringPointer)
	fmt.Printf("Value of someString from pointer is %v\n", *someStringPointer)

	callX(agePointer)
	fmt.Printf("2nd Value of age from pointer is %v\n", *agePointer)
	fmt.Printf("changed value of age variable is %v\n", age)
	callY(amountPointer)
	callZ(someStringPointer)

	// All values in Go have a so-called "Null Value" - i.e., the value that's set as a default if no value is assigned to a variable.
	// For a pointer, it's nil - a special value built-into Go.
	// nil represents the absence of an address value - i.e., a pointer pointing at no address / no value in memory.

	var a int
	fmt.Printf("Value of a: %v\n", a)

	var b float64
	fmt.Printf("Value of b: %v\n", b)
	e := &b
	fmt.Printf("Value of b from pointer: %v\n", *e)

	var c string
	fmt.Printf("Value of c: %v & length: %v\n", c, len(c))

	var choice int
	fmt.Print("Your choice: ")
	// here we are using the & operator to get the address of choice variable from memory and scanning it
	fmt.Scan(&choice)

	fmt.Printf("Choice: %v\n", choice)

}

func callX(agePointer *int) {
	fmt.Printf("hi 1 -> %v\n", *agePointer)
	// to change the value, we need to dereference the pointer
	*agePointer = *agePointer + 1
}

func callY(amountPointer *float64) {
	fmt.Printf("hi 2 -> %v\n", *amountPointer)
}

func callZ(someStringPointer *string) {
	fmt.Printf("hi 3 -> %v\n", *someStringPointer)
}
