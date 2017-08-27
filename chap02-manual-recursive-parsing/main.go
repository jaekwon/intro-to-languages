package main

import "fmt"
import "io/ioutil"

func main() {
}

//--------------------------------------------------------------------------------

type Node interface {
	Type() string         // Name of type of node, e.g. 'BinaryExpression', 'MonoidExpression'
	ParseInfo() ParseInfo // e.g. Start/end position of node in the original program
}

type ParseInfo struct {
	Start int
	End   int // exclusive
}

type Program struct {
	Nodes []Node
}

type Comment struct {
	info ParseInfo

	Comment string
}

type Expression interface {
	Evaluate() float64 // Compute the result of this algebraic expression
}

type BinaryExpression struct {
	info ParseInfo

	Left     Expression
	Right    Expression
	Operator string // e.g. '+', '-', '*', '/', '%'
}

type MonoidExpression struct {
	info ParseInfo

	PrefixOperator string // e.g. '-' for negation
	Expression     Expression
}

type ConstantExpression struct {
	info ParseInfo

	Value float64
}

//--------------------------------------------------------------------------------

func readProgram(filename string) (node Node, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	prog := string(bytes)
	fmt.Printf("Read program: %v\n", prog)

	var cursor = 0 // Where we're currently reading.

OUTER_LOOP:
	for true {

		// Try reading comment
		node, end := readComment(prog, cursor)
		if node != nil {
			fmt.Println("Comment: %v\n", node.(
			cursor = end
			continue OUTER_LOOP // NOTE: labels are optional
		}

		// Try reading expression
		node, end

	}
}

func readExpression(prog string, start int) (node Node, end int) {
}
