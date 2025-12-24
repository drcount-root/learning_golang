package variadic_func

// VariadicFuncSum calculates the sum of a variable number of float64 numbers.
// It expects a variable number of pointers to float64 numbers.
// It treats these pointers, as a slice of float64 values.
// It returns the sum as a float64.

func VariadicFuncSum1(nums ...*float64) float64 {
	sum := 0.0

	for _, v := range nums {
		sum += *v
	}

	return sum
}

func VariadicFuncSum2(nums ...float64) float64 {
	sum := 0.0

	for _, v := range nums {
		sum += v
	}

	return sum
}

// VariadicFunc calculates the average of a variable number of float64 numbers.
// It expects a pointer to an integer 'count' representing the number of float64 pointers passed,
// followed by a variable number of pointers to float64 numbers. It treats these pointers, as a slice of float64 values.
// It returns the average as a float64.
func VariadicFuncAvg(count *int, nums ...*float64) float64 {
	sum := 0.0

	for _, v := range nums {
		sum += *v
	}

	ct := *count

	return sum / float64(ct)
}
