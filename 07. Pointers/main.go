// =====================================================================
//  GO POINTERS - MASTER INTERVIEW NOTES (NO CONCEPT MISSED)
// =====================================================================
// These notes cover 100% pointer concepts for Go developer interviews.
// With runnable examples & deep explanations.
// =====================================================================

package main

import "fmt"

func main() {
	// --------------------------------------------------------------
	// 1) WHAT IS A POINTER?
	// --------------------------------------------------------------
	// A pointer stores the MEMORY ADDRESS of a value.

	x := 10
	var p *int = &x // & gives address of x

	fmt.Println("x =", x)
	fmt.Println("address of x =", &x)
	fmt.Println("p (pointer value) =", p)
	fmt.Println("*p (value at pointer) =", *p) // dereference

	// Changing value using pointer
	*p = 99
	fmt.Println("x after pointer update =", x) // 99

	// --------------------------------------------------------------
	// 2) SHORT POINTER DECLARATION
	// --------------------------------------------------------------
	y := 200
	py := &y
	fmt.Println(*py) // 200

	// --------------------------------------------------------------
	// 3) NIL POINTER
	// --------------------------------------------------------------
	var np *int
	fmt.Println(np)        // <nil>
	fmt.Println(np == nil) // true

	// Dereferencing nil pointer causes panic
	// fmt.Println(*np) // ❌ panic: invalid memory address

	// --------------------------------------------------------------
	// 4) POINTERS & FUNCTIONS (PASS BY REFERENCE)
	// --------------------------------------------------------------
	val := 50
	fmt.Println("Before:", val)
	updateValue(&val)          // sending address
	fmt.Println("After:", val) // 100

	// --------------------------------------------------------------
	// 5) POINTERS WITH STRUCTS (VERY IMPORTANT)
	// --------------------------------------------------------------
	person := Person{name: "Rahul", age: 26}
	modifyPerson(&person)
	fmt.Println(person) // age updated

	// --------------------------------------------------------------
	// 6) NEW() vs & LITERAL
	// --------------------------------------------------------------
	// new(type) allocates memory and returns a pointer
	p1 := new(int)
	fmt.Println(*p1) // 0 default
	*p1 = 42
	fmt.Println(*p1) // 42

	// literal way
	p2 := &Person{name: "Amit", age: 30}
	fmt.Println(p2)

	// --------------------------------------------------------------
	// 7) POINTER TO POINTER (RARE BUT ASKED)
	// --------------------------------------------------------------
	a := 10
	pA := &a
	ppA := &pA         // pointer to pointer
	fmt.Println(**ppA) // 10

	// --------------------------------------------------------------
	// 8) ARRAYS vs SLICES WITH POINTERS
	// --------------------------------------------------------------
	arr := [3]int{1, 2, 3}
	arrP := &arr // pointer to array
	(*arrP)[0] = 100
	fmt.Println(arr) // [100 2 3]

	// Slice already works with reference semantics
	sl := []int{1, 2, 3}
	slP := sl
	slP[0] = 500
	fmt.Println(sl) // [500 2 3]

	// --------------------------------------------------------------
	// 9) POINTERS + METHODS (VERY IMPORTANT)
	// --------------------------------------------------------------
	u := User{name: "Allen", balance: 1000} //User ke phle & ni doge to bhi chlega, automatically go handle krleta hai
	u.AddBalance(500)                       // pointer receiver method
	fmt.Println(u.balance)                  // 1500
}

// function to modify int using pointer\ n
func updateValue(p *int) {
	*p = 100
}

// struct example
type Person struct {
	name string
	age  int
}

func modifyPerson(p *Person) {
	//(*p).age = 30 // it will also work
	p.age = 30. //go automatically dereference krdeta hai, yahan pr (*p).age likhne k jrurt ni h
}

// method with pointer receiver
type User struct {
	name    string
	balance int
}

func (u *User) AddBalance(amount int) {
	u.balance += amount
}

// =====================================================================
// END OF POINTERS MASTER NOTES
// =====================================================================
