package main

import "fmt"

// Uma das interfaces mais comuns é a Stringer, definida pelo pacote fmt.
// Um Stringer é um tipo que pode se descrever como uma string.
// O pacote fmt (e muitos outros) procura essa interface para imprimir valores.
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	fmt.Println(a, z)
}
