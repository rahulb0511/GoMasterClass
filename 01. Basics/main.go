package main

import "fmt"

func main() {

	// --------------------------------------------------------------
	// 1) IF / ELSE IF / ELSE
	// --------------------------------------------------------------

	age := 20

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}

	// if initialization syntax
	if score := 75; score >= 60 {
		fmt.Println("Pass")
	}

	// --------------------------------------------------------------
	// 2) SWITCH
	// --------------------------------------------------------------

	day := "tue"
	switch day {
	case "mon":
		fmt.Println("Monday")
	case "tue", "wed":
		fmt.Println("Mid week days")
	default:
		fmt.Println("Other day")
	}

	// switch without condition (acts like if/else chain)
	x := 10
	switch {
	case x < 0:
		fmt.Println("Negative")
	case x == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive")
	}

	// --------------------------------------------------------------
	// 3) FOR LOOPS (only loop keyword in Go)
	// --------------------------------------------------------------

	// Standard for loop
	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}

	// While-like loop
	j := 0
	for j < 3 {
		fmt.Println("j =", j)
		j++
	}

	// Infinite loop
	// for {
	//     fmt.Println("Loop forever")
	// }

	// --------------------------------------------------------------
	// 4) for-range loop
	// --------------------------------------------------------------

	arr := []string{"go", "lang"}
	for idx, val := range arr {
		fmt.Println(idx, val)
	}

	text := "Hi"
	for _, ch := range text { // iterates runes
		fmt.Println(string(ch))
	}

	// --------------------------------------------------------------
	// 5) BREAK and CONTINUE
	// --------------------------------------------------------------

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // skip index 3
		}
		if i == 5 {
			break // stop loop
		}
		fmt.Println(i)
	}

	// --------------------------------------------------------------
	// 6) LABEL + BREAK / CONTINUE (nested loops)
	// --------------------------------------------------------------

OuterLoop:
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			if a == 2 && b == 2 {
				break OuterLoop
			}
			fmt.Println("a, b =", a, b)
		}
	}

	// --------------------------------------------------------------
	// 7) GOTO (rare, but sometimes used)
	// --------------------------------------------------------------

	count := 1
Start:
	fmt.Println("count =", count)
	count++
	if count <= 3 {
		goto Start
	}
}

// =====================================================================
// END OF CONTROL FLOW MASTER NOTES
// =====================================================================
