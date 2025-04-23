package main

import "fmt"

func main() {
	fmt.Println("Switch case tutorial")

	// Default case is used when no other case matches
	day := 3

	// Alternative: You can declare the variable directly in the switch statement, e.g., switch day := 1; day { }
	switch day {
	case 1:
		fmt.Println("Monday")
		// Use fallthrough to execute the next case (e.g., Tuesday) even if the current case matches
	case 2:
		fmt.Println("Tuesday")
	case 3, 4: // Multiple values can be specified in a single case; this matches if day is 3 or 4
		fmt.Println("Case for 3 or 4")
	default:
		fmt.Println("Default case")
	}
}
