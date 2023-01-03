package strategy

import "fmt"

// Strategy
type Strategy interface {
	execute(int, int) int
}

// Concrete Strategies
type Add struct{}
type Subtract struct{}
type Multiply struct{}

var _ Strategy = Add{}
var _ Strategy = Subtract{}
var _ Strategy = Multiply{}

func (Add) execute(a int, b int) int {
	return a + b
}

func (Subtract) execute(a int, b int) int {
	return a - b
}

func (Multiply) execute(a int, b int) int {
	return a * b
}

// Context
type Context struct {
	strategy Strategy
}

func (c *Context) setStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) execute(a int, b int) int {
	return c.strategy.execute(a, b)
}

// Client
func CalculateNums(a int, b int, method string) int {
	context := &Context{}
	switch method {
	case "addition":
		context.setStrategy(Add{})
	case "subtraction":
		context.setStrategy(Subtract{})
	case "multiplication":
		context.setStrategy(Multiply{})
	default:
		fmt.Println("no such calculation method")
		return -1
	}

	return context.strategy.execute(a, b)
}
