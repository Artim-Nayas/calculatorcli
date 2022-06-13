package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Run("should return calculator with 0 value", func(t *testing.T) {
		calculator := NewCalculator()
		assert.Equal(t, 0.0, calculator.value)
	})

	t.Run("should implement Calculator", func(t *testing.T) {
		calc := NewCalculator()
		assert.Implements(t, new(Calculator), calc)
	})
}

func TestCalculatorFmt(t *testing.T) {
	t.Run("should implement stringer interface", func(t *testing.T) {
		calc := NewCalculator()
		assert.Implements(t, new(fmt.Stringer), calc)
	})

	t.Run("should format with 4 digit precision", func(t *testing.T) {
		testSamples := map[float64]string{
			2.2: "2.2000", 2.223: "2.2230", 2.22453: "2.2245", 2.262789: "2.2628",
		}
		for initValue, expectedFormat := range testSamples {
			t.Run(fmt.Sprintf("given the inital value: %0.4f", initValue), func(t *testing.T) {
				calc := calculator{value: initValue}
				assert.Equal(t, expectedFormat, calc.String())
			})
		}
	})
}
