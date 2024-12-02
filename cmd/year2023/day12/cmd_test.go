package day12

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		expected int64
		input    string
		fn       func(string) int64
	}{
		// {
		// 	expected: 21,
		// 	input:    `test.txt`,
		// 	fn:       part1,
		// },
		{
			expected: 525152,
			input:    `test.txt`,
			fn:       part2,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)))
	}
}

func TestArrangements(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		record   []int
	}{
		{
			expected: 1,
			input:    `???.###`,
			record:   []int{1, 1, 3},
		},
		{
			expected: 4,
			input:    `.??..??...?##.`,
			record:   []int{1, 1, 3},
		},
		{
			expected: 1,
			input:    `?#?#?#?#?#?#?#?`,
			record:   []int{1, 3, 1, 6},
		},

		{
			expected: 1,
			input:    `????.#...#...`,
			record:   []int{4, 1, 1},
		},
		{
			expected: 4,
			input:    `????.######..#####.`,
			record:   []int{1, 6, 5},
		},
		{
			expected: 10,
			input:    `?###????????`,
			record:   []int{3, 2, 1},
		},
		{
			expected: 6,
			input:    `??#???#???#??????`,
			record:   []int{3, 11},
		},
		{
			expected: 2,
			input:    `.???##?.?#?`,
			record:   []int{6, 2},
		},
		{
			expected: 2,
			input:    `??????????#.#?????.?`,
			record:   []int{2, 7, 1, 4, 1},
		},
		{
			expected: 3,
			input:    `.???....#?`,
			record:   []int{1, 1},
		},
		{
			expected: 2,
			input:    `.???...??.`,
			record:   []int{2, 2},
		},
		{
			expected: 19,
			input:    `???.??.????..`,
			record:   []int{1, 2},
		},
		{
			expected: 2,
			input:    `???.?#???#.??#`,
			record:   []int{2, 6, 2},
		},
		{
			expected: 6,
			input:    `???????##???????`,
			record:   []int{5, 5},
		},
	}
	// ???????##??????? 5,5
	// #####??#####????
	// #####?#####?????
	// ?#####?#####????
	// ????#####?#####?
	// ????#####??#####
	// ?????#####?#####

	// #####??##??#####

	for _, test := range tests {

		assert.Equal(t, test.expected, arrangements(test.input, test.record), test.input, test.record)
	}
}

// ????????????? 1,5,1,1
// ???????????
// #?#####?#?# (11 - 1-5-1-1 - (4-1)) = 0
// ???????????? (12 - 1-5-1-1 - (4-1)) = 1
// #?#####?#?#?
// #?#####?#??#
// #?#####??#?#
// ?#?#####?#?#
