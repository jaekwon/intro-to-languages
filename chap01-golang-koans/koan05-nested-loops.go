package main

import (
	"fmt"
	"reflect"
)

// This program should loop over elements of an array,
// and keep a sum of all the numbers until it sees "stop".

func main() {

	array := []interface{}{1, 2, 3, 4, "stop", 5, 6}
	tally := 0

	for _, element := range array {
		switch e := element.(type) {
		case int:
			tally += e
		case string:
			if e == "stop" {
				break
			} else {
				panic(fmt.Sprintf("Unrecognized string %v", e))
			}
		default:
			panic(fmt.Sprintf("Unrecognized element type %v", reflect.TypeOf(element)))
		}
	}

	// Fix this program so it gets the correct tally.
	if tally == 10 {
		fmt.Println("Got the correct tally")
	} else {
		fmt.Printf("Expected tally of 10 but got %v\n", tally)
	}
}
