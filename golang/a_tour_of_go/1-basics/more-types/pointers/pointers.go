package main

import "fmt"

func pointers() {
	// Ponteiros armazenam endereços de memória de um valor
	// Todo ponteiro não inicializado, tem o seu valor como nulo, famoso nil
	var p *int
	fmt.Println(p) // p = <nil>

	// O operador &, pega o endereço de memória de uma variavel
	i := 42
	p = &i
	fmt.Println(p) // mostra o endereço

	// Desreferenciar ou indirecionar é pegar o valor armazenado
	// no endereço de memória
	z := *p
	fmt.Println(z) // obtendo o valor 42
}

func main() {
	pointers()
}
