package main

import (
	"calculatorcli/cli_input"
	"calculatorcli/handler"
	"calculatorcli/models"
	"calculatorcli/views"
	"os"
)

func main() {
	calc := models.NewCalculator()
	for true {
		views.RenderCliInput()
		cliInput := cli_input.CliInput(os.Stdin)
		handler.GetHandler(cliInput.Operation())(calc, cliInput.Operand())
	}
}
