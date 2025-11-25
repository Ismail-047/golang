package main

import "fmt"

func main() {
	/*
		ARRAY -
		IF NOT ASSIGNED THEN ZEROED VALUE OF THAT DATA TYPE IS ASSIGNED ON EACH INDEX
	*/
	var arr [3]int
	fmt.Println(arr)
	fmt.Println(len(arr)) // ARRAY LENGTH

	// ARRAY WITH INITIAL VALUES
	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr2)

	// ARRAY WITH INITIAL VALUES
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 2D ARRAY
	arr4 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr4) // [[1 2] [3 4]]

}
