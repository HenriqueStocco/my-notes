package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrucao string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e eu ja arranquei", x.dentesarrancados, "dentes, e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()

	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salario)
	case arquiteto:
		fmt.Println("Eu construo:", g.(arquiteto).tipodeconstrucao)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Mister Dente",
			sobrenome: "Da Silva",
			idade:     50,
		},
		dentesarrancados: 8973,
		salario:          536.56,
	}
	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Mister Prédio",
			sobrenome: "Sobrenome",
			idade:     51,
		},
		tipodeconstrucao: "Hospícios",
		tamanhodaloucura: "Demais",
	}

	// Neste exemplo, cada tipo é diferente e cada método também
	// Método especifico de cada struct/interface
	// mrdente.oibomdia()
	// mrpredio.oibomdia()

	// Aqui é o mesmo tipo mas que recebe um método diferente
	// Uma única interface para diferentes tipos
	// Isso aqui é polimorfismo, onde eu posso ter o mesmo metodo/funcao mas com a funcionalidade diferente
	serhumano(mrdente)
	serhumano(mrpredio)
}
