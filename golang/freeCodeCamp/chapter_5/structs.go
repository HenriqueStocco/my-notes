package main

// Utilizamos structs in Go para representar um dado estruturado.
type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

//Nested structs
type Wheel struct {
	Radius   int
	Material int
}

// Nest struct anônima como campos com outras structs
type Something struct {
	Do   string
	Make string
	Some struct {
		Thing int
	}
}

// Embedded structs - compor uma struct com outra, como se fosse a extensão de uma classe
// Mas Go não é uma linguagem orientada a objetos completamente
type mouse struct {
	color string
	size  string
}

type tech struct {
	// incorpora a struct mouse em tech
	// agora todos os campos de mouse estão dentro desta struct tech
	mouse
	space string
}

// Metodos de struct
type rect struct {
	width  int
	height int
}

// Os metodos são apenas funções que possuem um receptor, que é um parâmetro especial
// que, sintaticamente, vem antes do nome da função
// (r rect) é o receptor
func (r rect) area() int {
	return r.width * r.height
}

func main() {
	// Os campos da struct podem ser acesados com .
	myCar := car{}
	myCar.FrontWheel.Radius = 5

	// structs anônimas
	// Normalmente se utiliza structs anônimas quando não será utilizada mais de uma vez
	anotherCar := struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3",
	}
	anotherCar.Model = ""

	// Usando embedded structs
	idk := tech{
		space: "something",
		mouse: mouse{
			color: "black",
			size:  "4x4",
		},
	}

	// Usando os metodos de struct
	r := rect{
		width:  4,
		height: 6,
	}
	r.area()
}
