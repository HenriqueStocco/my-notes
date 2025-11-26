package main

import "fmt"

// Um tipo de interface é definido como um conjunto de
// assinatura de metodos
// Um valor do tipo interface pode conter qualquer valor
// que implemente esses metodos
// As interfaces são implementadas implicitamente
// Um tipo implementa uma interface ao implementar seus metodos
type I interface {
	M()
}

type T struct {
	S string
}

//	Este método significa que o tipo T implementa a interface I,
//
// mas não precisamos declarar explicitamente que isso ocorre.
func (t T) M() {
	fmt.Println(t.S)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i)
}

// Nos bastidores, os valroes de inteface podem ser considerados como uma tupla
// de um valor e um tipo concreto: (valor, tipo)

func main() {
	var i I = T{"hello"}
	i.M()

	// O tipo de interface que especifica zero metodos é chamada de interface vazia
	// Um interface vazia pode conter valores de qualquer tipo
	// São utilizadas por códigos que lidam com valores de tipo desconhecido,
	// Um exemplo, o fmt.Print aceita qualquer número de arqgumentos do tipo interface {}
	var y interface{}
	describe(y)
}
