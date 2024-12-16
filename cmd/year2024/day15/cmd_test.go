package day15

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {

	b, err := os.ReadFile(`test.txt`)
	assert.NoError(t, err)

	// bsmall, err := os.ReadFile(`testsmall.txt`)
	// assert.NoError(t, err)

	tests := []struct {
		expected int
		input    []byte
		fn       func([]byte) int
	}{
		{
			expected: 10092,
			input:    b,
			fn:       part1,
		},
		{
			expected: 9021,
			input:    b,
			fn:       part2,
		},
		// {
		// 	expected: 618,
		// 	input:    bsmall,
		// 	fn:       part2,
		// },
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}
