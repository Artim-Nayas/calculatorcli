package models

func (c *calculator) Cancel() {
	c.value = 0
}
