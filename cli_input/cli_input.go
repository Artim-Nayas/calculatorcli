package cli_input

import (
	"calculatorcli/handler"
	"fmt"
	"io"
)

type Parser interface {
	Operation() handler.Operation
	Operand() float64
}

type input struct {
	operation handler.Operation
	operand   float64
}

func (i input) Operation() handler.Operation {
	return i.operation
}

func (i input) Operand() float64 {
	return i.operand
}

func CliInput(reader io.Reader) (cliInput input) {
	fmt.Fscanf(reader, "%s %f", &cliInput.operation, &cliInput.operand)
	return
}
