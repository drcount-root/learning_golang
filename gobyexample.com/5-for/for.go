package main

import "fmt"

func main(){
	a := 1

	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Print("----------\n")

	for b:= 1; b < 3; b++ {
		fmt.Println(b)
	}

	fmt.Print("----------\n")

	for c := range 3 {
		fmt.Println("range", c)
	}

	fmt.Print("----------\n")

	for {
		fmt.Println("hello")
		break
	}

	fmt.Print("----------\n")

	for d := range 10 {
		if d % 2 == 0 {
			continue
		}

		fmt.Println(d)
	}
}