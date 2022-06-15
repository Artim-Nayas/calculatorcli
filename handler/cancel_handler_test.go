package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestCancelHandler(t *testing.T) {
	c := models.NewCalculator()
	GetHandler(addOperation)(c, 12.23)

	assert.Equal(t, "12.2300", c.String())

	GetHandler(cancelOperation)(c, 0)
	assert.Equal(t, "0.0000", c.String())
}

func TestCancelHandlerOutput(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	calc := models.NewCalculator()
	GetHandler(cancelOperation)(calc, 0)

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, "0.0000", string(bytes))
}

func TestCancelHandlerRegistration(t *testing.T) {
	assert.IsType(t, handlerFunc(CancelHandler), GetHandler("cancel"))
}
