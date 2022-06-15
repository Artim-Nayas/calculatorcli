package handler

import (
	"calculatorcli/models"
	"calculatorcli/views"
)

func NoopHandler(c models.Calculator, value float64) {
	views.RenderInvalidOperation(c.String())
}
