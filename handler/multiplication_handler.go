package handler

import (
	"calculatorcli/models"
	"calculatorcli/views"
)

const mulOperation Operation = "mul"

func MulHandler(c models.Calculator, v float64) {
	c.Mul(v)
	views.Render(c.String())
}

func init() {
	RegisterHandler(mulOperation, MulHandler)
}
