package day11

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		expected int64
		input    string
		scale    int
		fn       func(string, int) int64
	}{
		{
			expected: 374,
			input:    `test.txt`,
			fn:       part1,
			scale:    2,
		},
		{
			expected: 1030,
			input:    `test.txt`,
			fn:       part1,
			scale:    10,
		},
		{
			expected: 8410,
			input:    `test.txt`,
			fn:       part1,
			scale:    100,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b), test.scale))
	}
}
