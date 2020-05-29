package main

import "fmt"

// Global Variables and Constants
var (
	name  = "Biswas Nandamuri"
	age   = 25
	place = "Melbourne"
)

const (
	nameConst  = "Biswas Nandamuri"
	ageConst   = 25
	placeConst = "Melbourne"
)

func main() {
	// There are two types to declare variables
	// 1. Specific / Declarative
	var firstName string = "Biswas"

	// 2. Auto Inferring (2 types)
	lastName := "Nandamuri"
	var age = 25

	// Constants
	const nature = "Cool"

	// Declaring multiple variables at once
	// Emails
	gmail, outlook := "biswas@gmail.com", "biswas@outlook.com"

	fmt.Printf("%s %s is %d years old and he is %s\n", firstName, lastName, age, nature)
	fmt.Printf("%s's emails are %s, %s\n", firstName, gmail, outlook)
	fmt.Printf("\nTypes:\nfirstName: %T\nlastName: %T\nage: %T\n", firstName, lastName, age)
}
