package main

import "fmt"

// A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

// To declare a structure(struct) in Go, use the type and struct keywords:
type Person struct {
	name string
	age  int

	// phonenumbers []int  we deifne with [] that means its can hold mutiple numbers
}

func main() {
	fmt.Println("struct class")
	var pers1 Person
	var pers2 Person
	// person1
	pers1.name = "ali"
	pers1.age = 27

	// person2
	pers2.name = "Hassan"
	pers2.age = 22

	// person1
	// complete print
	fmt.Println("perosn1", pers1)
	// singel with key
	fmt.Println("person1 name", pers1.name)
	fmt.Println("person1 age", pers1.age)

	// person2
	// complete print
	fmt.Println("perosn2", pers2)
	// singel with key
	fmt.Println("person2 name", pers2.name)
	fmt.Println("person2 age", pers2.age)

}
