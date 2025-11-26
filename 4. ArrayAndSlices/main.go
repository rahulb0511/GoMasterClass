// =====================================================================
//  GO ARRAYS & SLICES - MASTER INTERVIEW NOTES
// =====================================================================
// These notes cover 100% of Array & Slice concepts for Go developer interviews.
// Use this file for revision before interviews. All examples are runnable.
// =====================================================================

package main

import "fmt"

func main() {
	// --------------------------------------------------------------
	// 1) ARRAY BASICS
	// --------------------------------------------------------------
	// Arrays are fixed-size and stored in continuous memory.

	var arr [3]int   // all elements default to 0
	fmt.Println(arr) // [0 0 0]

	arr2 := [3]int{1, 2, 3} // declare + initialize
	fmt.Println(arr2)       // [1 2 3]

	arr3 := [...]int{10, 20, 30, 40} // auto-size array
	fmt.Println(arr3, len(arr3))     // [10 20 30 40] 4

	// --------------------------------------------------------------
	// 2) ARRAY COPY BEHAVIOR (VERY IMPORTANT: VALUE COPY)
	// --------------------------------------------------------------
	// Arrays are copied by value, not by reference!
	// Meaning: modifying one does NOT change the other.

	a := [3]int{1, 2, 3}
	b := a // full copy created
	b[0] = 99
	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [99 2 3]

	// --------------------------------------------------------------
	// 3) ARRAYS & LOOPS
	// --------------------------------------------------------------
	for i := 0; i < len(arr2); i++ {
		fmt.Println("Index", i, "=", arr2[i])
	}

	for i, v := range arr3 {
		fmt.Println("range -> i, v:", i, v)
	}

	// --------------------------------------------------------------
	// 4) MULTI-DIMENSIONAL ARRAY
	// --------------------------------------------------------------
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix) // [[1 2 3] [4 5 6]]

	// ==============================================================
	// 5) SLICE BASICS (MOST IMPORTANT PART FOR GO INTERVIEWS)
	// ==============================================================

	// Slice is dynamic, flexible, backed by an underlying array.
	s := []int{10, 20, 30}
	fmt.Println(s) // [10 20 30]

	// append grows slice capacity automatically
	s = append(s, 40)
	fmt.Println(s) // [10 20 30 40]

	// --------------------------------------------------------------
	// 6) MAKE() for slices
	// --------------------------------------------------------------
	s2 := make([]int, 2, 5)           // length 2, capacity 5
	fmt.Println(s2, len(s2), cap(s2)) // [0 0] 2 5

	// --------------------------------------------------------------
	// 7) SLICING ARRAYS
	// --------------------------------------------------------------
	base := [5]int{10, 20, 30, 40, 50}
	s3 := base[1:4] // slice from array
	fmt.Println(s3) // [20 30 40]

	// VERY IMPORTANT: slices share underlying array
	s3[0] = 999
	fmt.Println(base) // [10 999 30 40 50]

	// --------------------------------------------------------------
	// 8) COPY SLICE SAFELY (SO CHANGES DON'T AFFECT ORIGINAL)
	// --------------------------------------------------------------
	original := []int{1, 2, 3}
	clone := make([]int, len(original))
	copy(clone, original)

	clone[0] = 99
	fmt.Println(original) // [1 2 3]
	fmt.Println(clone)    // [99 2 3]

	// --------------------------------------------------------------
	// 9) CAPACITY GROWTH WITH APPEND
	// --------------------------------------------------------------
	grow := []int{}
	for i := 1; i <= 8; i++ {
		grow = append(grow, i)
		fmt.Println("len=", len(grow), "cap=", cap(grow), grow)
	}

	// --------------------------------------------------------------
	// 10) NIL SLICE vs EMPTY SLICE
	// --------------------------------------------------------------
	var x []int           // nil slice
	y := []int{}          // empty slice, not nil
	fmt.Println(x == nil) // true
	fmt.Println(y == nil) // false

	// --------------------------------------------------------------
	// 11) ARRAY IN FUNCTION (VALUE SEMANTICS)
	// --------------------------------------------------------------
	arrFunc := [3]int{1, 2, 3}
	modify(arrFunc)
	fmt.Println(arrFunc) // [1 2 3] unchanged!

	// --------------------------------------------------------------
	// 11) SLICE IN FUNCTION (REFERENCE SEMANTICS)
	// --------------------------------------------------------------
	arrFunc2 := []int{1, 2, 3}
	modify2(arrFunc2)
	fmt.Println(arrFunc2) // [100 2 3] changed!

	// --------------------------------------------------------------
	// 12) APPEND TRAP (WHEN SLICE RE-ALLOCATES)
	// --------------------------------------------------------------
	t := []int{1, 2, 3}
	fmt.Println(len(t), cap(t))
	t2 := t
	t = append(t, 4) // may allocate new array
	t[0] = 999
	fmt.Println(t)  // [999 2 3 4]
	fmt.Println(t2) // [1 2 3] (unchanged after reallocation!)

	// --------------------------------------------------------------
	// 13) CONVERT ARRAY TO SLICE & SLICE TO ARRAY
	// --------------------------------------------------------------
	arrConv := [3]int{5, 6, 7}
	sliceConv := arrConv[:] // array -> slice
	fmt.Println(sliceConv)

	// slice to array (only same size allowed)
	arrFromSlice := [3]int(sliceConv)
	fmt.Println(arrFromSlice)

	// --------------------------------------------------------------
	// 14) TWO-DIMENSIONAL SLICES
	// --------------------------------------------------------------
	grid := [][]int{
		{1, 2},
		{3, 4, 5},
	}
	fmt.Println(grid)

	// --------------------------------------------------------------
	// 15) SORTING SLICES
	// --------------------------------------------------------------
	nums := []int{5, 1, 4, 3, 2}
	sortInts(nums)
	fmt.Println(nums) // [1 2 3 4 5]

	// --------------------------------------------------------------
	// 16) REMOVE ELEMENT FROM SLICE
	// --------------------------------------------------------------
	rem := []int{10, 20, 30, 40, 50}
	index := 2
	rem = append(rem[:index], rem[index+1:]...)
	fmt.Println(rem) // [10 20 40 50]

	// --------------------------------------------------------------
	// 17) INSERT ELEMENT IN SLICE
	// --------------------------------------------------------------
	ins := []int{1, 2, 4, 5}
	idx := 2
	ins = append(ins[:idx], append([]int{3}, ins[idx:]...)...)
	fmt.Println(ins) // [1 2 3 4 5]

	// --------------------------------------------------------------
	// INTERVIEW CHECKLIST - ARRAYS & SLICES
	// --------------------------------------------------------------
	// ✔ array vs slice difference
	// ✔ value copy vs reference
	// ✔ underlying array sharing
	// ✔ append reallocation trap
	// ✔ make(), len(), cap(), copy()
	// ✔ nil slice vs empty slice
	// ✔ slicing behavior
	// ✔ deletion & insertion patterns
	// ✔ multi-dimensional slices
	// ✔ capacity growth behavior
}

// helper function that modifies array
func modify(s [3]int) {
	s[0] = 100
}

// helper function that modifies slice
func modify2(s []int) {
	s[0] = 100
}

// custom sort (simple bubble sort for demonstration)
func sortInts(s []int) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
