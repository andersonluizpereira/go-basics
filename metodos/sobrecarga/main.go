package main

import "fmt"

type triangle struct {
	size int
}

func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

type coloredTriangle struct {
	triangle
	color string
}

func main() {

	t2 := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", t2.size)
	fmt.Println("Perimeter", t2.perimeter())
}
