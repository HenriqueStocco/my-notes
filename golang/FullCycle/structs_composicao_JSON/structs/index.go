package structs

import "fmt"

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo metodo: ", c.Nome)
}

// Compondo ClienteInternacional com Cliente
// Dentro das crases é uma TAG, para o json
// Quando o json for converter a struct para JSON, ao invés de pegar o nome que está
// na struct, ele identifica esta tag e usa a tag como chave
type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"`
}

func Structs() {
	cliente := Cliente{
		Nome:  "Henrique",
		Email: "h@h.com.br",
		CPF:   12345678900,
	}
	// fmt.Println(cliente)

	cliente2 := Cliente{"Mari", "m@m.com", 123456}
	// fmt.Printf("Nome: %s. Email: %s. CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Dave",
			Email: "d@d.com",
			CPF:   12345678900,
		},
		Pais: "EUA",
	}
	// fmt.Printf("Nome: %s, Email: %s, CPF: %d, País: %s.\n", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)

	cliente.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()
}
