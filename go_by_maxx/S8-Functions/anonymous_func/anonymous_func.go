package anonymous_func

func AnonymousFunc() []int {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(n *int) int {
		return *n * 5
	})

	/*
		This is the anonymous function

		func(n *int) int {
			return *n * 5
		}
	*/

	return transformed
}

func transformNumbers(numbers *[]int, transform func(*int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(&val))
	}

	return dNumbers
}
