package day24

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {

	b, err := os.ReadFile(`test.txt`)
	assert.NoError(t, err)

	b2, err := os.ReadFile(`test2.txt`)
	assert.NoError(t, err)

	tests := []struct {
		expected int
		input    []byte
		fn       func([]byte) int
	}{
		{
			expected: 4,
			input:    b,
			fn:       part1,
		},
		{
			expected: 2024,
			input:    b2,
			fn:       part1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}

func TestPart2(t *testing.T) {

	b, err := os.ReadFile(`1.txt`)
	assert.NoError(t, err)

	tests := []struct {
		expected string
		input    []byte
		fn       func([]byte) string
	}{
		{
			expected: `bmn,jss,mvb,rds,wss,z08,z18,z23`,
			input:    b,
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}
