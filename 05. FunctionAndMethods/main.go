// =====================================================================
//  GO FUNCTIONS & METHODS - MASTER INTERVIEW NOTES (NO CONCEPT MISSED)
// =====================================================================
// Includes: basic funcs, params, return values, named return, variadic,
// multiple returns, anonymous funcs, closures, higher-order funcs,
// defer with functions, methods (value/pointer receivers), interfaces link.
// =====================================================================

package main

import "fmt"

func main() {

	// --------------------------------------------------------------
	// 1) BASIC FUNCTION CALL
	// --------------------------------------------------------------
	greet("Rahul")

	// --------------------------------------------------------------
	// 2) MULTIPLE PARAMETERS & RETURN VALUES
	// --------------------------------------------------------------
	sum, mul := addMultiply(3, 4)
	fmt.Println("sum:", sum, "mul:", mul)

	// --------------------------------------------------------------
	// 3) NAMED RETURNS
	// --------------------------------------------------------------
	r := divide(20, 4)
	fmt.Println("divide:", r)

	// --------------------------------------------------------------
	// 4) VARIADIC FUNCTIONS (...type)
	// --------------------------------------------------------------
	total := sumAll(1, 2, 3, 10)
	fmt.Println("sumAll:", total)

	nums := []int{5, 6, 7}
	fmt.Println(sumAll(nums...)) // spread slice

	// --------------------------------------------------------------
	// 5) ANONYMOUS FUNCTION
	// --------------------------------------------------------------
	func(msg string) {
		fmt.Println("Anonymous says:", msg)
	}("Hello")

	// --------------------------------------------------------------
	// 6) CLOSURE (function capturing outer variables)
	// --------------------------------------------------------------
	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// --------------------------------------------------------------
	// 7) FUNCTIONS AS VALUES / HIGHER ORDER FUNCTIONS
	// --------------------------------------------------------------
	operate := func(a, b int, fn func(int, int) int) int {
		return fn(a, b)
	}
	fmt.Println(operate(5, 6, func(x, y int) int { return x * y }))

	// --------------------------------------------------------------
	// 8) DEFER WITH FUNCTIONS
	// --------------------------------------------------------------
	defer fmt.Println("Runs at the END of main()")
	fmt.Println("Runs BEFORE deferred output")

	// --------------------------------------------------------------
	// 9) METHODS (Value & Pointer Receivers)
	// --------------------------------------------------------------
	p := Person{Name: "Amit", Age: 40}
	fmt.Println(p.Info())

	p.Birthday() // pointer receiver updates original
	fmt.Println(p.Info())

	// --------------------------------------------------------------
	// 10) POINTER RECEIVER BENEFITS
	// --------------------------------------------------------------
	// - Allows method to modify original struct
	// - Prevents copying large struct on method call

	// --------------------------------------------------------------
	// 11) METHODS ON NON-STRUCT TYPES
	// --------------------------------------------------------------
	var s MyString = "go-lang"
	fmt.Println(s.Upper())

	// --------------------------------------------------------------
	// INTERVIEW CHECKLIST
	// --------------------------------------------------------------
	// ✔ basic functions
	// ✔ multiple return values
	// ✔ named return values
	// ✔ variadic functions
	// ✔ anonymous & closure
	// ✔ functions as values / HOF
	// ✔ defer + functions
	// ✔ methods: value vs pointer receivers
	// ✔ methods on custom types
	// ✔ connection to interfaces
}

// =====================================================================
// FUNCTION DEFINITIONS
// =====================================================================

// simple function
func greet(name string) {
	fmt.Println("Hello", name)
}

// multiple returns
func addMultiply(a, b int) (int, int) {
	return a + b, a * b
}

// named return values
func divide(x, y int) (result int) {
	result = x / y
	return // implicit return
}

// variadic function
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// closure example
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// =====================================================================
// METHODS
// =====================================================================

type Person struct {
	Name string
	Age  int
}

// value receiver (does not modify original)
func (p Person) Info() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

// pointer receiver (modifies original)
func (p *Person) Birthday() {
	p.Age++
}

// =====================================================================
// METHODS ON NON-STRUCT TYPES
// =====================================================================

type MyString string

func (s MyString) Upper() string {
	return fmt.Sprint("UPPER:", string(s))
}

// =====================================================================
// END OF FUNCTIONS & METHODS MASTER NOTES
// =====================================================================
