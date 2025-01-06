package day23

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {

	tests := []struct {
		expected       int
		input          string
		fn             func(string, bool) int
		forceDirection bool
		name           string
	}{
		{
			expected:       94,
			input:          `test.txt`,
			fn:             part1,
			forceDirection: true,
			name:           `part2 wrapped with forcedirection = true`,
		},
		// {
		// 	expected:       154,
		// 	input:          `test.txt`,
		// 	fn:             part2,
		// 	forceDirection: false,
		// },
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b), test.forceDirection), test.name)
	}
}
