package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// wg := &sync.WaitGroup{}

	// wg.Add(2)

	// go readerFn(12, wg)

	// go sayHello(wg)

	// wg.Wait()

	wg := &sync.WaitGroup{}

	st := time.Now()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&i, wg)
	}

	wg.Wait()

	tt := time.Since(st)

	fmt.Printf("All workers finished their work! Time taken: %v\n", tt)
}

// func readerFn(n int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Reader %v\n", n)
// }

// func sayHello(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("sayHello Fn executed after 3 seconds")
// }

func worker(n *int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\nWorker %v started\n", *n)

	time.Sleep(3 * time.Second) // thing that some work is taking 3 secs to finish

	fmt.Printf("\nWorker %v ended\n", *n)
}
