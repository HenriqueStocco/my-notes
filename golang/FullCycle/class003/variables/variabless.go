package variables

import f "fmt"

var variavel string

func Variaveis() {
	// Atribuição da variavel
	variavel = "Henrique"
	// Declaração + Atribuição
	nome := 2025
	// Re-atribuição
	nome = 1

	f.Println(variavel, nome)

	// String interpolation
	f.Printf("%v \n", nome)
	f.Printf("%v \n", nome)
}
