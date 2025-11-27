package main

import "fmt"

func getBolas() int8 {
	return 2
}

func conditions() {
	bolas := 4

	if bolas > 1 {
		fmt.Println("Mais que 1 bola")
	} else if bolas > 4 {
		fmt.Println("Quantidade exagerada de bolas")
	} else {
		fmt.Println("Um pobre sem bolas")
	}

	if bola := getBolas(); bola < 2 {
		fmt.Println("HaHa, um mono bola")
	} else {
		fmt.Println("A, um muitas bolas")
	}
}
