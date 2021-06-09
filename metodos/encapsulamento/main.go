package main

import (
	"fmt"
	"microsoft/metodos/encapsulamento/geometry"
)

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())
}
