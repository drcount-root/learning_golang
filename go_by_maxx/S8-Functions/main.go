package main

import (
	"fmt"

	"example.com/functions/anonymous_func"
	"example.com/functions/closure"
	"example.com/functions/recursion"
	"example.com/functions/variadic_func"
)

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

	anonymousFuncResult := anonymous_func.AnonymousFunc()
	fmt.Println(anonymousFuncResult)

	closureFnc := closure.ClosureDemonstrationFunc()
	closureFncResult := closureFnc(&sliceOfInts[3])
	fmt.Println(closureFncResult) // 10 + 3 = 13

	k := 6
	recursionResult := recursion.RecursiveFactorialFunc(&k)
	fmt.Printf("Fact of %v is : %v\n", k, recursionResult) // 720

	k = 0
	recursionResult = recursion.RecursiveFactorialFunc(&k)
	fmt.Printf("Fact of %v is : %v\n", k, recursionResult) // 1

	n := 8
	fibSlice := recursion.RecursiveFibonacciFunc(&n)
	fmt.Printf("Fibonacci series of %v is : %v\n", n, fibSlice) // [0 1 1 2 3 5 8 13]

	kx := []float64{7.8, 2.1, 5.3, 9.7}
	resVariadicFnSum1 := variadic_func.VariadicFuncSum1(&kx[0], &kx[1], &kx[2], &kx[3])
	fmt.Printf("Result of Variadic Function1 Sum is : %v\n", resVariadicFnSum1) // 24.9

	kx2 := []float64{7.8, 2.1, 5.3, 9.7}
	resVariadicFnSum2 := variadic_func.VariadicFuncSum2(kx2...)
	fmt.Printf("Result of Variadic Function2 Sum is : %v\n", resVariadicFnSum2) // 24.9

	kj := 6
	ki := []float64{7.8, 2.1, 5.3, 9.7}
	resVariadicFnAvg := variadic_func.VariadicFuncAvg(&kj, &ki[0], &ki[1], &ki[2], &ki[3])
	fmt.Printf("Result of Variadic Function Avg is : %.2f\n", resVariadicFnAvg) // 4.15

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
