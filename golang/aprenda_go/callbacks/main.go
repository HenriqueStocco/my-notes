package main

import "fmt"

func soma(x ...int) int {
	n := 0

	for v := range x {
		n += v
	}

	return n
}

// Um callback é quando uma funcao/metodo recebe uma funcao como parametro e após a execução da funcao que recebe uma funcao,
// ela chama a funcao recebida
func somentePares(f func(x ...int) int, y ...int) int {
	var slice []int

	for v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}

	total := f(slice...)

	return total
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int

	for v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}

	total := f(slice...)

	return total
}

func main() {
	p := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	i := somenteImpares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)

	fmt.Println(p)
	fmt.Println(i)
}
