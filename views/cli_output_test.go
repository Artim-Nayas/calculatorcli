package views

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	Render("output should be this")

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, "output should be this", string(bytes))
}

func TestViewRenderInvalidOperation(t *testing.T) {
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	RenderInvalidOperation("2.2300")

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, "Invalid operation - current value: 2.2300", string(bytes))
}

func TestRenderCliInput(t *testing.T) {
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = fakeStdout
	RenderCliInput()

	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, "\n>", string(bytes))
}
