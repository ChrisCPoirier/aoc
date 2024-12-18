package day17

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	b, err := os.ReadFile(`test.txt`)
	assert.NoError(t, err)

	tests := []struct {
		expected  string
		input     []byte
		fn        func([]byte) string
		expectMem map[string]int
	}{
		{
			expected: `4,6,3,5,6,3,5,2,1,0`,
			input:    b,
			fn:       part1,
		},
		// {
		// 	expected: 64,
		// 	input:    b,
		// 	fn:       part2,
		// },
		{
			expected: `4,2,5,6,7,7,7,7,3,1,0`,
			input:    []byte("Register A: 2024\nRegister B: 0\nRegister C: 0\n\nProgram: 0,1,5,4,3,0"),
			fn:       part1,
		},
		{
			expected: `0,1,2`,
			input:    []byte("Register A: 10\nRegister B: 0\nRegister C: 0\n\nProgram: 5,0,5,1,5,4"),
			fn:       part1,
		},
		{
			expected:  ``,
			input:     []byte("Register A: 0\nRegister B: 0\nRegister C: 9\n\nProgram: 2,6"),
			fn:        part1,
			expectMem: map[string]int{`B`: 1},
		},
		{
			expected:  ``,
			input:     []byte("Register A: 0\nRegister B: 29\nRegister C: 0\n\nProgram: 1,7"),
			fn:        part1,
			expectMem: map[string]int{`B`: 26},
		},
		{
			expected:  ``,
			input:     []byte("Register A: 0\nRegister B: 2024\nRegister C: 43690\n\nProgram: 4,0"),
			fn:        part1,
			expectMem: map[string]int{`B`: 44354},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
		if test.expectMem != nil {
			for k := range test.expectMem {
				assert.Equal(t, test.expectMem[k], mem[k])
			}
		}
	}
}

func TestPart2(t *testing.T) {

	// b, err := os.ReadFile(`test.txt`)
	// assert.NoError(t, err)

	tests := []struct {
		expected  int
		input     []byte
		fn        func([]byte) int
		expectMem map[string]int
	}{
		{
			expected: 117440,
			input:    []byte("Register A: 2024\nRegister B: 0\nRegister C: 0\n\nProgram: 0,3,5,4,3,0"),
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
		if test.expectMem != nil {
			for k := range test.expectMem {
				assert.Equal(t, test.expectMem[k], mem[k])
			}
		}
	}
}
