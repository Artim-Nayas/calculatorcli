package models

func (c *calculator) Sub(operand float64) {
	c.SaveInput(c.Sub, operand)
	c.value = c.value - operand
}
