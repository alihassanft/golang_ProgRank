package main

import "fmt"

func main() {
	fmt.Println("slicing vs array")

	// this sample of array where we use [...] or define  lenth [4]int
	var numbers [4]int = [4]int{22, 33, 44, 55} //i notice that if we define 4 and we give three values its show thier 0,
	fmt.Println("array numbers", numbers)

	// their deiifenrce only we dont define range of array so its differet from slicing
	// slicing
	var number []int = []int{1, 2, 3, 4}
	fmt.Println("slicing numbers", number)

}
