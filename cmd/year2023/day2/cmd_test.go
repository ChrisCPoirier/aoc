package day2

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 8,
			input:    `test.txt`,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, part1(b))
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 2286,
			input:    `test.txt`,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, part2(b))
	}
}
