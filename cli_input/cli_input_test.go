package cli_input

import (
	"calculatorcli/handler"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestCliInput(t *testing.T) {
	t.Run("should accept operation name and numerical operand", func(t *testing.T) {
		var reader io.Reader = strings.NewReader("operation 12.34")
		input := CliInput(reader)
		assert.Equal(t, handler.Operation("operation"), input.Operation())
		assert.Equal(t, 12.34, input.Operand())
	})
}
