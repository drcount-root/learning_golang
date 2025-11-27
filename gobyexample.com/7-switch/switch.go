package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()

	fmt.Println("Time t =", t)

	fmt.Println("Weekday", t.Weekday())

	fmt.Println("Time", t.Hour())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAMI := func(i interface{}) {
		switch p := i.(type) {
		case int:
			fmt.Println("I'm an int")

		case string:
			fmt.Println("I'm a string")

		case bool:
			fmt.Println("I'm a bool")

		default:
			fmt.Printf("Don't know type %T\n", p)
		}
	}

	whatAMI(true)
	whatAMI(1)
	whatAMI("hello")
	whatAMI(4.6)
}
