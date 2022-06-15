package handler

import (
	"calculatorcli/models"
	"fmt"
)

type Operation string
type HandlerFunc func(c models.Calculator, v float64)

var handlers = map[Operation]HandlerFunc{}

func RegisterHandler(op Operation, h HandlerFunc) {
	if _, found := handlers[op]; found {
		err := fmt.Errorf("duplicate handlers registration for Operation: %s", op)
		panic(err)
	}
	handlers[op] = h
}

func GetHandler(operation Operation) (h HandlerFunc) {
	if h, found := handlers[operation]; found {
		return h
	}
	return NoopHandler
}
