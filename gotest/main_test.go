package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name     string
	request  []int
	expected []int
}

var deret = deretBilang{5}

func TestFibonacciTable(t *testing.T) {
	tests := []testTable{
		{
			name:     "TestPrimeInput",
			request:  deret.prima(),
			expected: []int{2, 3},
		},
		{
			name:     "TestOddsInput",
			request:  deret.ganjil(),
			expected: []int{1, 3},
		},
		{
			name:     "TestEvensInput",
			request:  deret.genap(),
			expected: []int{2, 4},
		},
		{
			name:     "TestFibonacciInput",
			request:  deret.fibonacci(),
			expected: []int{0, 1, 1, 2, 3, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.request
			assert.Equal(t, test.expected, result)
		})
	}
}
