package day9

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
			expected: 114,
			input:    `test.txt`,
			fn:       part1,
		},
		{
			expected: 2,
			input:    `test.txt`,
			fn:       part2,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)))
	}
}

func TestNextVal(t *testing.T) {
	tests := []struct {
		expected int64
		input    []int64
	}{
		{
			expected: 18,
			input:    []int64{0, 3, 6, 9, 12, 15},
		},
		{
			expected: 28,
			input:    []int64{1, 3, 6, 10, 15, 21},
		},
		{
			expected: 68,
			input:    []int64{10, 13, 16, 21, 30, 45},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, nextVal(test.input))
	}
}
