package models

func (c *calculator) Repeat(value int) {
	length := len(history)
	for i := length - value; i < length; i++ {
		history[i].operation(history[i].operand)
	}
}
