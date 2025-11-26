package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go não tem classe e POO, mas é possivel definir metodos para structs/types
// Um metodo é uma função com um argumento receptor especial
// Este metodo Abs é um metodo que tem o receptor do tipo Vertex nomeado v
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

// É possivel definir metodos com receptores que não são structs também
// Obs: não é possivel definir um metodo onde o tipo está em outro arquivo, tem que ser no mesmo.
func (f MyFloat) methodContinued() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// É possivel criar metodos com receptores de ponteiros
// Como os metodos frequentemente precisam modificar seu receptor, receptores de ponteiros são mais comuns
func (v *Vertex) methodPointerReceiver(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Diferenças entre um metodo que recebe um receptor de ponteiro
// e uma função que recebe um ponteiro de uma struct
// Metodo
// métodos com receptores de ponteiro aceitam um valor ou um ponteiro como receptor
// quando são chamados
// Por conveniência, o Go interpreta a instrução v.methodOfStructPointerReceiver(2)
// como (&v).methodOfStructPointerReceiver(2)
func (v *Vertex) methodOfStructPointerReceiver(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Função
// funções com um argumento de ponteiro devem receber um ponteiro
// se não da erro na compilação
func functionPointerParam(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// O mesmo acontece quando utilizamos receptores como valores em comparação
// com funções que recebem um argumento, no caso da função, ele deve receber algo exato
// referente ao tipo do parametro, o metodo com receptor pode receber o valor ou ponteiro
// Metodo
func (v Vertex) methodValueReceptor() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Função
func functionReceiveArg(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Há duas razões para usar um receptor de ponteiro.
// A primeira é para que o método possa modificar o valor para o qual seu receptor aponta.
// A segunda é para evitar copiar o valor em cada chamada de método.
// Isso pode ser mais eficiente se o receptor for uma estrutura grande

func main() {
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt2)
	v.methodPointerReceiver(10)
	v.methodOfStructPointerReceiver(2)
	functionPointerParam(&v, 10)
	fmt.Println(v.methodValueReceptor())
	fmt.Println(functionReceiveArg(v))

	p := &Vertex{4, 3}
	fmt.Println(p.methodValueReceptor())
	fmt.Println(functionReceiveArg(*p))

	p.methodOfStructPointerReceiver(3)
	functionPointerParam(p, 8)

	fmt.Println(v, p)

	fmt.Println(v.Abs())
	fmt.Println(v.Abs())
	fmt.Println(f.methodContinued())
}
