package models

import (
	"fmt"
	"math"
	"os"
)

type calculator struct {
	value float64
}

func (c *calculator) String() string {
	return fmt.Sprintf("%0.4f", c.value)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

type Calculator interface {
	fmt.Stringer
	Add(operand float64)
	Sub(operand float64)
	Mul(operand float64)
	Div(operand float64)
	Cancel()
	Repeat(value int)
}

func NewCalculator() *calculator {
	return &calculator{0}
}

func init() {
	os.Create("history")
}
