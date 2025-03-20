package main

import "fmt"

// deine Multi-dimensional Arrays in line 50

func main() {
	// 1st method
	// Declaration with size and type
	var numbers [4]int = [4]int{22, 33, 44, 55} //i notice that if we define 4 and we give three values its show thier 0,
	// if like we only give two data its give us 0,2,and otehr two data
	// if we extra values its show error

	fmt.Println("array numbers", numbers)
	fmt.Println("length of array", len(numbers)) //lenth of array
	fmt.Println("accsing index", numbers[0])     //access by index liek from 0,1----

	// 2nd method
	// Declaration and initialization
	numbers_2 := [3]int{17, 19, 20}
	fmt.Println("direct declare numvers2", numbers_2)
	// Modifies the  element using indexing
	numbers_2[2] = 99
	fmt.Println()
	fmt.Println("after modeification", numbers_2)

	fmt.Println()

	// 3rd nethod
	// if we not sure about size we  Using ellipsis (...)
	str_ell := [...]string{"Ali", "CS", "A+", "mobile"}
	numbers_ell := [...]int{23, 44, 55, 6}

	fmt.Println("String Array Ellipsis", str_ell)
	fmt.Println("String Ellipsis length of array", len(str_ell))
	fmt.Println()
	fmt.Println("Number Array Ellipsis", numbers_ell)
	fmt.Println("Number Ellipsis length of array", len(numbers_ell))

	fmt.Println("rnage")
	// suppose we deinne lenthy array so we need specfic data or kind of range of
	range_number := [...]int{2, 3, 4, 5, 6, 7}
	fmt.Println()
	fmt.Println("tetsing index", range_number[0])
	fmt.Println("range numbers", range_number[1:5]) //assume we need from 3-6 (so 3 have 1index ,6hvae 4 index(but we use one more index liek 5index))
	// if we [:3] if we intail emptye so its pick from start
	fmt.Println("intail index", range_number[4:]) //output [6 7]

	// if we intial have value and strt have no value then its pocik start point and to till ned index

	fmt.Println("end indexing", range_number[:2]) //ouutput [2 3]

	//
	println()
	println("Multi-dimensional Arrays")
	// Multi-dimensional Arrays
	// [2][3]int 2is row amd 3 is column
	var muti_dim [2][3]int = [2][3]int{{1, 2, 4}, {2, 4, 3}}
	var muti_dim_2 [3][2]int = [3][2]int{{1, 2}, {2, 4}}
	fmt.Println("Multi-dimensional Value", muti_dim)
	fmt.Println("Multi-dimensional Value 2", muti_dim_2)
	// accessing index so in array their dict {0index{},iindex{}} and in the dict data will be again 0index,2index and so on
	fmt.Println("index mutidimentional", muti_dim[1][2])

}
