package models

type operationFunc func(value float64)

type input struct {
	operation operationFunc
	operand   float64
}

var history []input

func (c *calculator) SaveInput(operation operationFunc, value float64) {
	history = append(history, input{operation, value})
}

func LengthHistory() float64 {
	return float64(len(history))
}
