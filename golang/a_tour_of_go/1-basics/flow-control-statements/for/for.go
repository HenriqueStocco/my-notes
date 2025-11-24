package main

import "fmt"

/*
 Go só tem apenas um construtor de loops, é o for

    Componentes do for
        declaracao inicial - executado antes da primeira iteração
        expressao - verificado antes de cada iteração
        proxima declaração - executada no final de cada iteração
*/
func commonFor() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += 1
	}

	fmt.Println(sum)
}

/*
   A declaração inicial e posterior são opcionais
*/
func forContinued() {
	// Omitindo a declaração posterior com ;
	for sum := 0; sum < 10; {
		sum += 1
		fmt.Println(sum)
	}

	// Omitindo a declaração inicial com ;
	i := 0
	for ; i < 2; i++ {
		println(i)
	}
}

// Neste ponto, podemos ignorar o ;, em Go, while é chamado de for
func while() {
	sum := 90

	for sum < 100 {
		sum += 1
	}

	fmt.Println(sum)
}

func infiniteLoop() {
	// Loop infinito é só não passar nada no for
	for {
	}
}

func main() {
	commonFor()
	forContinued()
	while()
	// PERIGOSO
	// infiniteLoop()
}
