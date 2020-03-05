package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // 1.2

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal) // 6

	// cuidado...
	fmt.Println("Teste " + string(97)) // 97 representa 'a' na tabela ASCII

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
