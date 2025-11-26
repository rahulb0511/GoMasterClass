// =====================================================================
//  GO INTERFACES - MASTER INTERVIEW NOTES (MOST IMPORTANT CONCEPTS)
// =====================================================================
//  ✔ Interfaces are satisfied implicitly (no "implements")
//  ✔ A type implements an interface if it has all required methods
//  ✔ Interface value = (dynamic type + dynamic value)
//  ✔ Dynamic dispatch → correct method runs based on dynamic type
//  ✔ Empty interface → accepts any value
//  ✔ Type assertion → extract value from interface
//  ✔ Type switch → multiple type checks
//  ✔ Interface embedding → combine interfaces (like inheritance)
//  ✔ Nil interface trap → typed nil ≠ nil interface
// =====================================================================

package main

import "fmt"

func main() {

	// --------------------------------------------------------------
	// 1) BASIC INTERFACE IMPLEMENTATION (IMPLICIT IMPLEMENTATION)
	// --------------------------------------------------------------
	// Speaker interface ka koi bhi type tabhi implement karega
	// jab uske paas Speak() method ho.
	// Go me explicit "implements" keyword nahi hota.

	var s Speaker

	s = Person{Name: "Rahul"} // Person → Speak() present
	s.Speak()

	s = Dog{Breed: "Labrador"} // Dog → Speak() present
	s.Speak()

	// --------------------------------------------------------------
	// 2) DYNAMIC DISPATCH / POLYMORPHISM
	// --------------------------------------------------------------
	// Runtime me actual type ke method call hote hain (dynamic type)

	speakers := []Speaker{
		Person{"Amit"},
		Dog{"Husky"},
	}
	for _, sp := range speakers {
		sp.Speak() // Different output based on dynamic type
	}

	// --------------------------------------------------------------
	// 3) EMPTY INTERFACE (interface{}) → ANY TYPE ALLOWED
	// --------------------------------------------------------------
	// Go me koi bhi type interface{} me store ho sakta hai.

	var any interface{}
	any = 10
	fmt.Println(any)

	any = "GoLang"
	fmt.Println(any)

	// --------------------------------------------------------------
	// 4) TYPE ASSERTION → extract actual value
	// --------------------------------------------------------------
	var x interface{} = 200

	v, ok := x.(int)   // safe assertion
	fmt.Println(v, ok) // 200 true

	_, ok = x.(string) // wrong type
	fmt.Println(ok)    // false

	// --------------------------------------------------------------
	// 5) TYPE SWITCH → multiple type checks
	// --------------------------------------------------------------
	checkType(100)
	checkType("go")
	checkType(true)

	// --------------------------------------------------------------
	// 6) INTERFACE EMBEDDING → multiple behaviors combine
	// --------------------------------------------------------------

	var a Animal
	a = Dog{"Pug"} // Dog implements Speaker + Mover
	a.Speak()
	a.Move()

	// --------------------------------------------------------------
	// 7) INTERFACE INTERNALS → (dynamic type, dynamic value)
	// --------------------------------------------------------------

	var sp Speaker = Person{"Neha"}
	fmt.Printf("Dynamic Type=%T Dynamic Value=%v\n", sp, sp)

	// --------------------------------------------------------------
	// 8) NIL INTERFACE TRAP → IMPORTANT!
	// --------------------------------------------------------------
	// interface == nil ONLY when:
	//     dynamic type = nil
	//     dynamic value = nil

	var i interface{} = nil
	fmt.Println(i == nil) // true  → completely nil interface

	var spk Speaker = (*Person)(nil)
	fmt.Println(spk == nil) // false → type stored, value nil
}

// =====================================================================
// INTERFACE DEFINITIONS & IMPLEMENTING TYPES
// =====================================================================

// Speaker interface → 1 method requirement
type Speaker interface {
	Speak()
}

// Struct implementing Speaker
type Person struct {
	Name string
}

func (p Person) Speak() {
	fmt.Println("Person says: Hi, I am", p.Name)
}

// Dog also implements Speaker
type Dog struct {
	Breed string
}

func (d Dog) Speak() {
	fmt.Println("Dog barks: Woof! Breed:", d.Breed)
}

// =====================================================================
// TYPE SWITCH EXAMPLE
// =====================================================================

func checkType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("int =>", val)
	case string:
		fmt.Println("string =>", val)
	default:
		fmt.Println("unknown type")
	}
}

// =====================================================================
// INTERFACE EMBEDDING EXAMPLE (COMPOSITION)
// =====================================================================

// 1st interface
type Mover interface {
	Move()
}

// 2nd interface embedding another one
// Animal = Speaker + Mover
type Animal interface {
	Speaker
	Mover
}

// Dog implements Move() also
func (d Dog) Move() {
	fmt.Println("Dog is running...")
}

// =====================================================================
// END OF GO INTERFACES MASTER NOTES
// =====================================================================
