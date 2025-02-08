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
		time     int
		fn       func([]byte, int) int
	}{
		{
			expected: 1120,
			time:     1000,
			input:    b,
			fn:       part1,
		},
		{
			expected: 689,
			time:     1000,
			input:    b,
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input, test.time))
	}
}
