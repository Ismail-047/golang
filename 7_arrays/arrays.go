package main

import "fmt"

/*
	NOTE: WE ALMOST NEVER USE ARRAYS DIRECTLY. IN REAL GO CODE, WE USE SLICES INSTEAD. ARRAYS ARE TOO RIGID - THEIR FIXED SIZE MAKES THEM IMPRACTICAL FOR MOST USE CASES.
*/

func main() {
	/*
		ARRAY - IF NOT INITIALIZED, THEN ZEROED VALUE OF THAT DATA TYPE IS ASSIGNED ON EACH INDEX BY DEFAULT
	*/

	// DECLARATION
	var arr [5]int // [0 0 0 0 0]

	var names [3]string // ["" "" ""]

	primes := [4]int{2, 3, 5, 7} // DECLARATION & INITIALIZATION

	auto := [...]int{1, 2, 3} // COMPILER COUNTS LENGTH

	// ARRAY LENGTH
	var length = len(arr)

	// 2D ARRAY
	arr4 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr4) // [[1 2] [3 4]]

	// NOTE: ARRAYS ARE VALUES, NOT REFERENCES. WHEN YOU ASSIGN OR PASS AN ARRAY, GO COPIES THE ENTIRE THING:
	a := [3]int{1, 2, 3}

	b := a // b IS A COPY
	b[0] = 99

	fmt.Println(a) // [1 2 3] - UNCHANGED
	fmt.Println(b) // [99 2 3] - CHANGED

	fmt.Println(arr, names, primes, auto, length)

}
