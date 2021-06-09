package main

import "fmt"

type Employee1 struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main1() {
	employee := Employee1{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee)
	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)
}
