package gosimplecalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T){

	testCases := []struct{
		name string
		a, b, expectedAdd, expectedSubs, expectedMul, expectedDiv int
	}{
		{
			name: "Negative and Negative",
			a : -5,
			b : -2,
			expectedAdd: -7,
			expectedSubs: -3,
			expectedMul: 10,
			expectedDiv: 2,
		},
		{
			name: "Negative and Positive",
			a : -5,
			b: 5,
			expectedAdd: 0,
			expectedSubs: -10,
			expectedMul: -25,
			expectedDiv: -1,
		},
		{
			name: "Positive and Positive",
			a: 2,
			b: 8,
			expectedAdd: 10,
			expectedSubs: -6,
			expectedMul: 16,
			expectedDiv: 0,
		},
		{
			name: "Positive and Negetive",
			a: 5,
			b: -3,
			expectedAdd: 2,
			expectedSubs: 8,
			expectedMul: -15,
			expectedDiv: -1,
		},
	}

	for _, c := range testCases{
		t.Run("TestAdd - " + c.name, func(t *testing.T){
			result := Add(c.a, c.b)
			assert.Equal(t, c.expectedAdd, result)
		})

		t.Run("TestSubs - " + c.name, func(t *testing.T){
			result := Subtract(c.a, c.b)
			assert.Equal(t, c.expectedSubs, result)
		})

		t.Run("TestMul - " + c.name, func(t *testing.T){
			result := Multiply(c.a, c.b)
			assert.Equal(t, c.expectedMul, result)
		})

		t.Run("TestDiv - " + c.name, func(t *testing.T){
			result := Divide(c.a, c.b)
			assert.Equal(t, c.expectedDiv, result)
		})
	}
}
