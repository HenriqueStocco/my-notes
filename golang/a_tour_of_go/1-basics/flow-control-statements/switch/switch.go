package main

import (
	"fmt"
	"runtime"
	"time"
)

// Basico
func switchStatement() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("Mac")
	default:
		fmt.Printf("Nenhum acima: %s\n", os)
	}
}

// Switch sem condição é igual a true
func noCondition() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
func main() {
	switchStatement()
	noCondition()
}
