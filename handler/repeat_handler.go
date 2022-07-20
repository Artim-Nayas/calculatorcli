package handler

import (
	"calculatorcli/models"
	"calculatorcli/views"
	"fmt"
	"math"
)

const repeatOperation Operation = "repeat"

func RepeatHandler(c models.Calculator, v float64) {
	if math.Mod(v, 1) == 0 && v <= models.LengthHistory() {
		c.Repeat(int(v))
		views.Render(c.String())
	} else if math.Mod(v, 1) != 0 {
		fmt.Printf("Repeat needs whole numbers as argument")
	} else {
		fmt.Printf("Cannot implement repeat %0.0f;\t Number of mathematical commands executed: %0.0f", v, models.LengthHistory())
	}
}

func init() {
	RegisterHandler(repeatOperation, RepeatHandler)
}
