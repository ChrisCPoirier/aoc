package day11

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {

	b, err := os.ReadFile(`test.txt`)
	assert.NoError(t, err)

	tests := []struct {
		expected string
		input    []byte
		fn       func([]byte) string
	}{
		{
			expected: `ghjaabcc`,
			input:    b,
			fn:       part1,
		},
		// {
		// 	expected: 237746,
		// 	input:    b,
		// 	fn:       part2,
		// },
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}
