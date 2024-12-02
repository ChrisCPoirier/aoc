package day17

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		expected int64
		input    string
		fn       func(string, int, int) int64
		min, max int
	}{
		{
			expected: 102,
			input:    `test.txt`,
			fn:       part1,
			min:      0,
			max:      3,
		},
		{
			expected: 94,
			input:    `test.txt`,
			fn:       part1,
			min:      4,
			max:      10,
		},
		{
			expected: 71,
			input:    `test2.txt`,
			fn:       part1,
			min:      4,
			max:      10,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b), test.min, test.max))
	}
}
