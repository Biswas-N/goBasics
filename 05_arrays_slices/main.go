package main

import "fmt"

func main() {
	// Arrays - Should be fixed length
	// Declaring array
	var fruitArr [2]string

	// Assigning the values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Initializing an array (Declare + Assing)
	fruitArr2 := [4]string{"Apple", "Orange", "Grapes", "Jackfruit"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr2)

	// Slices - Can be of variable length and be derived from arrays
	// Initializing slice from array
	fruitSlice := fruitArr2[1:3]
	// To append anything to an existing slice, we use append function
	fruitSlice = append(fruitSlice, "Custed Apple", "Mango", "Goa")

	fmt.Println(fruitSlice)
}
