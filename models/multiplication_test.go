package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	testSamples := map[float64][]map[float64]float64{
		0.0:     {map[float64]float64{0.1: 0, 0.2: 0}},
		42.0:    {map[float64]float64{-1.0: -42.0, 2.0: 84.0}},
		10.0001: {map[float64]float64{10.0001: 100.002, 20: 200.002, 2: 20.0002}},
	}
	for initialValue, testSample := range testSamples {
		t.Run(fmt.Sprintf("when initial value is %f", initialValue), func(t *testing.T) {
			for _, testCase := range testSample {
				for mulValue, expectedValue := range testCase {
					t.Run(fmt.Sprintf("upon multiplying %f", mulValue), func(t *testing.T) {
						calc := calculator{value: initialValue}
						calc.Mul(mulValue)
						assert.Equal(t, expectedValue, calc.value)
					})
				}
			}
		})
	}
}
