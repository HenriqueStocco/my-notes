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

//  Este método significa que o tipo T implementa a interface I,
// mas não precisamos declarar explicitamente que isso ocorre.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
