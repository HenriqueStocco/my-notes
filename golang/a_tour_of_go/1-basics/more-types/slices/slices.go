package main

import (
	"fmt"
	"strings"
)

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// Slices são listas com o tamanho flexivel mas com itens de mesmo tipo
	// definidos pelo range [low : high]
	var slice []int = primes[1:4]
	fmt.Println(slice)
}

func slicesLikeRefToArray() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	// Slices são como referências do array
	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	// Se você altera um valor de um indice, o array mostrará essas alterações
	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(names)
}

func literals() {
	// Array literal -> [3]bool{true, true, false}
	// Slice literal -> []bool{true, true, false}
	q := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}

	fmt.Println(r)

	// Struct de slice
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}

func padraoSlice() {
	// Podemos omitir os limites do slice, high or low
	// para usar os valores zeros dele
	s := []int{2, 3, 5, 7, 11, 13}

	// Com os limites
	s = s[1:4]
	fmt.Println(s)

	// Sem o limite superior
	s = s[:2]
	fmt.Println(s)

	// Sem o limite inferior
	s = s[1:]
	fmt.Println(s)
}

func sliceLengthAndCapacity() {
	printSlice := func(s []int) { fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) }

	// Slices possuem o comprimento e a capacidade
	// O comprimento é o numero de elementos que ele contém
	// A capacidade é o numero de elementos no array subjacente, contando a partir do primeiro elemento
	// O tamanho pode ser obtido pela expressão: len(slice)
	// A capacidade pode ser obtida pela expressão: cap(slice)
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// fatie a fatia para dar comprimento 0
	// o comprimento é 0 por que o range é do inicio até 0 elementos
	// a capacidade é 6 por que é o numero de elementos contidos no slice
	s = s[:0]
	printSlice(s)

	// extendendo seu comprimento
	// o comprimento agora é 4 por que foi definido do inicio até o 4 elemento
	s = s[:4]
	printSlice(s)

	// Dropa os primeiros 2 valores
	// o comprimento é 2 por que o inicio foi definido 2
	// a capacidade é 4 porque os 2 primeiros valores foram dropados
	// por que começa do indice 2 e vai até o final
	s = s[2:]
	printSlice(s)
}

func nilSlices() {
	// O valor zero de um slice é nil
	// Um nil slice possue sua capacidade e comprimento iguais a 0
	var s []int

	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}

func sliceWithMake() {
	printSlice := func(s string, x []int) { fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x) }

	// Criando um array com um tamanho dinâmico
	// A função interna make, aloca um array zerado e retorna um slice que se refere a este array
	a := make([]int, 5) // len(a)=5
	printSlice("a", a)

	// Para especificar a capacidade, passe o terceiro argumento
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func sliceOfSlice() {
	// Exemplo com um Jogo da Velha
	/* Exemplo do Tour of Go
	    board := [][]string{
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
		}
	*/
	// O LSP do Go recomenda tirar os tipos redundantes
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// Exemplo do Tour of Go
	/*
		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}
	*/

	// O LSP do GO recomenda fazer o range direto
	for i := range board {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendingToASlice() {
	printSlice := func(s []int) { fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) }

	var s []int
	printSlice(s)

	// append funciona em nil slices
	s = append(s, 0)
	printSlice(s)

	// A slice cresce conforme necessário
	s = append(s, 1)
	printSlice(s)

	// Nós podemos adicionar mais de um elemento de uma vez
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeOverSlice() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// Quando usamos o range sobre um slice, 2 valores são retornados para cada iteração
	// O primeiro é o index e o segundo é uma cópia do elemento neste indice.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func rangeOverSliceContinued() {
	pow := make([]int, 10)

	// Omitindo o valor, deixando apenas o i
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// Omitindo o indice e pegando apenas o valor
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	slices()
	slicesLikeRefToArray()
	literals()
	padraoSlice()
	sliceLengthAndCapacity()
	nilSlices()
	sliceWithMake()
	sliceOfSlice()
	appendingToASlice()
	rangeOverSlice()
	rangeOverSliceContinued()
}
