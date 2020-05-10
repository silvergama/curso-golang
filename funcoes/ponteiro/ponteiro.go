package main

import "fmt"

func inc1(n int) {
	n++
}

// Revisão: Um ponteiro é representado por *
func inc2(n *int) {
	// Revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// Revsão: & é usado para obter o valor da variável
	inc2(&n)
	fmt.Println(n)

}
