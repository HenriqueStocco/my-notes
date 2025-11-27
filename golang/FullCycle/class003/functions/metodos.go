package functions

import "fmt"

// Tipo
type Car struct {
	Name string
}

// Metodo da struct Car, chamada de Bind
func (c Car) run() {
	fmt.Println(c.Name, "Andou")
}

// Metodos: Funcao para usar o metodo run da struct Car
func Metodos() {
	car := Car{
		Name: "Rodolfo",
	}
	car.run()
}
