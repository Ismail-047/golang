package main

import "fmt"

func main() {
	
	var name string = "John" // EXPLICITLY MENTION TYPE
	var age = 25             // TYPE AUTO INFERRED
	count := 10              // SHORT DECLARATION (INSIDE FUNCTIONS ONLY)

	// MULTIPLE VARIABLES:
	var x, y int = 1, 2
	a, b := "hello", 42

	// CONSTANTS
	const pi = 3.14
	const ( // MULTIPLE CONSTANTS
		port = 3000
		host = "localhost"
	)

	fmt.Println(name, age, count, x, y, a, b);
}

// NOTE: UNUSED VARIABLES = COMPILE ERROR (GO DOESN'T TOLERATE DEAD CODE)