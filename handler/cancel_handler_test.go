package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCancelHandler(t *testing.T) {
	c := models.NewCalculator()
	GetHandler(addOperation)(c, 12.23)

	assert.Equal(t, "12.2300", c.String())

	GetHandler(cancelOperation)(c, 0)
	assert.Equal(t, "0.0000", c.String())
}
