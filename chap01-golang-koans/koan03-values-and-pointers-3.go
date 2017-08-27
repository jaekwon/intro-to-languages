package main

import "fmt"

type Animal interface {
	Sound() string
}

type Cat struct{}

func (_ Cat) Sound() string { return "meow" }

type Dog struct{}

func (_ *Dog) Sound() string { return "woof" }

func main() {
	var animal Animal
	animal = Cat{}
	fmt.Println(animal.Sound())

	// Why does the following not compile?
	// Make it work for a dog.

	// animal = Dog{}
	// fmt.Println(animal.Sound())
}
