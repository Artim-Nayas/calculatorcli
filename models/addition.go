package models

func (c *calculator) Add(operand float64) {
	c.SaveInput(c.Add, operand)
	c.value = c.value + operand
}
