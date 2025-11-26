// ================================
// Example 1: Exported vs Unexported Functions
// ================================

package utils

// Add is EXPORTED because it starts with a capital letter.
// Accessible outside this package.
func Add(a, b int) int {
	return a + b
}

// subtract is UNEXPORTED because it starts with lowercase.
// Accessible only inside this package.
func subtract(a, b int) int {
	return a - b
}

// ================================
// Example 2: Struct Export Rules
// ================================

type Person struct { // Exported struct
	Name string // Exported field (Public)
	age  int    // Unexported field (Private)
}
