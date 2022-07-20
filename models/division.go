package models

func (c *calculator) Div(operand float64) {
	c.SaveInput(c.Div, operand)
	c.value = roundFloat(c.value/operand, 4)
}
