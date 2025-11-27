package main

import "fmt"

// Type switches é uma estrutura que permite várias asserções de tipo em série
// func do(i interface{}) {
func do(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I dont't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
