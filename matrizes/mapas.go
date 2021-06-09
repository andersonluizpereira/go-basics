package main

import "fmt"

func main() {
	//inicializando a variavel
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)
	//sem inicializar use o make
	studentsAge1 := make(map[string]int)
	studentsAge1["john"] = 32
	studentsAge1["bob"] = 31
	fmt.Println(studentsAge1)
	//buscando por uma chave
	fmt.Println("Bob's age is", studentsAge["bob"])
	fmt.Println("Christy's age is", studentsAge["christy"])

	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}
	//removendo um elemento de uma das chaves
	delete(studentsAge, "john")
	fmt.Println(studentsAge)

	// varrendo o mapa
	studentsAge["john"] = 32
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}