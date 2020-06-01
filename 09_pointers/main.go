package main

import "fmt"

func main() {
	// & : address_of
	normalVar := 10
	pointerVar := &normalVar

	fmt.Printf("Typeof pointerVar: %T\n", pointerVar) // *int
	// %p prints Hexadecimal (base 16) value
	fmt.Printf("Value in pointerVar: %p\n", pointerVar) // 0xc0000120b0

	// * : value_at
	fmt.Println(*pointerVar) // 10
	*pointerVar += 10
	fmt.Println(*pointerVar) // 20
}
