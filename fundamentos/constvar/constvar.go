package main

import (
	"fmt"
	m "math" // m alias
)

func main() {
	const PI float64 = 3.1415 // Por padrão um literal é float64
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma variável
	area := PI * m.Pow(raio, 2) // Sempre que criar uma variável, é necessário usá-la
	fmt.Println("A área da circunferência é", area)

	// Blocos de construção de constantes
	const (
		a = 1
		b = 2
	)

	// Blocos de construção de variáveis
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// Declaração de várias variáveis em uma única linha
	var e, f bool = true, false
	fmt.Println(e, f)
 
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)

}