package main

import "fmt"

func main() {
	x, y, z := 5, 10, 15

	// if-else
	if x < y {
		fmt.Println("The main if block")
	} else if x < z {
		fmt.Println("The else if block")
	} else {
		fmt.Println("The else block")
	}

	// Using switch
	switch x {
	case 10:
		fmt.Println("X is 10")
	case 20:
		fmt.Println("X is 20")
	default:
		fmt.Println("X is NOT 10 or 20")
	}
}
