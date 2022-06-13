package models

type calculator struct {
	value float64
}

type Calculator interface {
}

func NewCalculator() *calculator {
	return &calculator{0}
}
