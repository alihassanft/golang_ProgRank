package main

import "fmt"

func main() {
	// Declare a regular integer variable
	num := 42
	fmt.Println("Original value of num:", num) // Prints: 42

	// Declare a pointer to num using & (address-of operator)
	ptr := &num
	fmt.Println("Pointer address:", ptr)   // Prints the memory address
	fmt.Println("Value at pointer:", *ptr) // Prints: 42 (dereferencing)

	// Modify the value at the pointer's address
	*ptr = 100
	fmt.Println("New value of num:", num) // Prints: 100 (num changed via pointer)

	// Pass pointer to a function to modify the original value
	updateValue(&num)
	fmt.Println("After function call, num:", num) // Prints: 200
}

// Function that takes a pointer and modifies the value at that address
func updateValue(n *int) {
	*n = 200 // Dereference and set the value to 200
}
