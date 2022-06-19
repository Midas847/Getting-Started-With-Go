package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string) *Person {
	p := Person{name: name}
	p.age = 40
	return &p
}

func main() {
	fmt.Println(Person{"Saswata", 20})
	fmt.Println(NewPerson("Saswata"))
}