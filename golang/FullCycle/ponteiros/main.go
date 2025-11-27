package main

import "fmt"

/*
Ponteiros

Pega o endereço de memória:
ex -> &variavel

Cria um ponteiro:
ex -> var variavel *tipo = valor

Desreferencia um ponteiro para buscar o valor:
ex -> *variavel
*/

// Crio um objeto (quase uma classe)
type Car struct {
	Name string
}

// Crio um metodo para o objeto(classse)
// E falo que a propriedade Name vai ter o valor de BMW
// Este é o objeto original
// Caso o metodo não receba a struct como ponteiro
// O valor variavel Name na copia não é alterado após a chamada do metodo
// Isso se dá por que ambas as variaveis estão em endereços de memoria
// diferentes, mas caso seja um ponteiro, o valor é alterado.
func (c *Car) run() {
	c.Name = "BMW"
	fmt.Printf("Andou: %v\n", c.Name)
}

func main() {
	// Crio uma instancia do objeto(ou copia)
	// E atribuo o valor de Ka para a propriedade Name
	car := Car{
		Name: "Ka",
	}
	fmt.Printf("Name before: %v\n", car.Name)
	car.run()
	fmt.Printf("Name after: %v\n", car.Name)
}

// func baseDePonteiros() {
// 	a := 10
// 	var pointer *int = &a
// 	// Atribuir um novo valor ao ponteiro
// 	*pointer = 50
// 	b := &a
// 	*b = 60

// 	fmt.Printf("Variable before: %v\n", a)
// 	fmt.Printf("Pointer: %v \n", *pointer)
// 	fmt.Printf("Variable after: %v\n", a)
// 	fmt.Printf("B: %v\n", *b)
// }

// func abc(a *int) int {
// 	*a = 200

// 	return *a
// }
