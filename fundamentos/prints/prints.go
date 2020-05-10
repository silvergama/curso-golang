package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é " + x)
	xs := fmt.Sprint(x) // Ao invés de imprimir no console essa funcão retorna o valor transformado em string
	fmt.Println("O valor de x é " + xs)

	fmt.Println("O valor de x é", x) // mesmo resultado sem a necessidade de criar uma variável

	fmt.Printf("O valor de x é %.2f", x) // com Printf, temos mais controle. .2f indica que queremos apenas duas casas decimais. .2f também faz o arredondamento do valor

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
