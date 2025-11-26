// =====================================================================
//  GO VARIABLES, PRIMITIVE TYPES & CONSTANTS - MASTER NOTES (COMPLETE)
// =====================================================================
// Covers: var, :=, multiple declarations, scopes, primitive data types,
// zero-values, type conversion, constants, iota, typed vs untyped consts.
// =====================================================================

package main

import "fmt"

func main() {

	// --------------------------------------------------------------
	// 1) VARIABLE DECLARATION
	// --------------------------------------------------------------

	var a int = 10 // explicit type
	var b = 20     // type inferred
	c := 30        // short declaration (only inside functions)
	fmt.Println(a, b, c)

	// multiple variables
	var x, y, z = 1, 2, 3
	fmt.Println(x, y, z)

	// --------------------------------------------------------------
	// 2) ZERO VALUES (VERY IMPORTANT)
	// --------------------------------------------------------------
	var i int     // 0
	var s string  // ""
	var f float64 // 0.0
	var bl bool   // false

	fmt.Println(i, s, f, bl)

	// --------------------------------------------------------------
	// 3) PRIMITIVE DATA TYPES
	// --------------------------------------------------------------
	var n8 int8 = 127
	var un8 uint8 = 255
	var fl float32 = 3.14
	var complexNum complex64 = 1 + 2i

	fmt.Println(n8, un8, fl, complexNum)

	// string is immutable, supports UTF-8
	str := "GoLang"
	fmt.Println(str, len(str))

	// rune = alias for int32 for Unicode characters
	var r rune = '😀'
	fmt.Println(r, string(r))

	// --------------------------------------------------------------
	// 4) TYPE CONVERSION (explicit)
	// --------------------------------------------------------------
	var num int = 100
	var decimal float64 = float64(num)
	fmt.Println(decimal)

	// --------------------------------------------------------------
	// 5) CONSTANTS
	// --------------------------------------------------------------

	const PI = 3.1415        // untyped constant
	const maxUsers int = 100 // typed constant

	fmt.Println(PI, maxUsers)

	// Constants cannot be modified
	// PI = 3.14 //  compile error

	// --------------------------------------------------------------
	// 6) MULTIPLE CONSTANTS
	// --------------------------------------------------------------
	const (
		A = 1
		B = 2
		C = 3
	)
	fmt.Println(A, B, C)

	// --------------------------------------------------------------
	// 7) IOTA (Auto incrementing constant generator)
	// --------------------------------------------------------------

	const (
		r1 = iota // 0
		r2 = iota // 1
		r3 = iota // 2
	)
	fmt.Println(r1, r2, r3)

	const (
		_  = iota             // skip
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB = 1 << (10 * iota) // 1 << 20
		GB = 1 << (10 * iota) // 1 << 30
	)
	fmt.Println(KB, MB, GB)

	// --------------------------------------------------------------
	// 8) UNTYPED VS TYPED CONSTANTS
	// --------------------------------------------------------------

	const un = 5        // untyped (flexible in operations)
	var ff float64 = un // ok

	const tn int = 5 // typed
	// var ff2 float64 = tn // ❌ not allowed without explicit conversion

	fmt.Println(un, ff, tn)

	// --------------------------------------------------------------
	// 9) SHADOWING (IMPORTANT)
	// --------------------------------------------------------------

	value := 50
	{
		value := 99 // shadows outer variable
		fmt.Println("inner:", value)
	}
	fmt.Println("outer:", value)

	// --------------------------------------------------------------
	// INTERVIEW CHECKLIST
	// --------------------------------------------------------------
	// ✔ var, short declaration :=, multiple vars
	// ✔ zero-values (int→0, string→"", bool→false, pointers→nil)
	// ✔ primitive types: int, uint, float, complex, string, rune, byte
	// ✔ type conversion (explicit)
	// ✔ constants, typed vs untyped
	// ✔ iota usage (auto-increment, bit shifting)
	// ✔ important: shadowing rules
}

// =====================================================================
// END OF VARIABLES, PRIMITIVES & CONSTANTS MASTER NOTES
// =====================================================================
