package main

import (
	"fmt"
	"time"
)

func main() {
	greet("1")
	greet("2")
	slowGreet("3")
	greet("4")

	// doneChannel is an unbuffered channel used to signal completion between goroutines. It is used to notify the main goroutine that the slowGreet function has finished.
	doneChannel := make(chan bool)

	go slowGreet2("5", doneChannel) // running slowGreet2 in a goroutine

	<-doneChannel // wait for the data & ends the program
}

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
}

func slowGreet2(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)

	// the data true is flowing in the channel doneChan
	doneChan <- true
}
