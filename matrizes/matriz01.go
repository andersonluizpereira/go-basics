package main

import "fmt"

func main() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("cities", cities)
	fmt.Println("cities", cities[2])
	//Outra maneira de declarar e inicializar uma matriz quando você não sabe de quantas posições precisará,
	//mas sabe quantos dados terá, é usar reticências (...), como neste exemplo:
	cities1 := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities1:", cities1[3])

	//Outra maneira interessante de inicializar uma matriz é usar reticências e só especificar um valor para a
	//última posição. Por exemplo, use o seguinte código:
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))

	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)

	var threeD [3][5][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("\nAll at once:", threeD)
}