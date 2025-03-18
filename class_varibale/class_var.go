package main

import "fmt"

// 30:56minit end

func main() {
	// two mtehod declar varibale

	// 1st method
	var a int //declaration
	a = 11    //initialization

	fmt.Println("a value:", a)

	// 2nd method
	var b int = 13 //declaration and initialization
	fmt.Println("b value: ", b)

	//3rd we can direct give value to variable withoiut define type
	var c = 21 //declaration and initialization
	println("c value", c)

	// 4th method without using keyword var and define data type
	d := 27
	println("d value", d)

	// var a int = 10
	// var b int
	// fmt.Println("variable a:", a)
	// fmt.Println("print zero", b)
}
