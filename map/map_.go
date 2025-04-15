package main

import "fmt"

// map is collection of key pair value like dict

func main() {
	fmt.Println("map tutorial")

	// empSalary := make(map[string]int) //declartion
	// // intialization
	// empSalary = map[string]int{
	// 	"Ali":    122000,
	// 	"Hassan": 121000,
	// }
	// fmt.Println("employe salary", empSalary)

	// declartion and intailzation
	empSalary := map[string]int{
		"Ali":      9000,
		"Hassan":   1000,
		"shahazad": 1000,
	}

	// add new pair
	empSalary["Noman"] = 3000

	// delete
	delete(empSalary, "Ali")

	/* check availablity when we access a key from map its return value but
	we access a value whihc is not in map still its return zero
	fmt.Println("employe salary", empSalary["ahmed"]) its return zero even its not existin in our map
	so we check availblity its return with true and false as response
	if not availble its retun
	check availblity 0 false
	if availble its return 1000 is here ket is value
	check availblity 1000 true

	*/
	x, status := empSalary["Hassan"]
	fmt.Println("check availblity", x, status)

	fmt.Println("employe salary", empSalary)
	// check length
	fmt.Println("length map: ", len(empSalary))
}
