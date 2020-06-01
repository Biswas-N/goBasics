package main

import "fmt"

func main() {
	// Closures - Inline functions like Lambda or arrow functions
	// in python and javascript
	greeter := func(name string) string {
		return "Hello " + name
	}

	fmt.Println(greeter("Biswas"))
}
