package handler

import (
	"calculatorcli/models"
	"os"
)

const exitOperation Operation = "exit"

func ExitHandler(c models.Calculator, v float64) {
	os.Exit(0)
}

func init() {
	RegisterHandler(exitOperation, ExitHandler)
}
