package main

import (
	"fmt"
	"sync"
)

func main() {

	/* ------------------ approach 1 ------------------ */

	var wg sync.WaitGroup

	wg.Add(1) // We need to wait for one goroutine, so we Add(1) *before* starting it.

	go func() {
		// defer Done() right away to ensure it's called, even if the goroutine panics or returns early.
		defer wg.Done()
		fmt.Println("Hello!")
	}()

	go func() {
		fmt.Println("Hey")
	}()

	go func() {
		fmt.Println("Hey 2")
	}()

	wg.Wait() // Wait blocks until the counter is zero.

	/* ------------------ approach 2 ------------------ */

	// wg := &sync.WaitGroup{}

	// wg.Add(1) // We need to wait for one goroutine, so we Add(1) *before* starting it.

	// go func() {
	// 	defer wg.Done() // defer Done() right away to ensure it's called, even if the goroutine panics or returns early.

	// 	fmt.Println("Hello!")
	// }()

	// wg.Wait() // Wait blocks until the counter is zero.

	/* ------------------ best practice ------------------ */

	// wg := &sync.WaitGroup{}

	// wg.Go(func() {
	// 	fmt.Println("Hello!")
	// })

	// wg.Wait()
}

/*
	Problem is that it doesn't prints anything

	Question is WHY?
	=> main func finishes and it kills the program before the goroutine has a chance to run
*/
