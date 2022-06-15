package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMulHandler(t *testing.T) {
	c := models.NewCalculator()
	AddHandler(c, 1)
	GetHandler(mulOperation)(c, 12.23)
	assert.Equal(t, "12.2300", c.String())

	MulHandler(c, 2)
	assert.Equal(t, "24.4600", c.String())
}
