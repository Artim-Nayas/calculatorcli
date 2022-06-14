package handler

import (
	"calculatorcli/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestNoopHandler(t *testing.T) {
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	NoopHandler(models.NewCalculator(), 23.22)

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()

	assert.Equal(t, "Invalid operation - current value: 0.0000", string(bytes))
}
