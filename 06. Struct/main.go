package main

import (
	"encoding/json"
	"fmt"
)

// --------------------------------------------------------------
// STRUCTS → Custom data types
// --------------------------------------------------------------
type User struct {
	Name string
	Age  int
}

// Constructor-like function → struct banane ka clean tarika
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

// Value receiver → struct ki copy par call hota hai (original change nahi hota)
func (u User) Description() string {
	return fmt.Sprintf("User: %s (%d)", u.Name, u.Age)
}

// Pointer receiver → original struct modify hota hai
func (u *User) IncrementAge() {
	u.Age++
}

// Embedded structs → Go ka inheritance-style (is-a nahi, has-a + promotes fields)
type Address struct {
	City string
	Zip  int
}

type Employee struct {
	user    User
	address Address
	Salary  int
}

func main() {

	// --------------------------------------------------------------
	// 1) STRUCT INITIALIZATION
	// --------------------------------------------------------------
	u1 := User{"Rahul", 25}           // positional
	u2 := User{Name: "Amit", Age: 26} // named fields
	var u3 User                       // zero-value struct

	fmt.Println(u1, u2, u3)

	// --------------------------------------------------------------
	// 2) UPDATE STRUCT FIELDS
	// --------------------------------------------------------------
	u2.Age = 30
	fmt.Println(u2)

	// --------------------------------------------------------------
	// 3) POINTER TO STRUCT (Go auto-dereferences pointer)
	// --------------------------------------------------------------
	u4 := &User{Name: "Mohit", Age: 28}
	u4.Age = 29 // *u4 likhne ki need nahi
	fmt.Println(u4)

	// --------------------------------------------------------------
	// 4) CONSTRUCTOR-STYLE CREATION
	// --------------------------------------------------------------
	newUser := NewUser("Sita", 40)
	fmt.Println(newUser)

	// --------------------------------------------------------------
	// 5) METHODS
	// --------------------------------------------------------------
	u5 := User{Name: "Ramesh", Age: 50}

	fmt.Println(u5.Description()) // value method → sirf read

	u5.IncrementAge() // pointer method → original u5 change hoga
	fmt.Println(u5)

	// --------------------------------------------------------------
	// 6) STRUCT EMBEDDING (Composition)
	// --------------------------------------------------------------
	u := User{"Komal", 33}
	emp := Employee{
		user:    u,                                   //m1
		address: Address{City: "Delhi", Zip: 110001}, //m2
		Salary:  80000,
	}

	fmt.Println(emp.user.Name, emp.address.City, emp.Salary)

	// --------------------------------------------------------------
	// 7) ANONYMOUS STRUCT (Quick one-time struct)
	// --------------------------------------------------------------
	anon := struct {
		Title string
		Likes int
	}{"Post Title", 500}

	fmt.Println(anon)

	// --------------------------------------------------------------
	// 8) STRUCT COMPARISON (field-by-field comparison allowed)
	// --------------------------------------------------------------
	p1 := User{"A", 1}
	p2 := User{"A", 1}
	fmt.Println(p1 == p2) // true (struct is comparable)

	// --------------------------------------------------------------
	// 9) JSON TAGS + MARSHAL/UNMARSHAL
	// --------------------------------------------------------------
	type Product struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Cost int    `json:"price"` // json field different name
	}

	prod := Product{ID: 1, Name: "Book", Cost: 299}

	// struct → JSON
	data, _ := json.Marshal(prod)
	fmt.Println(string(data))

	// JSON → struct
	var prod2 Product
	json.Unmarshal(data, &prod2)
	fmt.Println(prod2)

	// --------------------------------------------------------------
	// 10) ZERO VALUE RULE
	// --------------------------------------------------------------
	var empty User
	fmt.Printf("%+v\n", empty) // Name="", Age=0
}
