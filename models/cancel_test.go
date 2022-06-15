package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCancelFunction(t *testing.T) {
	t.Run("should change value to 0", func(t *testing.T) {
		calc := calculator{10}
		calc.Cancel()
		assert.Equal(t, 0.0, calc.value)
	})
}
