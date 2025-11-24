package main

import "fmt"

// structs são coleções de campos
type Vertex struct {
	X, Y int
}

func printVertex() {
	fmt.Println(Vertex{1, 2})
}
func accessingStructFields() {
	s := Vertex{1, 2}
	fmt.Println(s.X, s.Y)
}

func pointersToStruct() {
	v := Vertex{4, 5}
	p := &v
	// o Go permite utilizar esta notação simples para acessar um ponteiro
	p.X = 1e9
	// Mas poderia ser feito assim, de forma mais complicada
	(*p).Y = 45

	fmt.Println(v)
}

func structLiterals() {
	// Structs literais
	var (
		v1 = Vertex{1, 2}  // struct de modo literal, com X: 1 e Y: 2
		v2 = Vertex{X: 9}  // definir Y de modo implicito (valor 0), pode ser feito em qualquer ordem, ex Y e X
		v3 = Vertex{}      // struct literal com os valores zeros: {0,0}
		p  = &Vertex{8, 7} // ponteiro para a struct literal
	)

	// Com um exemplo confuso deste, segue um exemplo de uma struct não literal
	var s Vertex
	s.X = 0
	s.Y = 12

	fmt.Println(v1, v2, v3, p, s)
}

func main() {
	printVertex()
	accessingStructFields()
	pointersToStruct()
	structLiterals()
}
