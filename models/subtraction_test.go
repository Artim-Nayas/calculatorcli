package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubtraction(t *testing.T) {
	testSamples := map[float64][]map[float64]float64{
		0.0:  {map[float64]float64{0.1: -0.1, 0.2: -0.2}},
		42.0: {map[float64]float64{-1.0: 43.0, -2.0: 44.0}},
		10.0: {map[float64]float64{-0.0001: 10.0001, 12.1234: -2.1234}},
	}
	for initialValue, testSample := range testSamples {
		t.Run(fmt.Sprintf("when initial value is %f", initialValue), func(t *testing.T) {
			for _, testCase := range testSample {
				for subValue, expectedValue := range testCase {
					t.Run(fmt.Sprintf("upon adding %f", subValue), func(t *testing.T) {
						calc := calculator{value: initialValue}
						calc.Sub(subValue)
						assert.Equal(t, expectedValue, calc.value)
					})
				}
			}
		})
	}
}
