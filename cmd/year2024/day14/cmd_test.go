package day14

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {

	b, err := os.ReadFile(`test.txt`)
	assert.NoError(t, err)

	tests := []struct {
		expected int
		input    []byte
		fn       func([]byte) int
	}{
		{
			expected: 12,
			input:    b,
			fn:       func(b []byte) int { return part1(b, 11, 7) },
		},
		{
			// Tree does not exist in test input
			expected: 0,
			input:    b,
			fn:       func(b []byte) int { return part2(b, 11, 7) },
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}
