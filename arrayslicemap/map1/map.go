package main

import "fmt"

func main() {
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678911] = "Maria"
	aprovados[12345678912] = "Pedro"
	aprovados[12345678913] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678913])
	delete(aprovados, 12345678913)
	fmt.Println(aprovados[12345678913])

}
