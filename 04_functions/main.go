package main

import "fmt"

// Basic function
func greeter(name string) string {
	return "Hello " + name
}

// Shortened input argument types, if they are same
func add(num1, num2 int) int {
	return num1 + num2
}

// Multiple return values
func swap(num1, num2 int) (int, int) {
	return num2, num1
}

func main() {
	fmt.Println("Greet: ", greeter("Biswas"))
	fmt.Println("Sum: ", add(10, 30))

	num1, num2 := 10, 20
	num1, num2 = swap(num1, num2)
	fmt.Println("Swap: ", num1, num2)
}
