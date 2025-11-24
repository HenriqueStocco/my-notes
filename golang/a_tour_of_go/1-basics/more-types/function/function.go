package main

import (
	"fmt"
	"math"
)

func funcValue(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funClosure() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x

		return sum
	}
}

func main() {
	// Função como valor
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(funcValue(hypot))
	fmt.Println(funcValue(math.Pow))

	pos, neg := funClosure(), funClosure()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
