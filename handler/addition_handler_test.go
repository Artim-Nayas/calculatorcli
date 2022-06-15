package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestAddHandler(t *testing.T) {
	c := models.NewCalculator()
	GetHandler(addOperation)(c, 12.23)

	assert.Equal(t, "12.2300", c.String())

	AddHandler(c, 13.22)
	assert.Equal(t, "25.4500", c.String())
}

func TestAddHandlerOutput(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	calc := models.NewCalculator()
	GetHandler(addOperation)(calc, 23.22)

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, "23.2200", string(bytes))

	r, fakeStdout, err = os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout

	GetHandler(addOperation)(calc, 22.23)
	fakeStdout.Close()
	bytes, err = io.ReadAll(r)
	r.Close()
	assert.Equal(t, "45.4500", string(bytes))
}

func TestAddHandlerRegistration(t *testing.T) {
	assert.IsType(t, handlerFunc(AddHandler), GetHandler("Add"))
}
