// Go Export Rules (Capital vs small letters)
// -----------------------------------------
// Go does NOT use public/private keywords. Instead:
// - Identifiers starting with CAPITAL letters are EXPORTED (public)
// - Identifiers starting with small letters are UNEXPORTED (package private)
// This applies to:
//   * Functions, Variables, Constants, Structs, Struct fields, Methods, Interfaces

// ================================
// Example 3: Usage in main package (Different file)
// ===============================
package main

import (
	"fmt"
	"package/utils" // import path depending on your module
)

func main() {
	fmt.Println(utils.Add(5, 3)) // ✔ Allowed (exported)
	//fmt.Println(utils.subtract(5, 3)) // ❌ Compile error (unexported)

	p1 := utils.Person{}
	p1.Name = "Rahul"
	//p1.age = 25    // ❌ compile error, unexported field
}

// ================================
// Summary (Easy to Remember)
// ================================
// Capital letter => Exported => Public
// small letter   => Unexported => Private
// Example: fmt.Println is exported; fmt.print is not.
