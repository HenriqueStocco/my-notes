package main

import "fmt"

type calculator struct{}

func (c calculator) Addition(n1, n2 float32) float32 {
	return n1 + n2
}
func (c calculator) Subtraction(n1, n2 float32) float32 {
	return n1 - n2
}
func (c calculator) Division(n1, n2 float32) float32 {
	if n2 == 0 {
		return 0
	}

	return n1 / n2
}
func (c calculator) Multiplication(n1, n2 float32) float32 {
	if n1 == 0 || n2 == 0 {
		return 0
	}
	return n1 * n2
}

func main() {
	c := calculator{}

	fmt.Printf("Addition - Result %.2f\n", c.Addition(2, 2))
	fmt.Printf("Subtraction - Result %.2f\n", c.Subtraction(2, 2))
	fmt.Printf("Division - Result %.2f\n", c.Division(2, 2))
	fmt.Printf("Multiplication - Result %.2f\n", c.Multiplication(2, 2))
}
