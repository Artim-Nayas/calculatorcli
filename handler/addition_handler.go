package handler

import (
	"calculatorcli/models"
	"calculatorcli/views"
)

const addOperation Operation = "add"

func AddHandler(c models.Calculator, v float64) {
	c.Add(v)
	views.Render(c.String())
}

func init() {
	RegisterHandler(addOperation, AddHandler)
}
