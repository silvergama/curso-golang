package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		fatorialAnterior := fatorial(n - 1)
		return n * fatorialAnterior
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

	// Gera um erro: constant -4 overflows uint
	resultado2 := fatorial(-4)
	fmt.Println(resultado2)
}
