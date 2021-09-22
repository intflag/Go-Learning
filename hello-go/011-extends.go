package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) Name() string {
	return a.name
}

type Dog struct {
	Animal
}

func main() {
	dog := Dog{}
	dog.name = "dog"
	fmt.Println(dog.name)
	fmt.Println(dog.Name())
}