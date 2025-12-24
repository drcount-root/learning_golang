package recursion

func RecursiveFactorialFunc(n *int) int {
	if *n == 0 {
		return 1
	}

	decrementedVal := (*n - 1)

	return *n * RecursiveFactorialFunc(&decrementedVal)
}

func RecursiveFibonacciFunc(n *int) []int {

	if *n <= 0 {
		return []int{}
	}

	if *n == 1 {
		return []int{0}
	}

	if *n == 2 {
		return []int{0, 1}
	}

	prevN := *n - 1
	fibSlice := RecursiveFibonacciFunc(&prevN)

	nextFib := fibSlice[len(fibSlice)-1] + fibSlice[len(fibSlice)-2]
	fibSlice = append(fibSlice, nextFib)

	return fibSlice
}
