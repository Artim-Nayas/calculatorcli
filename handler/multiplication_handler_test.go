package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
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

func TestMulHandlerOutput(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	calc := models.NewCalculator()
	calc.Add(1)
	GetHandler(mulOperation)(calc, 23.22245)

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, "23.2224", string(bytes))

	r, fakeStdout, err = os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout

	GetHandler(mulOperation)(calc, 0.01)
	fakeStdout.Close()
	bytes, err = io.ReadAll(r)
	r.Close()
	assert.Equal(t, "0.2322", string(bytes))
}
