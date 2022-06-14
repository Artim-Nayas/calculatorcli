package handler

import (
	"calculatorcli/models"
	"fmt"
)

type Operation string
type handlerFunc func(c models.Calculator, v float64)

var handlers = map[Operation]handlerFunc{}

func RegisterHandler(op Operation, h handlerFunc) {
	if _, found := handlers[op]; found {
		err := fmt.Errorf("duplicate handlers registration for Operation: %s", op)
		panic(err)
	}
	handlers[op] = h
}

func GetHandler(operation Operation) (h handlerFunc) {
	return h
}
