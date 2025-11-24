package main

import "fmt"

func arrays() {
	// Um array é uma lista de itens de mesmo tipo e com o tamanho definido
	// o tamanho do array é parte do tipo dele, não pode ser redefinido
	// Não sei se segue a logica de literais mas o array abaixo não é literal
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"

	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	// Se seguir a logica, segue um array literal
	primes := [4]int{2, 3, 5, 7}
	fmt.Println(primes)
}

func main() {
	arrays()
}
