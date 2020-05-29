package main

import "fmt"

func main() {
	// Long method - like While in C
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// Short method
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	// Infinite loop - should have a condition to break
	breaker := 6
	k := 1
	for {
		if k == breaker {
			break
		}

		fmt.Println(k)
		k++
	}

	// Loops on Arrays or Slices
	sampleArr := [5]string{"One", "Two", "Three", "Four", "Five"}
	for position, value := range sampleArr {
		fmt.Printf("%d -> %s\n", position, value)
	}

	// Loops on Maps
	sampleMap := map[string]string{"AUD": "Australia", "INR": "Indian"}
	for key, value := range sampleMap {
		fmt.Printf("%s -> %s\n", key, value)
	}
}
