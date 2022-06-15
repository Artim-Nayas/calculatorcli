package models

import "fmt"

type calculator struct {
	value float64
}

func (c *calculator) String() string {
	return fmt.Sprintf("%0.4f", c.value)
}

type Calculator interface {
	fmt.Stringer
	Add(operand float64)
	Sub(operand float64)
}

func NewCalculator() *calculator {
	return &calculator{0}
}
