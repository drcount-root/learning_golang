package main

import "fmt"

type customStr string

func (text customStr) log() {
	fmt.Println(text)
}

func main() {
	var description customStr = "This is a custom string type"

	description.log()
}
