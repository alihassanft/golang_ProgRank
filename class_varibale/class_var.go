package main

import "fmt"

// 30:56minit end strat scope varaivble

func main() {
	// two mtehod declar varibale

	// test
	k := 9
	fmt.Println("K value", k)

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

	// scope varaible
	// 30:56minit end strat scope varaivble
	// three type of scope variable
	// global: in all over the program even out of function
	// local:  only withint functionn
	// package_level:

	// type and value using prinf

	z := "ali"

	fmt.Printf("%v, %T", z, z)

	// start primitive data types go 52:32
}
