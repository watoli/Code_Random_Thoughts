package Interpreter

import (
	"fmt"
	"testing"
)

func TestAdd_Interpret(t *testing.T) {
	expression := "11 + 2 - 6"
	interpreter := NewInterpreter()
	result := interpreter.Interpret(expression)
	fmt.Printf("%s = %d\n", expression, result)
}
