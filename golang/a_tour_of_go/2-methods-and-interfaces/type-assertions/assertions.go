package main

import "fmt"

// uma asserção de tipo fornece ao valor concreto implicito de um valor de interface
func assert(i interface{}) {
	s, ok := i.(string)

	fmt.Println(s, ok)
}

func main() {
	var i interface{} = "Hello"
	assert(i)
}
