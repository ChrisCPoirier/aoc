package day1

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
		expected float64
		input    []byte
		fn       func([]byte) float64
	}{
		{
			expected: 11.0,
			input:    b,
			fn:       part1,
		},
		{
			expected: 31.0,
			input:    b2,
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}
