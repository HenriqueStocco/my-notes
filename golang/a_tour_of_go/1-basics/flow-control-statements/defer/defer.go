package main

import "fmt"

// defer adia a execução de uma função, até que a função ao redor retorne
// Em meus testes, funções acima do defer, são executadas primeiro
// Abaixo do defer, são executadas depois das de cima e ai sim o defer
func simpleDefer() {
	fmt.Println("Before defer")

	defer fmt.Println("Defered")

	fmt.Println("After defer")
}

// Mas, se eu tiver defer encima, no meio e abaixo, e uma sem defer
// A sem defer executa, a debaixo executa, a do meio executa e depois a de cima
func inOrder() {
	fmt.Println("First Normal function")

	defer fmt.Println("Last execution")
	defer fmt.Println("Middle execution")
	defer fmt.Println("First execution before the normal function")

	fmt.Println("Second Normal function")
}

/*
   As chamadas de função diferidas são empurradas para uma pilha.
   Quando uma função retorna, suas chamadas diferidas são executadas na ordem último
   a entrar, primeiro a sair.

   counting executa
   done executa
   no for:
   9,8,7,6,... executam
*/
func stackingDefers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	simpleDefer()
	inOrder()
	stackingDefers()
}
