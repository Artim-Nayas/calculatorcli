package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubHandler(t *testing.T) {
	resetHandler()
	c := models.NewCalculator()
	RegisterHandler(subOperation, SubHandler)
	GetHandler(subOperation)(c, 12.23)
	assert.Equal(t, "-12.2300", c.String())

	SubHandler(c, 13.22)
	assert.Equal(t, "-25.4500", c.String())
}
