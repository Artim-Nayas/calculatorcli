package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivision(t *testing.T) {
	testSamples := map[float64][]map[float64]float64{
		0.0:  {map[float64]float64{0.1: 0.0, 0.2: 0.0}},
		42.0: {map[float64]float64{-1.0: -42.0, -2.0: -21.0}},
		10.0: {map[float64]float64{5: 2, 8: 1.25, 16: 0.625, 32: 0.3125, 64: 0.1563, 128: 0.0781}},
	}
	for initialValue, testSample := range testSamples {
		t.Run(fmt.Sprintf("when initial value is %f", initialValue), func(t *testing.T) {
			for _, testCase := range testSample {
				for divValue, expectedValue := range testCase {
					t.Run(fmt.Sprintf("upon dividing %f", divValue), func(t *testing.T) {
						calc := calculator{value: initialValue}
						calc.Div(divValue)
						assert.Equal(t, expectedValue, calc.value)
					})
				}
			}
		})
	}
}
