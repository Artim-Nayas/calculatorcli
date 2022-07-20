package models

func (c *calculator) Mul(operand float64) {
	c.SaveInput(c.Mul, operand)
	c.value = roundFloat(c.value*operand, 4)
}
