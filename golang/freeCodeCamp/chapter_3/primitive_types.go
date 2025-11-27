package main

import "fmt"

func primitiveTypes() {
	// Tipos primitivos
	// BOOLEANO
	var boolean bool = true

	// CARACTERES & TEXTO
	var text string = "Hello World!"

	// INTEIROS
	// COM SINAL
	// depende da arch, pode ser 32 ou 64 bits
	var signed4b int = 0
	var signed8b int8 = 127
	var signed16b int16 = 32767
	var signed32b int32 = 2147483647
	var signed64b int64 = 9223372036854775807
	var aliasInt32 rune = 2147483647
	// SEM SINAL
	// depende da arch, pode ser 32 ou 64 bits
	var unsigned4b uint = 0
	var unsigned8b uint8 = 255
	var unsigned16b uint16 = 65535
	var unsigned32b uint32 = 4294967295
	var unsiged64b uint64 = 18446744073709551615
	// depende da arch, pode ser 32 ou 64 bits
	// usado somente para armazenagem de ponteiros
	var uintpointer uintptr = 1
	var aliasUint8b byte = 255

	// PONTO FLUTUANTE
	var float32b float32 = 3.1415
	var float64b float64 = 3.1415

	// COMPLEXO (IHNI)
	var complex64b complex64 = 2e2
	var complex128b complex128 = 4e4

	fmt.Println("Boolean: ", boolean)
	fmt.Println("String: ", text)
	fmt.Println("Default Signed Int (32 | 64) bits: ", signed4b)
	fmt.Println("Signed Int 8 bits: ", signed8b)
	fmt.Println("Signed Int 16 bits: ", signed16b)
	fmt.Println("Signed Int 32 bits: ", signed32b)
	fmt.Println("Signed Int 64 bits: ", signed64b)
	fmt.Println("Default Unsigned Int (32 | 64) bits: ", unsigned4b)
	fmt.Println("Unsigned Int 8 bits: ", unsigned8b)
	fmt.Println("Unsigned Int 16 bits: ", unsigned16b)
	fmt.Println("Unsigned Int 32 bits: ", unsigned32b)
	fmt.Println("Unsigned Int 64 bits: ", unsiged64b)
	fmt.Println("Alias Unsigned Int 8 bits: ", aliasUint8b)
	fmt.Println("Alias Signed Int 32 bits: ", aliasInt32)
	fmt.Println("Pointers only Int 32 bits: ", uintpointer)
	fmt.Println("Floating point number 32 bits: ", float32b)
	fmt.Println("Floating point number 64 bits: ", float64b)
	fmt.Println("Complex number 64 bits: ", complex64b)
	fmt.Println("Complex number 128 bits: ", complex128b)
}
