package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExitHandlerRegistration(t *testing.T) {
	assert.IsType(t, handlerFunc(ExitHandler), GetHandler("exit"))
}
