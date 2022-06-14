package views

import "fmt"

func Render(output string) {
	fmt.Printf(output)
}

func RenderInvalidOperation(value float64) {
	fmt.Printf("invalid operations; current value: %0.4f", value)
}
