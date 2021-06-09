package main

import "fmt"

func main() {
	g0(0)
	fmt.Println("Program finished successfully!")
}

func g0(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic("Panic in g() (major)")
	}
	defer fmt.Println("Defer in g()", i)
	fmt.Println("Printing in g()", i)
	g0(i + 1)
}