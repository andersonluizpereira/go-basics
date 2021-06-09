package main

import "microsoft/calculator"

func main() {
	total := calculator.Sum(3, 5)
	println(total)
	println("Version: ", calculator.Version)
}
