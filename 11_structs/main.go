package main

import "fmt"

// Structs can be used to model real life objects,
// almost similar to classes in object oriented languages
type person struct {
	// struct definitions
	name, DOB string
	age       int
}

// Struct methods
// 1. Value receivers (eg: (p person) someMethod(....) )
// 2. Pointer receivers (eg: (p *person) someMethod(....) )

// Value receiver example
func (p person) sayHi() string {
	return "Hi, " + p.name
}

// Pointer receiver example
func (p *person) increaseAge() {
	p.age++
}

func main() {
	// Initialize a struct
	biswas := person{
		name: "Biswas",
		age:  25,
		DOB:  "03021995",
	}

	// Accessing elements on a struct
	fmt.Println(biswas.name)
	fmt.Println(biswas.DOB)
	fmt.Println(biswas.age)

	// Type of struct variable
	fmt.Printf("%T\n", biswas) // main.person

	// Struct methods
	fmt.Println(biswas.sayHi())
	biswas.increaseAge()
	fmt.Println(biswas.age)
}
