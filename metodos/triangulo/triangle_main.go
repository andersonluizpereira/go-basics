package main

import "fmt"

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (s square) perimeter() int {
	return s.size * 4
}

type coloredTriangle struct {
	triangle
	color string
}

func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())
	t1 := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", t1.size)
	fmt.Println("Perimeter", t1.perimeter())
}
