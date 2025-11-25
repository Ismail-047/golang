package main

import "fmt"

/*
	KEY TAKEAWAY: len(str) COUNTS BYTES, NOT CHARACTERS. ALWAYS USE RANGE LOOP FOR PROPER CHARACTER ITERATION.
*/

func StringRuneByte() {
	// STRING
	name := "HELLO"
	fmt.Println(name)      // HELLO
	fmt.Println(len(name)) // 5 (BYTES)

	emoji := "HI 🔥"
	fmt.Println(len(emoji)) // 7 (BYTES, NOT CHARACTERS!)

	// RUNE (SINGLE CHARACTER)
	var letter rune = 'A'
	fmt.Println(letter)        // 65 (ASCII CODE)
	fmt.Printf("%c\n", letter) // A (AS CHARACTER)

	fire := '🔥'
	fmt.Printf("%c\n", fire) // 🔥

	// BYTE
	var b byte = 'Z'
	fmt.Println(b)        // 90 (ASCII CODE)
	fmt.Printf("%c\n", b) // Z

	// ITERATING STRING BY RUNE VS BYTE
	text := "GO🔥"

	for i := 0; i < len(text); i++ {
		fmt.Printf("%c ", text[i]) // G O ï ¿ ½ (BROKEN - BYTE ITERATION)
	}
	fmt.Println()

	for _, r := range text {
		fmt.Printf("%c ", r) // G O 🔥 (CORRECT - RUNE ITERATION)
	}
}
