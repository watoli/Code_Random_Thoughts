package Interpreter

import (
	"strconv"
	"strings"
)

// Expression 定义抽象表达式接口
type Expression interface {
	Interpret() int
}

// NumberExpression 定义终结符表达式
type NumberExpression struct {
	number int
}

func (n *NumberExpression) Interpret() int {
	return n.number
}

// PlusExpression 定义非终结符表达式
type PlusExpression struct {
	left  Expression
	right Expression
}

func (p *PlusExpression) Interpret() int {
	return p.left.Interpret() + p.right.Interpret()
}

type MinusExpression struct {
	left  Expression
	right Expression
}

func (m *MinusExpression) Interpret() int {
	return m.left.Interpret() - m.right.Interpret()
}

// Context 定义环境
type Context struct {
	expressionStack []Expression
}

func (c *Context) PushExpression(expr Expression) {
	c.expressionStack = append(c.expressionStack, expr)
}

func (c *Context) PopExpression() Expression {
	if len(c.expressionStack) > 0 {
		expr := c.expressionStack[len(c.expressionStack)-1]
		c.expressionStack = c.expressionStack[:len(c.expressionStack)-1]
		return expr
	}
	return nil
}

// Interpreter 定义解释器
type Interpreter struct {
	context *Context
}

func NewInterpreter() *Interpreter {
	return &Interpreter{&Context{}}
}

func (i *Interpreter) Interpret(expression string) int {
	tokens := strings.Split(expression, " ")
	var output []string
	var operators []string
	for _, token := range tokens {
		if _, err := strconv.Atoi(token); err == nil {
			output = append(output, token)
		} else {
			switch token {
			case "+", "-":
				for len(operators) > 0 && (operators[len(operators)-1] == "+" || operators[len(operators)-1] == "-") {
					output = append(output, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}
				operators = append(operators, token)
			case "(", ")":
				if token == "(" {
					operators = append(operators, token)
				} else {
					for operators[len(operators)-1] != "(" {
						output = append(output, operators[len(operators)-1])
						operators = operators[:len(operators)-1]
					}
					operators = operators[:len(operators)-1]
				}
			}
		}
	}
	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	for _, token := range output {
		if n, err := strconv.Atoi(token); err == nil {
			i.context.PushExpression(&NumberExpression{n})
		} else {
			var left, right Expression
			right = i.context.PopExpression()
			left = i.context.PopExpression()
			switch token {
			case "+":
				i.context.PushExpression(&PlusExpression{left, right})
			case "-":
				i.context.PushExpression(&MinusExpression{left, right})
			}
		}
	}
	result := i.context.PopExpression()
	return result.Interpret()
}
