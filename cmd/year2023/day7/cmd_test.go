package day7

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
		fn       func(string, bool) int64
	}{
		{
			expected: 6440,
			input:    `test.txt`,
			jokers:   false,
			fn:       part1,
		},
		{
			expected: 5905,
			input:    `test.txt`,
			jokers:   true,
			fn:       part1,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b), test.jokers))
	}
}
