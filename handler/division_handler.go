package handler

import (
	"calculatorcli/models"
	"calculatorcli/views"
)

const divOperation Operation = "div"

func DivHandler(c models.Calculator, v float64) {
	c.Div(v)
	views.Render(c.String())
}

func init() {
	RegisterHandler(divOperation, DivHandler)
}
