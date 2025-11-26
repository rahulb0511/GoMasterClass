// ============================================================================
//  GO CHANNELS - MASTER INTERVIEW NOTES (COMPLETE + INTERVIEW READY)
// ============================================================================
// Channels allow goroutines to communicate safely without explicit locks.
// Think of them as typed message queues with blocking semantics.
// ============================================================================

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("================ UNBUFFERED CHANNEL ================")
	unbufferedChannelDemo()

	fmt.Println("\n================ BUFFERED CHANNEL ================")
	bufferedChannelDemo()

	fmt.Println("\n================ CHANNEL CLOSE & RANGE ================")
	closeAndRangeDemo()

	fmt.Println("\n================ DIRECTIONAL CHANNELS ================")
	directionalChannelsDemo()

	fmt.Println("\n================ SELECT STATEMENT ================")
	selectDemo()

	fmt.Println("\n================ TIMEOUT WITH SELECT ================")
	timeoutSelectDemo()

	fmt.Println("\n================ FAN-IN PATTERN ================")
	fanInDemo()

	fmt.Println("\n================ FAN-OUT (WORKERS) ================")
	fanOutWorkerDemo()

	fmt.Println("\nCHANNELS MASTER DEMO COMPLETE ✔✔✔")
}

// ============================================================================
// 1) UNBUFFERED CHANNEL - Send blocks until a receiver is ready
// ============================================================================
// * Unbuffered channels = synchronous communication.
// * Sending will blocks until someone receives.
// * Receiving will blocks until someone sends.
// * Used for strict ordering and synchronization.
// ============================================================================
func unbufferedChannelDemo() {
	ch := make(chan string) // unbuffered = capacity 0

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- "Hello from Goroutine" // BLOCKS until main receives
	}()

	msg := <-ch // waits here (main blocks until goroutine sends)
	fmt.Println("Received:", msg)
}

// ============================================================================
// 2) BUFFERED CHANNEL - Allows queueing without blocking
// ============================================================================
// * Buffered channels = asynchronous up to buffer capacity.
// * Send blocks only when buffer is full.
// * Receive blocks only when buffer is empty.
// ============================================================================
func bufferedChannelDemo() {
	ch := make(chan int, 2) // buffer size = 2

	ch <- 10 // does NOT block
	ch <- 20 // does NOT block (buffer full now)

	// Now receives
	fmt.Println(<-ch, <-ch)
}

// ============================================================================
// 3) CLOSE CHANNEL + RANGE
// ============================================================================
// * close(ch) signals "no more values will be sent".
// * After close, receives continue until buffer is drained.
// * `range ch` automatically stops when channel closes.
// * Sending on closed channel => PANIC (common interview mistake).
// ============================================================================
func closeAndRangeDemo() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // required for ranging (otherwise deadlock)
	}()

	for val := range ch { // auto stops when channel closes
		fmt.Println("Received:", val)
	}
}

// ============================================================================
// 4) DIRECTIONAL CHANNELS (chan<- send only, <-chan receive only)
// ============================================================================
// Why directional?
// * API safety — prevents misuse
// * Makes intention clear
// * Compiler enforces correct communication direction
// ============================================================================
func producer(out chan<- int) { // send-only
	for i := 1; i <= 3; i++ {
		out <- i
	}
	close(out)
}

func consumer(in <-chan int) { // receive-only
	for v := range in {
		fmt.Println("Consumed:", v)
	}
}

func directionalChannelsDemo() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

// ============================================================================
// 5) SELECT STATEMENT - Multiplex multiple channels
// ============================================================================
// * select waits until one of the cases is ready.
// * Equivalent to non-deterministic wait.
// * Used for: fan-in, timeouts, cancellation, multiplexing.
// * If multiple cases are ready => Go picks one RANDOMLY.
// ============================================================================
func selectDemo() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(400 * time.Millisecond)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1: // whichever sends first
			fmt.Println("Received:", msg)
		case msg := <-ch2:
			fmt.Println("Received:", msg)
		}
	}
}

// ============================================================================
// 6) TIMEOUT WITH SELECT
// ============================================================================
// VERY COMMON INTERVIEW QUESTION.
// time.After() returns a channel that sends a value after duration.
// Perfect for API timeouts, DB timeouts, Goroutine watchdogs.
// ============================================================================
func timeoutSelectDemo() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "slow response"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second): // timeout case
		fmt.Println("Timeout! no response in time")
	}
}

// ============================================================================
// 7) FAN-IN PATTERN (combine multiple channels into one)
// ============================================================================
// * Aggregates results from multiple goroutines.
// * select ensures we receive from whichever channel is ready first.
// * Setting a channel to nil DISABLES that case in select.
// ============================================================================
func fanIn(input1, input2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		for input1 != nil || input2 != nil {
			select {
			case v, ok := <-input1:
				if !ok { // channel closed
					input1 = nil // disable case
					continue
				}
				out <- v
			case v, ok := <-input2:
				if !ok {
					input2 = nil
					continue
				}
				out <- v
			}
		}
	}()
	return out
}

func fanInDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 100; i <= 102; i++ {
			ch2 <- i
		}
	}()

	for v := range fanIn(ch1, ch2) {
		fmt.Println("Fan-in received:", v)
	}
}

// ============================================================================
// 8) FAN-OUT / WORKER POOL (simple version)
// ============================================================================
// * Multiple workers reading from SAME jobs channel.
// * Load automatically distributed (round-robin-ish).
// * Classic interview question for backend engineers.
//
// Notes:
// - Closing `jobs` tells workers no more work.
// - Workers exit NATURALLY when jobs channel closes.
// - `results` collects outputs (fan-in).
// ============================================================================
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // runs until jobs channel closes
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(200 * time.Millisecond) // simulating work
		results <- j * 2
	}
}

func fanOutWorkerDemo() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Spawn 2 workers (fan-out)
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	// send data to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // very important (otherwise workers block forever)

	// Collect results (fan-in)
	for r := 1; r <= numJobs; r++ {
		fmt.Println("Result:", <-results)
	}
}

// ============================================================================
// INTERVIEW & PRODUCTION CHECKLIST
// ============================================================================
// ✔ Channel = safe communication between goroutines
// ✔ Unbuffered = synchronous, Buffered = queue
// ✔ close() required for ranging
// ✔ Send on closed channel = PANIC
// ✔ Directional channels = safer APIs
// ✔ Select = multiplex multiple channels
// ✔ Timeout = select + time.After
// ✔ Fan-in = merge channels
// ✔ Fan-out = distribute work across goroutines
// ✔ Worker pool = standard backend concurrency pattern
// ============================================================================
