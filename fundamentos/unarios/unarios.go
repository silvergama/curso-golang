package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou x = x - 1
	fmt.Println(y)

	// fmt.Println(x == y--) não pode usar decrement ou impremet dentro de uma comparação

}
