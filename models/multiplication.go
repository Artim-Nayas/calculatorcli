package models

import "math"

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (c *calculator) Mul(operand float64) {
	c.value = roundFloat(c.value*operand, 4)
}
