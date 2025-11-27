package main

import "errors"

func functions(x int, y int) int {
	return x - y
}

func multipleParams(x, y int) int {
	return x + y
}

// Ignorando um retorno de função

func getParams() (x, y int) {
	return 3, 4
}

// Não é necessário return explicitamente algo com a keyword return, pois definimos os nomes e o tipo do retorno da função
// chamado de returno implicito
func namedReturnValues() (x, y int) {
	return
}

// Seria o mesmo que esta função
// chamado de retorno explicito - blank return
func copiaDeCima() (int, int) {
	var x int
	var y int
	return x, y
}

/*
O Maior motivo para utilizar retornos noemados é a documentação e o fácil entendimento
Ajuda em alguns casos a diminuir o código
Early return - quando temos apenas um if, onde normalmente invertemos a lógica da condição
mantendo apenas um if e caso esta condição não seja atingida, retorna outra coisa que seria um else
*/
func earlyReturn(divided, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return divided / divisor, nil
}

func main() {
	// _ ignora
	x, _ := getParams()

	print(x)
}
