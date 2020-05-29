package main

import "fmt"

func main() {
	// Maps - These are key value pairs, just like dictionaries in python
	// Define a map
	emails := make(map[string]string)

	// Assignment
	emails["Biswas"] = "biswas@gmail.com"
	emails["Ritu"] = "ritu@gmail.com"
	emails["Pallavi"] = "pallavi@gmail.com"

	fmt.Println(emails)

	// Initialization
	curr := map[string]string{"AUD": "Australia", "INR": "Indian"}
	fmt.Println(curr)

	// Accessing an item
	fmt.Println(curr["AUD"])

	// Deleting an item
	delete(emails, "Biswas")
	fmt.Println(emails)
}
