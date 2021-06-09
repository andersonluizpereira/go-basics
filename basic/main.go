package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	//Com variaveis inicializaveis
	// firstName   string  = "John"
	// lastName    string  = "Doe"
	// age         int     = 32
	// integer8    int8    = 127
	// integer16   int16   = 32767
	// integer32   int32   = 2147483647
	// integer64   int64   = 9223372036854775807
	// float321    float32 = 2147483647
	// float641    float64 = 9223372036854775807
	// rune                = 'G'
	// featureFlag bool    = true

	// Valores iniciais
	firstName   string
	lastName    string
	age         int
	integer8    int8
	integer16   int16
	integer32   int32
	integer64   int64
	float321    float32
	float641    float64
	rune        = 'G'
	featureFlag bool
)

const (
	StatusOK              = 0
	StatusConnectionReset = 1
	StatusOtherError      = 2
)

// \n para novas linhas
// \r para retornos de carro
// \t para tabulações
// \' para aspas simples
// \" para aspas duplas
// \\ para barras invertidas

func main() {
	fmt.Println("hello world")
	fmt.Println(firstName, lastName, age)
	fmt.Println(integer8, integer16, integer32, integer64)
	fmt.Println(rune)
	println(float321, float641)
	println(math.MaxFloat32, math.MaxFloat64)
	fmt.Println("featureFlag ", featureFlag)
	fullName := "John Doe \t(alias \"Foo\")\n"
	println(fullName)
	println(int32(integer16) + integer32)
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	println(i, s)
}
