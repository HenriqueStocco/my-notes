package errors

import (
	e "errors"
	f "fmt"
	"log"
	"net/http"
)

func ErrorHandler() {
	res, err := http.Get("https://google.com.br")

	if err != nil {
		log.Fatal("Erro ao fazer comunicação", err.Error())
	}
	f.Printf("Status: %v\n", res.Status)

	/*
		Podemos dar by pass no erro se não quisermos tratar, com blank indentifier, ao invés de passar a variavel de erro, passamos _ e apagamos a tratativa com if, e ai o erro é jogado fora, mas ai o retorno no caso de erro será o valor definido na condição
	*/
	result, _ := soma(2, 8)

	// if err != nil {
	// 	log.Fatal(error.Error())
	// }
	f.Printf("Result: %v \n", result)
}

func soma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, e.New("total maior que 10")
	}

	return res, nil
}
