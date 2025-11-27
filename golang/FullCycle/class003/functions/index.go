package functions

import "fmt"

func Functions() {
	value, message := retornoMultiplo(0, "Hello")
	fmt.Printf("Value: %v message: %v \n", value, message)

	name := retornoNomeado("Douglas")
	fmt.Printf("Name is: %v\n", name)

	somaTotal := somaTudo(2, 4, 6, 8)
	fmt.Printf("Total: %v\n", somaTotal)

	// Funcao dentro de uma variavel que retorna um funçao que retorna um int
	varFunction := func(x ...int) func() int {
		resultado := 0

		for _, value := range x {
			resultado += value
		}
		return func() int {
			return resultado * resultado
		}
	}
	fmt.Printf("Função em variavel: %v \n", varFunction(2, 2, 2, 2, 1)())
}

// retornoMultiplo: Função que retorna varias variaveis
func retornoMultiplo(value int, message string) (int, string) {
	return value, message
}

// retornoNomeado: Função que retorna explicitamente o que eu defini
func retornoNomeado(value string) (name string) {
	name = value
	return
}

// somaTudo: Rest params
func somaTudo(x ...int) int {
	resultado := 0

	for _, value := range x {
		resultado += value
	}

	return resultado
}
