package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddHandler(t *testing.T) {
	c := models.NewCalculator()
	GetHandler(addOperation)(c, 12.23)

	assert.Equal(t, "12.2300", c.String())

	AddHandler(c, 13.22)
	assert.Equal(t, "25.4500", c.String())
}
