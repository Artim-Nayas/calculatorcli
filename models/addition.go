package models

func (c *calculator) Add(operand float64) {
	c.value = c.value + operand
}
