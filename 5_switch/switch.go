package main

import (
	"fmt"
	"time"
)

func main() {
	// SIMPLE SWITCH
	var num int = 3
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// MULTIPLE CONDITIONS SWITCH
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// TYPE SWITCH
	var i any = 1
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("other")
	}
}
