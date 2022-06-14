package views

import "fmt"

func Render(output string) {
	fmt.Printf(output)
}

func RenderInvalidOperation(value string) {
	fmt.Printf("Invalid operation - current value: %s", value)
}

func RenderCliInput() {
	fmt.Printf("\n>")
}
