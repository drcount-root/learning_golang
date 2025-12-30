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

# Why Concurrency Is Important (With Examples)

Concurrency is the ability of a program to **make progress on multiple tasks independently**.  
It allows a system to stay productive while some tasks are blocked waiting for external resources like network, disk, or time.

Most real systems spend more time **waiting** than **computing**.  
Concurrency exists to prevent that waiting from becoming wasted time.

---

## The Core Problem: Blocking Execution

### Sequential (Non-Concurrent) Example

```go
func main() {
	fetchUser()
	fetchOrders()
	fetchPayments()
}
```

If each function:

- Makes a network request
- Suppose each of them takes ~2 seconds

**Total execution time:** ~6 seconds

Each call blocks the next one.
The CPU is idle while waiting for I/O.

---

## Concurrency Solves This Inefficiency

### Concurrent Version

```go
func main() {
	go fetchUser()
	go fetchOrders()
	go fetchPayments()

	time.Sleep(3 * time.Second) // keep main alive
}
```

Now:

- All three requests start immediately
- They wait independently
- CPU stays available

**Total execution time:** ~2 seconds

Same work.
Same logic.
Dramatically better utilization.

---

## Example 1: Web Server Handling Requests

### Without Concurrency (Conceptual)

```go
func handleRequest() {
	readFromDB()   // blocks
	callAPI()      // blocks
	writeResponse()
}
```

If the server processes **one request at a time**:

- Every user waits for the previous user
- Throughput collapses under load

---

### With Concurrency (How Go Actually Works)

```go
func handler(w http.ResponseWriter, r *http.Request) {
	go readFromDB()
	go callAPI()

	writeResponse()
}
```

Each request:

- Runs independently
- Does not block other users
- Scales naturally with traffic

This is why Go servers can handle **thousands of concurrent connections** efficiently.

---

## Example 2: Background Work Without Freezing the Main Flow

### Problem

You want to:

- Process a file
- Log progress
- Send metrics
- Keep the app responsive

### Sequential Approach (Bad UX)

```go
processFile()
logMetrics()
sendAnalytics()
```

User waits for everything.

---

### Concurrent Approach

```go
go logMetrics()
go sendAnalytics()

processFile()
```

- Core task stays fast
- Secondary tasks run in the background
- System feels responsive

This pattern is everywhere in production systems.

---

## Example 3: Time-Based Waiting

### Sequential Waiting

```go
fmt.Println("Start")
time.Sleep(2 * time.Second)
fmt.Println("Task A done")

time.Sleep(4 * time.Second)
fmt.Println("Task B done")
```

**Total time:** 6 seconds

---

### Concurrent Waiting

```go
fmt.Println("Start")

go func() {
	time.Sleep(2 * time.Second)
	fmt.Println("Task A done")
}()

go func() {
	time.Sleep(4 * time.Second)
	fmt.Println("Task B done")
}()

time.Sleep(5 * time.Second)
```

**Total time:** ~4 seconds

Concurrency allows **time itself** to be overlapped.

---

## Concurrency vs Parallelism

- **Concurrency**: Structuring a program to handle multiple tasks at once
- **Parallelism**: Executing multiple tasks at the exact same time

A program can be concurrent without being parallel.

Go lets you design for concurrency while the runtime decides:

- How many OS threads to use
- How tasks are scheduled
- How CPU cores are utilized

---

## Why Go’s Concurrency Model Matters

Go avoids common pitfalls of traditional threading:

- Threads are heavy
- Shared memory is dangerous
- Locks are error-prone

Instead, Go promotes:

- Goroutines (lightweight execution units)
- Channels (safe communication)
- Structured concurrency

This makes concurrent code:

- Easier to reason about
- Easier to test
- Easier to scale

---

## The Cost of Avoiding Concurrency

Applications that avoid concurrency often suffer from:

- Poor performance under load
- Blocking bottlenecks
- Complex refactors later
- Over-provisioned infrastructure

Concurrency is cheapest when built **from the start**.

---

## Final Takeaway

Concurrency is important because:

- Real systems wait more than they compute
- Waiting should not block progress
- Scalability depends on independent execution
- Go makes concurrency safe and practical

Concurrency is not about speed.
It is about **control over time, resources, and complexity**.

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

```

---

## 📦 Data Model

```go
type Result struct {
	JobID int
	Err   error
}
```

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

---

## Preparing Statements vs Directly Executing Queries (Prepare() vs Exec()/Query())

DB.Exec() (when we created the tables)

Prepare() + stmt.Exec() (when we inserted data into the database)

DB.Query() (when we fetched data from the database)

Using Prepare() is 100% optional! You could send all your commands directly via Exec() or Query().

The difference between those two methods then just is whether you're fetching data from the database (=> use Query()) or your manipulating the database / data in the database (=> use Exec()).

But what's the advantage of using Prepare()?

Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).

This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.

And, indeed, in this application, we are calling stmt.Close() directly after calling stmt.Exec(). So here, it really wouldn't matter which approach you're using.
