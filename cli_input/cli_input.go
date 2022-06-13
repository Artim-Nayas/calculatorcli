package cli_input

import (
	"fmt"
	"io"
)

type Input struct {
	operation string
	operand   float64
}

func CliInput(reader io.Reader) Input {
	var operation string
	var operand float64
	fmt.Fscanf(reader, "%s %f", &operation, &operand)
	return Input{operation, operand}
}
