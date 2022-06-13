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
}

func NewCalculator() *calculator {
	return &calculator{0}
}
