package main

import "fmt"

func main() {

	// Why does this work?
	mutateStructGood()

	// Why does the following throw a compilation error?
	// mutateStructBad()

}

//----------------------------------------

type MyStruct struct {
	value int
}

func NewMyStruct(value int) MyStruct {
	return MyStruct{
		value: value,
	}
}

func (ms *MyStruct) Increment() int {
	ms.value += 1
	return ms.value
}

func (ms *MyStruct) AssertEquals(value int) {
	if ms.value != value {
		panic(fmt.Sprintf("Expected %v but ms.value is %v", value, ms.value))
	}
}

func mutateStructGood() {
	ms := NewMyStruct(0)
	ms.AssertEquals(0)
	ms.Increment()
	ms.AssertEquals(1)
}

/*
func mutateStructBad() {
	NewMyStruct(0).Increment()
}
*/
