package main

import "fmt"

func main() {
	var username string = "wagslane"
	var password int = 20947382822

    // Go é de tipagem forte, ou seja, nenhum tipo definido pode ser alterado
    // uma prova é a tentativa de concatenar uma string com um inteiro
	fmt.Println("Authorization: Basic", username+":"+password)
}
