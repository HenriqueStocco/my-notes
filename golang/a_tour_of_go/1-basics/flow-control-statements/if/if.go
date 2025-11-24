package main

import (
	"fmt"
	"math"
)

func shortStatement(x, n, lim float64) float64 {
	// Declara variavel local do if somente e verifica a condição
    // Isso também é early return, onde se a condição não for atingida
    // retorna algo, normalmente se valida algo ao contrario no if
    // e caso não atinja a condição, faz o que queremos
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		shortStatement(3, 2, 10),
		shortStatement(3, 2, 20),
	)
}
