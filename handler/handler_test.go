package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	t.Run("should be able to register handlers with operation and handlers func", func(t *testing.T) {
		assert.NotPanics(t, func() {
			RegisterHandler("add", func(c models.Calculator, v float64) {})
		})
	})
}
