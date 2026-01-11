package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go readerFn(12, wg)

	go sayHello(wg)

	wg.Wait()
}

func readerFn(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Reader %v\n", n)
}

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Println("sayHello Fn executed after 3 seconds")
}
