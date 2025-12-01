package main

import "fmt"

// Uma closure é uma funcao/metodo que retorna uma funcao
// a grande utilidade é que a funcao retornada possui o escopo da funcao que retorna tal funcao
// aplicação real desconhecida kkkkk
func i() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func main() {
	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
