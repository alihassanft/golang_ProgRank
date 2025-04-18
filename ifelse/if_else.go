package main

import "fmt"

// we use for opertor or (||) and (&&)

func main() {
	fmt.Println("If Else Tutorial")
	i := 4

	if i >= 5 {
		fmt.Println("if >5", i)
	}

	// ifelse

	if i%2 == 0 {
		fmt.Println("even value", i)
	} else {
		fmt.Println("odd value", i)
	}

	// elif
	fmt.Println("2nd_loop ")
	if i >= 2 && i <= 4 {
		fmt.Println("2nd_loop 2-4", i)
	} else if i >= 5 && i <= 7 {
		fmt.Println("2nd_loop 5-7", i)
	} else {
		fmt.Println("2nd_loop else")
	}
}
