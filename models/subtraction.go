package models

func (c *calculator) Sub(operand float64) {
	c.value = c.value - operand
}
