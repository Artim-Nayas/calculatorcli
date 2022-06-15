package models

func (c *calculator) Div(operand float64) {
	c.value = roundFloat(c.value/operand, 4)
}
