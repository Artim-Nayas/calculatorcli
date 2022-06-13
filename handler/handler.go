package handler

import (
	"calculatorcli/models"
)

type Operation string
type handlerFunc func(c models.Calculator, v float64)

func RegisterHandler(op Operation, h handlerFunc) {
}
