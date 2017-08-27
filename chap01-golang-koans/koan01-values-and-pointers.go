package main

import "fmt"

func main() {

	mutateStructGood()

	// Why does this fail, and how would you fix it?
	mutateStructBad()

}

//----------------------------------------

type MyStruct1 struct {
	value int
}

func NewMyStruct1(value int) *MyStruct1 {
	return &MyStruct1{
		value: value,
	}
}

func (ms *MyStruct1) Increment() int {
	ms.value += 1
	return ms.value
}

func (ms *MyStruct1) AssertEquals(value int) {
	if ms.value != value {
		panic(fmt.Sprintf("Expected %v but ms.value is %v", value, ms.value))
	}
}

func mutateStructGood() {
	ms := NewMyStruct1(0)
	ms.AssertEquals(0)
	ms.Increment()
	ms.AssertEquals(1)
}

//----------------------------------------

type MyStruct2 struct {
	value int
}

func NewMyStruct2(value int) MyStruct2 {
	return MyStruct2{
		value: value,
	}
}

func (ms MyStruct2) Increment() int {
	ms.value += 1
	return ms.value
}

func (ms MyStruct2) AssertEquals(value int) {
	if ms.value != value {
		panic(fmt.Sprintf("Expected %v but ms.value is %v", value, ms.value))
	}
}

func mutateStructBad() {
	ms := NewMyStruct2(0)
	ms.AssertEquals(0)
	ms.Increment()
	ms.AssertEquals(1)
}
