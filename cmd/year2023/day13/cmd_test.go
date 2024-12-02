package day13

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		expected             int64
		input                string
		fn                   func(string, int) int64
		allowedImperfections int
	}{
		// {
		// 	expected:             405,
		// 	input:                `test.txt`,
		// 	fn:                   part1,
		// 	allowedImperfections: 0,
		// },
		{
			expected:             401,
			input:                `test.txt`,
			fn:                   part1,
			allowedImperfections: 1,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b), test.allowedImperfections))
	}
}
