package main

import "fmt"

func main() {
	// defer schedules a statement to run when the surrounding function exits.
	// It is often used for cleanup tasks. Deferred statements execute in LIFO order.
	fmt.Println("Defer tutorial")        // Prints first
	defer fmt.Println("End of function") // Deferred: prints last when main exits

	fmt.Println("Start") // Prints second
	fmt.Println("End")   // Prints third

	// Call another function to demonstrate defer in a different context
	exampleMain()
}

func exampleMain() {
	fmt.Println("Start of exampleMain function")

	// Defer a print statement; it runs when exampleMain exits
	defer fmt.Println("This is deferred: executed last")

	// Defer another print statement; it runs before the previous defer (LIFO order)
	defer fmt.Println("This is deferred: executed second-to-last")

	fmt.Println("Doing some work in exampleMain")

	// Call another function to show nested defer behavior
	exampleFunction()

	fmt.Println("End of exampleMain function")
}

func exampleFunction() {
	fmt.Println("Inside exampleFunction")

	// Defer a statement; it runs when exampleFunction exits
	defer fmt.Println("Deferred in exampleFunction")
}

// Print statements
// Defer tutorial
// Start
// End
// Start of exampleMain function
// Doing some work in exampleMain
// Inside exampleFunction
// Deferred in exampleFunction
// End of exampleMain function
// This is deferred: executed second-to-last
// This is deferred: executed last
// End of function
