package main

import (
    "fmt"
    "github.com/drcount-root/calculator_module/mathx"
	"github.com/drcount-root/calculator_module/caller"
)

func main() {
    sum := mathx.Add(10, 20)
    fmt.Println("Sum:", sum)

	result := caller.Caller(143, 121)
	fmt.Println("Result:", result)
}
