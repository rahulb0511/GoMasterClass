// ============================================================================
//  GO CONCURRENCY MASTER NOTES (GOROUTINE, WAITGROUP, MUTEX) - COMPLETE
// ============================================================================
// Covers EVERYTHING needed for Go interviews + real backend development:
// ✔ Goroutines (lightweight concurrency)
// ✔ WaitGroup (sync goroutines)
// ✔ Mutex & RWMutex (data race prevention)
// ✔ Data race example + fix
// ✔ Best practices & interview checklist
// ============================================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("================ GOROUTINES ================")
	goroutineExample()

	fmt.Println("\n================ WAITGROUP ================")
	waitGroupExample()

	fmt.Println("\n================ DATA RACE DEMO ================")
	dataRaceDemo()

	fmt.Println("\n================ FIX WITH MUTEX ================")
	mutexFixDemo()

	fmt.Println("\n================ RWMutex Demo ================")
	rwMutexDemo()

	fmt.Println("\nALL CONCURRENCY DEMOS COMPLETED ✔✔✔")
}

// ============================================================================
// 1) Goroutines - lightweight threads managed by Go runtime
// ============================================================================
// - Use `go` keyword to start concurrent execution
// - Extremely cheap: can spawn thousands
// - Does NOT block main() unless waited
// ============================================================================
func goroutineExample() {
	go fmt.Println("Hello from Goroutine 🌀")
	fmt.Println("Hello from main thread")
	time.Sleep(100 * time.Millisecond) // wait so goroutine can finish
}

// ============================================================================
// 2) WAITGROUP - wait for goroutines to complete
// ============================================================================
// Steps:
//
//	var wg sync.WaitGroup
//	wg.Add(n)
//	go func(){ defer wg.Done() }()
//	wg.Wait()
//
// ============================================================================
func waitGroupExample() {
	var wg sync.WaitGroup

	wg.Add(3) // wait for 3 goroutines

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Println("Worker", id, "started")
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Worker", id, "finished")
		}(i)
	}

	wg.Wait() // blocks until all Done()
	fmt.Println("All workers completed 🎉")
}

// ============================================================================
// 3) DATA RACE EXAMPLE (unsafe concurrent write)
// ============================================================================
// To detect race: `go run -race main.go`
// ============================================================================
func dataRaceDemo() {
	count := 0
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			count++ // RACE CONDITION ❌
		}()
	}

	wg.Wait()
	fmt.Println("Final Count (RACE, unpredictable):", count)
}

// ============================================================================
// 4) FIX WITH MUTEX (Mutual Exclusion)
// ============================================================================
// - sync.Mutex: Lock/Unlock for write protection
// ============================================================================
func mutexFixDemo() {
	count := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final Count (SAFE):", count)
}

// ============================================================================
// 5) RWMutex - allows multiple readers but one writer
// ============================================================================
// - RLock/RUnlock for reading
// - Lock/Unlock for writes
// ============================================================================
func rwMutexDemo() {

	var wg sync.WaitGroup
	var rw sync.RWMutex
	shared := 0

	// Start writers
	wg.Add(1)
	go func() {
		defer wg.Done()
		rw.Lock()
		shared = 42
		fmt.Println("Writer updated value ->", shared)
		rw.Unlock()
	}()

	// Start readers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rw.RLock()
			fmt.Println("Reader", id, "read value ->", shared)
			rw.RUnlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("RWMutex demo done ✔")
}

// ============================================================================
// INTERVIEW & PRODUCTION CHECKLIST
// ============================================================================
// ✔ Why goroutines are lightweight? (stack grows, multiplexed, scheduler)
// ✔ Why race conditions happen? (shared write without sync)
// ✔ When to use WaitGroup?
// ✔ Mutex vs RWMutex (read-heavy systems use RWMutex)
// ✔ Why channels vs Mutex? (communication vs shared memory)
// ✔ `go run -race` debugging
// ✔ NEVER: share state without sync
// ============================================================================
