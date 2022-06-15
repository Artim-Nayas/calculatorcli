package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	t.Run("should be able to register handlers with operation and handler func", func(t *testing.T) {
		defer resetHandler()()
		assert.NotPanics(t, func() {
			RegisterHandler("add", func(c models.Calculator, v float64) {})
		})
	})

	t.Run("should panic if Operation already exists", func(t *testing.T) {
		defer resetHandler()()
		RegisterHandler("add", func(c models.Calculator, v float64) {})

		assert.Panics(t, func() {
			RegisterHandler("add", func(c models.Calculator, v float64) {})
		})
	})
}

func resetHandler() func() {
	handlers = map[Operation]handlerFunc{}
	return func() { handlers = map[Operation]handlerFunc{} }
}

func TestGetHandler(t *testing.T) {
	t.Run("should return a handler given a Operation", func(t *testing.T) {
		defer resetHandler()()
		RegisterHandler("add", func(c models.Calculator, v float64) {})

		assert.IsType(t, *new(handlerFunc), GetHandler("add"))
	})

	t.Run("should return noop handler when Operation is not found", func(t *testing.T) {
		defer resetHandler()()

		h := GetHandler("add")
		assert.IsType(t, handlerFunc(NoopHandler), h)
	})
}
