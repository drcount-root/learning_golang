package main

import "fmt"

func main(){
	for i:=0; i < 10; i++ {
		if i%2 == 0 {
			if i < 5 {
				fmt.Println(i, "is even and less than 5")
			} else {
				fmt.Println(i, "is even and greater than 5")
			}
		} else {
			fmt.Println(i, "is odd")
		}
	}
}