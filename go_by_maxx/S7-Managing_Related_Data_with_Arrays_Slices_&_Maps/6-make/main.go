package main

import "fmt"

func main() {
	userNames := []string{}
	userNames = append(userNames, "Max")
	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")
	fmt.Println(userNames)

	// use of make
	// make solves the problem of having to specify the size of an array or slice when it is declared. It allows us to create a slice of a certain size without having to populate it with values.

	// type, preintended length, capacity
	userNames2 := make([]string, 3) // Create a slice of length 3 backed by an underlying array of length 3.
	userNames2[0] = "Chris"
	userNames2[1] = "David"
	userNames2[2] = "Kate"
	// userNames2[3] = "Jade" // panic: runtime error: index out of range [3] with length 3

	/*
		for normal slics we cant assign the values using index

		usrNames := []string{}
		usrNames[0] = "max"

		panic: runtime error: index out of range [0] with length 0
	*/

	fmt.Println(userNames2)

	// useName := map [string]int {}

	// useName["CPU Cores"] = 8
	// useName["RAM"] = 16
	// useName["SSD"] = 512
	// fmt.Println(useName)

	// delete(useName, "SSD")
	// fmt.Println(useName)

	// type, preintended length
	courseRatings := make(map [string]float64, 3)

	courseRatings["Go"] = 9.5
	courseRatings["Python"] = 9.0
	courseRatings["JavaScript"] = 9.5
	courseRatings["Java"] = 8.7
	fmt.Println(courseRatings) // map[Go:9.5 Java:8.7 JavaScript:9.5 Python:9]

	type floatMap map [string] float64

	newCourseRatings := make(floatMap, 3)

	newCourseRatings["Go"] = 9.5
	newCourseRatings["Python"] = 9.0
	newCourseRatings["JavaScript"] = 9.5
	newCourseRatings["Java"] = 8.7
	fmt.Println(newCourseRatings)
}
