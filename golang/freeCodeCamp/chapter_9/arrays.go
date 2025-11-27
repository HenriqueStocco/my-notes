package main

// Arrays - São grupos de tamanho fixo de variaveis de mesmo tipo

func main() {
	// Declarando um array com 10 itens
	// Se não declarar o valor é tudo 0
	var myInts [10]int8

	for i := 0; i < 10; i++ {
		println(myInts[i])
	}

	// Inicializando um array
	prime := [4]int{2, 3, 5, 7}
	println(prime[0])
}
