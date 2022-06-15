package handler

import (
	"calculatorcli/models"
	"calculatorcli/views"
)

const subOperation Operation = "sub"

func SubHandler(c models.Calculator, v float64) {
	c.Sub(v)
	views.Render(c.String())
}

func init() {
	RegisterHandler(subOperation, SubHandler)
}
