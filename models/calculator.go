package models

type calculator struct {
	value float64
}

func NewCalculator() *calculator {
	return &calculator{0}
}
