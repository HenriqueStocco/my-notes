package maths

/*
	O que indica um export para o Go é a primeira letra do nome de uma função sendo maiúscula, se for minuscula o Go não deixa usar a função fora do arquivo

	É recomendado que toda exportação tenha uma documentação (comentario em cima)
*/

/*
Soma: função que faz a soma de dois valores inteiros
@example: a + b
*/
func Soma(a int, b int) int {
	return a + b
}
