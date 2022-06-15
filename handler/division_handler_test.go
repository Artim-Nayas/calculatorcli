package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivHandler(t *testing.T) {
	c := models.NewCalculator()
	AddHandler(c, 10)
	GetHandler(divOperation)(c, 5)
	assert.Equal(t, "2.0000", c.String())

	DivHandler(c, 1000)
	assert.Equal(t, "0.0020", c.String())
}
