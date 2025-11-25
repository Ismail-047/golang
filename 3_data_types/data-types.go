package main

import "fmt"

func main() {

	/* ----- NUMBERS ----- */

	// SIGNED INTEGERS - INCLUDES POSITIVE AND NEGATIVE NUMBERS
	var num int    // DEFAULT DEPENDS ON THE SYSTEM 32 OR 64
	var num1 int8  // -128 TO 127 - 1 BYTE
	var num2 int16 // -32768 TO 32767 - 2 BYTES
	var num3 int32 // -2 BILLION TO 2 BILLION - 4 BYTES
	var num4 int64 // -922 BILLION TO 922 BILLION - 8 BYTES

	// UNSIGNED INTEGERS - INCLUDES ONLY POSITIVE NUMBERS
	var num5 uint   // DEFAULT DEPENDS ON THE SYSTEM 32 OR 64
	var num6 uint8  // 0 TO 255 - 1 BYTE
	var num7 uint16 // 0 TO 65535 - 2 BYTES
	var num8 uint32 // 0 TO 4 BILLION - 4 BYTES
	var num9 uint64 // 0 TO 18 BILLION - 8 BYTES

	// FLOATING POINT NUMBERS
	var num10 float32 // ~ 7 DIGITS PRECISION - 4 BYTES
	var num11 float64 // ~ 15 DIGITS PRECISION (DEFAULT) - 8 BYTES

	// COMPLEX NUMBERS
	var com complex64    // REAL + IMAGINARY PARTS ARE FLOAT32
	var com1 complex128  // REAL + IMAGINARY PARTS ARE FLOAT64
	c := 3 + 4i          // COMPLEX NUMBER
	fmt.Println(real(c)) // 3
	fmt.Println(imag(c)) // 4

	/* ----- BOOLEAN ----- */
	var boolean bool // TRUE / FALSE

	/* ----- TEXT ----- */
	var str string // IMMUTABLE UTF-8
	var char rune  // SINGLE CHARACTER (ALIAS FOR INT32)
	var bytee byte // ALIAS FOR UINT8
	// SEE DETAILS IN string-rune-byte.go

	/* ----- COMPOSITE TYPES ----- */
	var int_arr []int            // SLICE (DYNAMIC)
	var int_arr_fixed [5]int     // ARRAY (FIXED SIZE)
	var key_value map[string]int // KEY-VALUE PAIRS
	var structure struct{}       // CUSTOM TYPES

	/* ----- SPECIAL ----- */
	var pointer *int      // INTEGER POINTER
	var channel chan int  // CHANNELS FOR CONCURRENCY
	var inter interface{} // ANY TYPE
	/*
		interface{} // old way (Go 1.17 and earlier)
		any         // new alias (Go 1.18+)
	*/

	fmt.Println(num, num1, num2, num3, num4, num5, num6, num7, num8, num9, num10, num11, com, com1, str, char, bytee, boolean, int_arr, int_arr_fixed, key_value, structure, pointer, inter, channel)
}

// WHY SO MANY TYPES?
type Pixel struct {
	R, G, B uint8 // RGB VALUES ARE 0-255, WHY WASTE 7 BYTES PER PIXEL?
}
