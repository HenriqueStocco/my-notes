package main

// o m na frente do pacote math, indica um alias para a importação
import (
	"fmt"
	m "math"
)

// Coisas a serem exportadas, o nome é sempre capitalizado, como o Pi abaixo
func main() {
	fmt.Printf("Pi = %.4f\n", m.Pi)
}
