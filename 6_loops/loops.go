package main

import "fmt"

/*
	FOR - ONLY CONSTRUCT USE FOR LOOPING IN GOLANG
	NO WHILE OR DO-WHILE
*/

func main() {
	// FOR LOOP
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// WHILE LOOP
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// RANGE
	for index, value := range "helloworld" {
		fmt.Println(index, value)
	}

	// INFINITE LOOP
	for {
		fmt.Println("INFINITE LOOP")
	}
}
