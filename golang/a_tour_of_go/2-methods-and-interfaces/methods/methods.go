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

func main() {
	// v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt2)

	// fmt.Println(v.Abs())
	fmt.Println(f.methodContinued())
}
