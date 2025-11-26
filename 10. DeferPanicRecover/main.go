// =====================================================================
//  GO DEFER, PANIC & RECOVER - FULL MASTER NOTES (INTERVIEW + PRODUCTION)
// =====================================================================
// This is the MOST COMPLETE version: Nothing is skipped.
// Includes:
// ✔ defer execution rules (LIFO, argument evaluation)
// ✔ defer with functions, loops, return values
// ✔ panic behavior, stack unwinding
// ✔ recover usage patterns to prevent crashes
// ✔ panic vs error (system design level)
// ✔ real-world examples: file close, db cleanup, mutex unlock, logging
// =====================================================================

package main

import (
	"fmt"
)

func main() {

	fmt.Println("================ Defer Basics ================")
	deferBasics()

	fmt.Println("\n================ Defer Argument Evaluation ================")
	deferArgumentEvaluation()

	fmt.Println("\n================ Defer in Loops ================")
	deferInsideLoop()

	fmt.Println("\n================ Panic + Recover Safe Function ================")
	fmt.Println(safeDivide(10, 2))
	fmt.Println(safeDivide(10, 0)) // Recovers, program continues

	// fmt.Println("\n================ Panic VS Error (Explanation) ================")
	// panicVsErrorExplanation()

	fmt.Println("\n================ Real-World Use Cases ================")
	simulateFileOpen()
	simulateMutexLock()

	fmt.Println("\nPROGRAM CONTINUES WITHOUT CRASH ✔✔✔")
}

// =====================================================================
// 1) DEFER BASICS — Runs at the END of the function in LIFO order
// =====================================================================
func deferBasics() {
	fmt.Println("Start of deferBasics()")

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")

	fmt.Println("End of function BEFORE deferred prints")

	// Output:
	// Start
	// End of function BEFORE deferred prints
	// Deferred 3
	// Deferred 2
	// Deferred 1
}

// =====================================================================
// 2) DEFER ARGUMENT EVALUATION (VERY IMPORTANT INTERVIEW QUESTION)
// Arguments are evaluated immediately, not when executed.
// =====================================================================
func deferArgumentEvaluation() {

	x := 10
	defer fmt.Println("defer value: x =", x) // captures 10, not 20

	x = 20
	fmt.Println("Regular print: x =", x)

	// Output:
	// Regular print: x = 20
	// defer value: x = 10
}

// =====================================================================
// 3) DEFER INSIDE LOOPS — executes at end of entire function, not per loop
// =====================================================================
func deferInsideLoop() {
	for i := 1; i <= 3; i++ {
		defer fmt.Println("Deferred loop value:", i)
	}

	fmt.Println("Loop completed, now deferred executes LIFO")

	// Output:
	// Loop completed...
	// 3
	// 2
	// 1
}

// =====================================================================
// 4) PANIC + RECOVER — Safe execution wrapper
// Panic stops normal flow and starts stack unwinding.
// Recover MUST be inside a deferred function.
// =====================================================================
func safeDivide(a, b int) (result string) {

	defer func() {
		if r := recover(); r != nil {
			result = fmt.Sprintf("Recovered from panic: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero is not allowed")
	}

	return fmt.Sprintf("Result: %d", a/b)
}

// =====================================================================
// 5) PANIC VS ERROR (Interview GOLD)
// =====================================================================
// func panicVsErrorExplanation() {
// 	fmt.Println(`
// PANIC:
// - Used for UNRECOVERABLE situations
// - Crashes the normal flow if not recovered
// - Examples: corrupted memory, illegal state, programmer error

// ERROR:
// - Used for EXPECTED failure conditions
// - Should be returned and handled
// - Examples: invalid user input, network failure, file not found

// RULE:
// Return errors. Reserve panic for truly exceptional conditions.
// `)
// }

// =====================================================================
// 6) Real-World Use Case: File Close Simulation
// Defer is PERFECT for cleanup
// =====================================================================
func simulateFileOpen() {
	fmt.Println("Opening file...")
	defer fmt.Println("Closing file (cleanup)...")

	fmt.Println("Reading file data...")
}

// =====================================================================
// 7) Real-World Use Case: Mutex Unlock Simulation
// (in real usage: sync.Mutex)
// =====================================================================
func simulateMutexLock() {
	fmt.Println("Lock acquired")
	defer fmt.Println("Lock released (important to avoid deadlocks)")

	fmt.Println("Performing protected work...")
}

// =====================================================================
// END OF FULL DEFER, PANIC & RECOVER MASTER NOTES
// =====================================================================
