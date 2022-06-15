package handler

import (
	"calculatorcli/models"
	"calculatorcli/views"
	"fmt"
)

const divOperation Operation = "div"

func DivHandler(c models.Calculator, v float64) {
	if v == 0 {
		fmt.Print("Division by 0 not permitted\n")
		views.Render(c.String())
	} else {
		c.Div(v)
		views.Render(c.String())
	}
}

func init() {
	RegisterHandler(divOperation, DivHandler)
}
