package main

import "fmt"

func formatStrings() {
	// Imprime uma string formatada na saida padrão,
	fmt.Printf("")

    // Retorna a string formatada
    fmt.Sprintf("")

    // Interpolação a representação padrão
    // %v - Valor qualquer, caso não saiba o tipo especifico
    fmt.Sprintf("Hello, %v!", "Opa")

    // %d - Valor inteiro
    fmt.Sprintf("I'm %d years old", 22)

    // %f - Valor com casas decimais
    fmt.Sprintf("PI %f", 3.1415)
}