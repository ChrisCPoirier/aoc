package day7

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
		expected int64
		input    []byte
		fn       func([]byte) int64
	}{
		{
			expected: 3749,
			input:    b,
			fn:       part1,
		},
		{
			expected: 11387,
			input:    b2,
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}

func TestGetMultiScore(t *testing.T) {
	tests := []struct {
		total   int64
		name    string
		numbers []int
	}{
		{
			name:    `1 multiplication`,
			total:   190,
			numbers: []int{10, 19},
		},
		{
			name:    `all but last multiplied`,
			total:   238612,
			numbers: []int{10, 19, 22, 34, 11},
		},
		{
			name:    `multi on left and right side ((10*19)+22)*34*11`,
			total:   79288,
			numbers: []int{10, 19, 22, 34, 11},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.total, getMultiScore(test.total, test.numbers))
	}
}

func TestGetPerm(t *testing.T) {
	tests := []struct {
		expected   [][]bool
		name       string
		mc, length int
	}{
		{
			name:     `1 multiplication`,
			mc:       3,
			length:   4,
			expected: [][]bool{[]bool{true, false, true, true}, []bool{true, true, false, true}, []bool{true, true, true, false}, []bool{false, true, true, true}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, getPermutations(test.mc, test.length))
	}
}
