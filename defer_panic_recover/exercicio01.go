//Para o caso de FizzBuzz, multiplica-se 3 por 5 porque os resultados são divisíveis por 3 e 5.
//Também é possível incluir uma condição AND para verificar se um número é divisível por 3 e 5.
package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(num)
}

func main() {
	for num := 1; num <= 100; num++ {
		fmt.Println(fizzbuzz(num))
	}
}