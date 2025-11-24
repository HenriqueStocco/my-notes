package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func maps() {
	// Um map mapeia chaves para valores
	// O valor zero de um map é nil
	// Um map nulo não tem chaves, nem é possivel adicionar chaves
	var m map[string]Vertex

	// A função make retorna um mapa do tipo especificado, inicializado e pronto para uso.
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}

func mapLiterals() {
	// map [key type] struct type{}
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}

func mapLiteralsContinued() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

func mutatingMaps() {
	m := make(map[string]int)

	// Adiciona um elemento
	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	// Atualiza um elemento
	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	// Remove um elemento
	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	// Se a chave existe no m, ok = true, se não ok = false
	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)

	e, ok := m["Response"]
	fmt.Println("The value: ", e, "Present?", ok)
}

func main() {
	maps()
	mapLiterals()
	mapLiteralsContinued()
	mutatingMaps()
}
