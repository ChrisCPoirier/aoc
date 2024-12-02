package day21

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {

	tests := []struct {
		expected int64
		input    string
		fn       func(string, int, bool) int64
		steps    int
		infinite bool
	}{
		// {
		// 	expected: 16,
		// 	input:    `test.txt`,
		// 	fn:       part1,
		// 	steps:    6,
		// 	infinite: false,
		// },
		// {
		// 	expected: 16,
		// 	input:    `test.txt`,
		// 	fn:       part1,
		// 	steps:    6,
		// 	infinite: true,
		// },
		{
			expected: 50,
			input:    `test.txt`,
			fn:       part1,
			steps:    10,
			infinite: true,
		},
		// {
		// 	expected: 1594,
		// 	input:    `test.txt`,
		// 	fn:       part1,
		// 	steps:    50,
		// },
		// {
		// 	expected: 668697,
		// 	input:    `test.txt`,
		// 	fn:       part1,
		// 	steps:    1000,
		// },
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b), test.steps, test.infinite))
	}
}
