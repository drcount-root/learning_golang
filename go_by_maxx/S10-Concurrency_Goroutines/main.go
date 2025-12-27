package main

import (
	"fmt"
	"time"

	"example.com/concurrency_and_goroutines/combo"
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

	doneAddChannel := make(chan bool)

	go add(3, 7, doneAddChannel)

	<-doneAddChannel

	// better way to handle the above code using a slice of channels
	doneChannels := make([]chan bool, 2)

	doneChannels[0] = make(chan bool)
	go slowGreet2("6", doneChannels[0])

	doneChannels[1] = make(chan bool)
	go add(8, 9, doneChannels[1])

	for _, doneCh := range doneChannels {
		<-doneCh
	}

	// tax calc part using goroutines
	combo.Combo()
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

func add(a, b int, doneAddChannel chan bool) {
	time.Sleep(5 * time.Second)
	fmt.Printf("Sum after 5 seconds: %v\n", a+b)
	doneAddChannel <- true
}

/*
==================== Go Routines & Concurrency Notes ====================

1. Goroutines
-----------------------------------------------------------------------
- A goroutine is a lightweight, managed thread of execution.
- Created using the `go` keyword.
- Goroutines run concurrently, not necessarily in parallel.
- The Go runtime schedules goroutines, not the OS.
- Goroutines communicate using channels and share data. They don't return values.

Example:
    go slowGreet2("hello", doneChan)

IMPORTANT:
- Calling a function normally runs it synchronously.
- Using `go` runs the function asynchronously in a new goroutine.

-----------------------------------------------------------------------

2. Concurrency vs Parallelism
-----------------------------------------------------------------------
- Concurrency: structuring a program to handle multiple tasks at once.
- Parallelism: executing multiple tasks at the same time on multiple cores.
- Go focuses on concurrency; parallelism depends on available CPUs.

-----------------------------------------------------------------------

3. Channels
-----------------------------------------------------------------------
- Channels are used to communicate between goroutines.
- They synchronize execution and safely pass data.

Types:
    - Unbuffered channel: blocks until sender and receiver are both ready.
    - Buffered channel: allows limited non-blocking sends.

Example:
    done := make(chan struct{})

-----------------------------------------------------------------------

4. Signaling Completion (Synchronization)
-----------------------------------------------------------------------
- Use channels to signal when a goroutine has finished work.
- Prefer `chan struct{}` for signal-only channels (zero allocation).
- One channel should have one clear responsibility.

Example:
    done <- struct{}{}

Main goroutine waits using:
    <-done

-----------------------------------------------------------------------

5. Deadlocks
-----------------------------------------------------------------------
- A deadlock occurs when all goroutines are blocked and none can proceed.
- Common causes:
    - Sending on a channel with no receiver
    - Receiving from a channel with no sender
    - Circular waiting between goroutines

Go detects deadlocks at runtime and panics.

-----------------------------------------------------------------------

6. Best Practices
-----------------------------------------------------------------------
- Do not use channels without goroutines.
- Do not reuse the same channel for unrelated tasks.
- Separate data flow from lifecycle signaling.
- For waiting on multiple goroutines, prefer sync.WaitGroup.
- For cancellation and timeouts, use context.Context.

=======================================================================
*/
