package day8

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		expected int64
		input    string
		jokers   bool
		fn       func(string) int64
	}{
		{
			expected: 2,
			input:    `test.txt`,
			fn:       part1,
		},
		{
			expected: 6,
			input:    `test2.txt`,
			fn:       part1,
		},
		{
			expected: 6,
			input:    `test3.txt`,
			fn:       part2,
		},
		{
			expected: 6,
			input:    `test3.txt`,
			fn:       part3,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)))
	}
}
