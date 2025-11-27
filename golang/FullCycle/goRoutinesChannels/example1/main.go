package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	contador("Sem Go routine")
	go contador("Com Go routine")

	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
	time.Sleep(time.Second)
	fmt.Println("Fim...")
}
