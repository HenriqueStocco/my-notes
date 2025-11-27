package json

import (
	c "class006/structs"
	"encoding/json"
	"fmt"
	"log"
)

func Json() {
	cliente := c.ClienteInternacional{
		Cliente: c.Cliente{
			Nome:  "Robsom",
			Email: "r@r.com",
			CPF:   12345678900,
		},
		Pais: "Africa",
	}

	// Transforma a nossa struct em JSON, mas em slice de bytes
	// É necessario converter estes bytes para string
	clienteJson, err := json.Marshal(cliente)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Cliente em JSON: \n", string(clienteJson))

	// Processo ao contrario
	// Pegando um objeto JSON hidratando(transformando) em struct
	cliente2Json := `{"Nome":"Dave","Email":"d@d.com","CPF":12345678900,"pais":"EUA}`
	// Como estou importando a struct, precisei passar como referencia para ser o mesmo
	// endereço de memoria, por causa de escopo
	jsonToStruct := &c.ClienteInternacional{}

	// Unmarshal recebe um array de bytes (que é o nosso json)
	// No caso nós temos a string e por isso convertemos para um array de bytes
	// E recebe a nossa struct que vamos hidratar, mas como referencia
	// pois se passarmos apenas o valor ela não será atribuida para o mesmo endereço que
	// a variavel que queremos.
	json.Unmarshal([]byte(cliente2Json), &jsonToStruct)
	fmt.Println("JSON para Struct: \n", jsonToStruct)
}
