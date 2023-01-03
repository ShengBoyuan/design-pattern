package strategy_test

import (
	"design-pattern/behavioral/strategy"
	"fmt"
	"testing"
)

func TestCalculateNums(t *testing.T) {
	a, b := 1, 3
	c := strategy.CalculateNums(a, b, "addition")
	fmt.Println(c)
	c = strategy.CalculateNums(a, b, "subtraction")
	fmt.Println(c)
	c = strategy.CalculateNums(a, b, "multiplication")
	fmt.Println(c)
	_ = strategy.CalculateNums(a, b, "division")
}
