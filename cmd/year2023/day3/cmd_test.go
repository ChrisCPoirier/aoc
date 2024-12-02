package day3

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		fn       func(string) int
	}{
		{
			expected: 4361,
			input:    `test.txt`,
			fn:       part1,
		},
		{
			expected: 467835,
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
