package main

import "fmt"

func main() {

	// --------------------------------------------------------------
	// 1) BASIC MAP DECLARATION
	// --------------------------------------------------------------
	// Maps are reference types: map[keyType]valueType

	var m1 map[string]int  // nil map (not usable for assignment yet)
	fmt.Println(m1 == nil) // true

	// Before using, initialize it using make()
	m1 = make(map[string]int)
	m1["a"] = 10
	fmt.Println(m1)

	// Short declaration (recommended)
	m2 := map[string]int{"x": 100, "y": 200}
	fmt.Println(m2)

	// --------------------------------------------------------------
	// 2) INSERT / UPDATE VALUES
	// --------------------------------------------------------------
	m2["y"] = 250 // update
	m2["z"] = 300 // insert
	fmt.Println(m2)

	// --------------------------------------------------------------
	// 3) ACCESS VALUE (with existence check)
	// --------------------------------------------------------------
	value, ok := m2["x"]
	if ok {
		fmt.Println("x exists with value", value)
	}

	// key not present
	v, ok := m2["p"]
	fmt.Println(v, ok) // 0 false (zero-value + false)

	// --------------------------------------------------------------
	// 4) DELETE A KEY
	// --------------------------------------------------------------
	delete(m2, "y")
	fmt.Println(m2)

	// --------------------------------------------------------------
	// 5) LENGTH OF MAP
	// --------------------------------------------------------------
	fmt.Println("len:", len(m2))

	// --------------------------------------------------------------
	// 6) LOOPING A MAP (order not guaranteed!)
	// --------------------------------------------------------------
	for key, val := range m2 {
		fmt.Println(key, val)
	}

	// --------------------------------------------------------------
	// 7) MAP OF STRUCTS
	// --------------------------------------------------------------
	type User struct {
		Name string
		Age  int
	}

	users := map[int]User{
		1: {Name: "Rahul", Age: 25},
		2: {Name: "Amit", Age: 30},
	}

	fmt.Println(users[1].Name)

	// --------------------------------------------------------------
	// 8) SLICE OF MAPS
	// --------------------------------------------------------------
	sm := make([]map[string]int, 2)

	sm[0] = map[string]int{"a": 1}
	sm[1] = map[string]int{"b": 2}

	fmt.Println(sm)

	// --------------------------------------------------------------
	// 9) MAP OF SLICES
	// --------------------------------------------------------------
	ms := map[string][]int{
		"nums": {1, 2, 3},
	}

	fmt.Println(ms)

	// --------------------------------------------------------------
	// 10) PASS MAP TO FUNCTION (reference behavior)
	// --------------------------------------------------------------
	fmt.Println(m2)
	modifyMap(m2)
	fmt.Println(m2) // changed!

	// --------------------------------------------------------------
	// 11) NIL MAP VS EMPTY MAP
	// --------------------------------------------------------------
	var n1 map[string]int  // nil
	n2 := map[string]int{} // empty but not nil

	fmt.Println(n1 == nil) // true
	fmt.Println(n2 == nil) // false

	// n1["x"] = 10 // ❌ panic: assignment to entry in nil map

	// --------------------------------------------------------------
	// 12) CHECKING MAP EQUALITY
	// --------------------------------------------------------------
	// Direct map comparison is NOT allowed except with nil.

	// m1 == m2 // ❌ invalid

	// To compare, loop over keys or serialize JSON.

	// --------------------------------------------------------------
	// 13) CONCURRENT MAP (IMPORTANT!)
	// --------------------------------------------------------------
	// Regular maps are NOT safe for concurrent writes.
	// Use sync.Map or mutex if needed.

	// --------------------------------------------------------------
	// INTERVIEW CHECKLIST - MAPS
	// --------------------------------------------------------------
	// ✔ map declaration & initialization
	// ✔ make() vs literal initialization
	// ✔ insert, update, delete, lookup with ok
	// ✔ looping (order not guaranteed)
	// ✔ map of structs, slice of maps, map of slices
	// ✔ reference behavior in functions
	// ✔ nil map vs empty map
	// ✔ why map cannot be compared directly
	// ✔ concurrency-safe map concepts
}

// function showing reference behavior
func modifyMap(m map[string]int) {
	m["z"] = 999
}

// =====================================================================
// END OF GO MAPS MASTER NOTES
// =====================================================================
