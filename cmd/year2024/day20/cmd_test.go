package day20

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
			expected: 44,
			input:    b,
			fn:       func(s []byte) int { return part1(s, 2, 2) },
		},
		{
			expected: 285,
			input:    b,
			fn:       func(s []byte) int { return part2(s, 20, 50) },
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}

func TestSort(t *testing.T) {

	tests := []struct {
		name     string
		expected [][]int
		input    [][]int
	}{
		{
			name:     `already in order, should not change`,
			input:    [][]int{{13, 3}, {13, 5}},
			expected: [][]int{{13, 3}, {13, 5}},
		},
		{
			name:     `out of order, should change`,
			input:    [][]int{{13, 5}, {13, 3}},
			expected: [][]int{{13, 3}, {13, 5}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, sort(test.input), test.name)
	}
}
