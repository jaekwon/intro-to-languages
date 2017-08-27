package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

func (_ *Dog) Sound() string { return "woof" }

func main() {
	var animal Animal
	if animal == nil {
		fmt.Println("Initially, the 'animal' variable is a nil interface")
	}

	var dog *Dog = &Dog{}
	animal = dog
	if animal != nil {
		fmt.Println("Now 'animal' is not nil, it's a Dog{}.")
	}

	dog = nil
	if dog == nil {
		fmt.Println("Now 'dog' is nil")
	}

	animal = dog
	if animal == dog {
		fmt.Println("Now 'animal' was set to that nil dog")
	}

	if animal == nil {
		fmt.Println("'animal' is a nil interface")
	} else {
		// What is going on?
		fmt.Println("'animal' is not a nil interface")
		fmt.Printf("'animal.Sound()' is %v\n", animal.Sound())
	}
}
