package handler

import (
	"calculatorcli/models"
	"calculatorcli/views"
)

const cancelOperation Operation = "cancel"

func CancelHandler(c models.Calculator, v float64) {
	c.Cancel()
	views.Render(c.String())
}

func init() {
	RegisterHandler(cancelOperation, CancelHandler)
}
