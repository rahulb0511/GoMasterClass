// =====================================================================
//  GO CLOSURES - MASTER INTERVIEW NOTES (MOST IMPORTANT CONCEPTS)
// =====================================================================
// ✔ Closure = function + surrounding variables ka combo
// ✔ Outer variables ko "capture" karta hai
// ✔ Captured variables reference se bind hote hain (copy nahi!)
// ✔ Har bar closure call par updated value milti hai
// ✔ Function returning another function is the common closure pattern
// ✔ Used for: stateful functions, counters, generators, decorators
// ✔ Loop variable capture bug (common interview question!)
// ✔ Goroutines + closures need careful variable capture
// =====================================================================

package main

import "fmt"

func main() {

	// --------------------------------------------------------------
	// 1) BASIC CLOSURE (inner function outer variable use karta hai)
	// --------------------------------------------------------------

	x := 10

	// anonymous function using x → closure
	printX := func() {
		fmt.Println("x =", x)
	}

	printX() // x = 10

	x = 20
	printX() // x = 20 (closure sees updated x)

	// --------------------------------------------------------------
	// 2) FUNCTION RETURNING A CLOSURE (MOST ASKED IN INTERVIEWS)
	// --------------------------------------------------------------

	counter := newCounter()

	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	// --------------------------------------------------------------
	// 3) MULTIPLE CLOSURES HOLD SEPARATE STATES
	// --------------------------------------------------------------

	c1 := newCounter()
	c2 := newCounter()

	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c2()) // 1 (separate state)

	// --------------------------------------------------------------
	// 4) LOOP VARIABLE CAPTURE (BUG + FIX)
	// --------------------------------------------------------------

	fmt.Println("\nWrong closure in loop:")
	wrongLoopClosure()

	fmt.Println("\nCorrect closure in loop:")
	correctLoopClosure()

	// --------------------------------------------------------------
	// 5) CLOSURE + GOROUTINES (MUST FIX LOOP VARIABLE)
	// --------------------------------------------------------------

	fmt.Println("\nClosure with goroutine (correct handling):")
	goroutineClosureFix()
}

// =====================================================================
// CLOSURE FACTORY – returns a function maintaining state
// =====================================================================

func newCounter() func() int {
	count := 0

	return func() int { // ← closure capturing 'count'
		count++
		return count
	}
}

// =====================================================================
// LOOP VARIABLE CAPTURE BUG
// =====================================================================
// For loop ka i variable ek hi memory location hota hai.
// Sare closures same i ko refer karte hain → buggy output.
// =====================================================================

func wrongLoopClosure() {
	funcs := []func(){}

	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println("i =", i)
		})
	}

	for _, f := range funcs {
		f() // all print 3!
	}
}

// =====================================================================
// FIX: Create new local variable inside loop
// =====================================================================

func correctLoopClosure() {
	funcs := []func(){}

	for i := 0; i < 3; i++ {
		j := i // NEW variable → each closure binds independent copy
		funcs = append(funcs, func() {
			fmt.Println("j =", j)
		})
	}

	for _, f := range funcs {
		f() // 0,1,2 (correct)
	}
}

// =====================================================================
// GOROUTINE CLOSURE FIX (must copy loop variable!)
// =====================================================================

func goroutineClosureFix() {
	done := make(chan bool)

	for i := 0; i < 3; i++ {
		ii := i // fix
		go func() {
			fmt.Println("goroutine sees:", ii)
			done <- true
		}()
	}

	// wait for all
	for i := 0; i < 3; i++ {
		<-done
	}
}

// =====================================================================
// END OF CLOSURES MASTER NOTES
// =====================================================================
