package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
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

func TestDivHandlerOutput(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	calc := models.NewCalculator()
	calc.Add(1)
	GetHandler(divOperation)(calc, 128)

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, "0.0078", string(bytes))

	r, fakeStdout, err = os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout

	GetHandler(divOperation)(calc, 0.01)
	fakeStdout.Close()
	bytes, err = io.ReadAll(r)
	r.Close()
	assert.Equal(t, "0.7800", string(bytes))
}

func TestDivHandlerRegistration(t *testing.T) {
	assert.IsType(t, handlerFunc(DivHandler), GetHandler("div"))
}

func TestDivisionException(t *testing.T) {
	t.Run("divide by 0 not allowed", func(t *testing.T) {
		r, fakeStdout, err := os.Pipe()
		require.NoError(t, err)
		os.Stdout = fakeStdout
		calc := models.NewCalculator()
		GetHandler(divOperation)(calc, 0)
		fakeStdout.Close()
		bytes, err := io.ReadAll(r)
		r.Close()
		assert.Equal(t, "Division by 0 not permitted\n0.0000", string(bytes))

	})
}
