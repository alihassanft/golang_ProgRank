package main

func main() {
	// 1st method thi be constent we cant change value
	const data = 25
	println("constant value", data)
	const (
		a = iota //idont undertand what is iota but its define a serilz and we dont need asign value to variable of consit
		b
		c
		// here in abouv three serils if i use _ ist skip that number
		_ //so here its skiping numver 3 becz its start fro 0
		d
	)
	println("a", a, "b", b, "c", c, "d", d)
}
