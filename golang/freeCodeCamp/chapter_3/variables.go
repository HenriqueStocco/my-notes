package main

import "fmt"

func variables() {
	// Variaveis com var precisam da definição explicita do tipo dela, como cada um abaixo
	// Declared variable
	var explicit string = "This is a string"

	fmt.Println("Variable with explicit type: ", explicit)

	// Variaveis com sem o var e com o operador de atribuição :=, indica uma variavel que possui o seu tipo definido implicitamente
	// Short variable declaration
	implicit := "This is a string too"

	fmt.Println("Variable with implicit type: ", implicit)

	// Constantes, variáveis não reatribuiveis
	// Não é possivel utilizar a atribuição implicita de tipo com :=
	const constant uint8 = 24
	fmt.Println("Constant variable: ", constant)
}
