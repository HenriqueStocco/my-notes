package main

import "fmt"

type person struct {
	name string
	age  byte
	city string
}

func (p person) Introduce() string {
	return "Name: " + p.name + ", City: " + p.city
}

func main() {
	json := person{
		name: "Json",
		age:  33,
		city: "Alabama",
	}
	freddie := person{
		name: "Freddie",
		age:  52,
		city: "New York",
	}
	renato := person{
		name: "Renato",
		age:  26,
		city: "Arizona",
	}

	fmt.Println(json.Introduce())
	fmt.Println(freddie.Introduce())
	fmt.Println(renato.Introduce())
}
