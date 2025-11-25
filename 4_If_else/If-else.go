package main

import "fmt"

/*
	NOTE: GO DOES NOT HAVE TERNARY (?:) OPERATOR
	WE HAVE TO USE IF-ELSE
*/

func main() {

	age := 20

	// IF
	if age >= 18 {
		fmt.Println("PERSON IS AN ADULT")
	}

	// IF-ELSE
	if age >= 18 {
		fmt.Println("PERSON IS AN ADULT")
	} else {
		fmt.Println("PERSON IS NOT AN ADULT")
	}

	// IF-ELSE IF
	if age >= 18 {
		fmt.Println("PERSON IS AN ADULT")
	} else if age >= 13 {
		fmt.Println("PERSON IS A TEENAGER")
	} else {
		fmt.Println("PERSON IS NOT AN ADULT")
	}

	role := "admin"
	hasPermission := true

	if role == "admin" && hasPermission {
		fmt.Println("YOU ARE AN ADMIN AND HAVE PERMISSION")
	} else if role == "user" && hasPermission {
		fmt.Println("YOU ARE A USER AND HAVE PERMISSION")
	} else {
		fmt.Println("YOU ARE NOT AN ADMIN OR USER")
	}

	// SHORT IF (VARIABLE DECLARATION IN IF)
	if id := 1; id == 1 {
		fmt.Println("ID IS 1", id)
	} else {
		fmt.Println("ID IS NOT 1", id)
	}
}
