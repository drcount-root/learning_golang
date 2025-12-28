# Process of setup & execution

## To create a go module

`go mod init example.com/hello`

## To install modules

`go mod tidy`

## To run

`go run .`

go run main.go

go build main.go

./main

---

## Notes for Packages

### Internal Packages

In go package name (fileops) & folder name (fileops) should be same, filename can be different

### To install external package

`go get github.com/Pallinder/go-randomdata`

---

## Go is not a dynamically typed language. It is a statically typed with type inference

Let’s frame it precisely, the way a system architect thinks:

1. Statically Typed

Go determines the type of every variable at compile time, not at runtime.
Even if you use:

x := 10

Go infers the type int, and after that, x is strictly an int.
You cannot later do:

x = "hello" // compile-time error

2. Type Inference

The := syntax lets the compiler infer the type so you don’t have to write it:

name := "Bro" // inferred as string
pi := 3.14 // inferred as float64

The key detail:
Type inference does not make Go dynamic; it only removes verbosity.

3. How it differs from dynamically typed languages

Languages like Python and JavaScript decide variable types at runtime.
You can do:

x = 10
x = "hello"

In Go, this is impossible. Types are fixed once assigned.

4. Go has some dynamic behavior via interfaces

If you use interface{}, Go can hold values of any type, but even then, the actual type is preserved internally.
This is controlled dynamic capability, not dynamic typing.

Final summary:

Go = statically typed + type inference + limited dynamic features via interfaces

# Go Concurrency: Complete Practical Example

The example models a **concurrent task processing system**, similar to what you’d build in backend services, job queues, or data pipelines.

---

## 🎯 Problem Statement

We want to process a list of jobs concurrently while ensuring:

- Controlled concurrency (worker pool)
- Proper error handling
- Graceful shutdown
- Timeout support
- No goroutine leaks
- Clear ownership of channels

---

## 🧰 Go Concurrency Features Used

| Feature           | Purpose                     |
| ----------------- | --------------------------- |
| Goroutines        | Concurrent execution        |
| Channels          | Safe communication          |
| Buffered channels | Backpressure                |
| `select`          | Multiplexing + cancellation |
| `context.Context` | Timeouts & cancellation     |
| `sync.WaitGroup`  | Lifecycle coordination      |
| Fan-out / Fan-in  | Worker pool pattern         |

---

## 🧠 High-Level Architecture
```

Job IDs
↓
Jobs Channel
↓
Worker Pool (Goroutines)
↓
Results Channel
↓
Aggregator (main goroutine)

````

---

## 📦 Data Model

```go
type Result struct {
	JobID int
	Err   error
}
````

Each job returns **exactly one result**:

- `Err == nil` → success
- `Err != nil` → failure

---

## 👷 Worker Implementation

```go
func worker(
	ctx context.Context,
	id int,
	jobs <-chan int,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// graceful shutdown
			return

		case jobID, ok := <-jobs:
			if !ok {
				return
			}

			// simulate work
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

			if jobID%5 == 0 {
				results <- Result{JobID: jobID, Err: errors.New("job failed")}
			} else {
				results <- Result{JobID: jobID, Err: nil}
			}
		}
	}
}
```

### Key Notes

- Workers **own no channels**
- Workers exit cleanly on:

  - context cancellation
  - closed jobs channel

- Each job sends **exactly one result**

---

## 🚀 Main Function (Orchestration)

```go
func main() {
	rand.Seed(time.Now().UnixNano())

	jobIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	jobs := make(chan int)
	results := make(chan Result)

	const workerCount = 3
	var wg sync.WaitGroup

	// fan-out: start workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &wg)
	}

	// producer
	go func() {
		defer close(jobs)
		for _, jobID := range jobIDs {
			select {
			case <-ctx.Done():
				return
			case jobs <- jobID:
			}
		}
	}()

	// fan-in: close results when workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// aggregate results
	for res := range results {
		if res.Err != nil {
			fmt.Printf("❌ Job %d failed: %v\n", res.JobID, res.Err)
		} else {
			fmt.Printf("✅ Job %d completed\n", res.JobID)
		}
	}

	fmt.Println("All processing complete.")
}
```

---

## 🔒 Concurrency Guarantees

- ✔ No goroutine leaks
- ✔ No deadlocks
- ✔ Bounded concurrency
- ✔ Deterministic shutdown
- ✔ Clear channel ownership
- ✔ Context-aware cancellation

---

## 🧠 Design Principles Demonstrated

- **One goroutine → one channel → one result**
- Channels are closed by **owners only**
- `select` is used for **control flow**, not decoration
- Context is the **control plane**
- WaitGroup is the **lifecycle guard**

---

## 🏗 Where This Pattern Is Used

- Background job processors
- API fan-out/fan-in
- Web crawlers
- ETL pipelines
- Event consumers
- Microservices orchestration

---

## 🎯 Final Takeaway

Go concurrency is not about “running things in parallel”.
It is about **coordinating work safely, predictably, and transparently**.

