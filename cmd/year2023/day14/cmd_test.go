package day14

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		expected int64
		input    string
		fn       func(string, int) int64
		cycles   int
	}{
		// {
		// 	expected: 136,
		// 	input:    `test.txt`,
		// 	fn:       part1,
		// 	cycles:   0,
		// },
		// {
		// 	expected: 87,
		// 	input:    `test.txt`,
		// 	fn:       part1,
		// 	cycles:   1,
		// },
		// {
		// 	expected: 69,
		// 	input:    `test.txt`,
		// 	fn:       part1,
		// 	cycles:   2,
		// },
		{
			expected: 64,
			input:    `test.txt`,
			fn:       part1,
			cycles:   1000000000,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b), test.cycles))
	}
}

// .....#....
// ....#...O#
// .....##...
// ..O#......
// .....OOO#.
// .O#...O#.#
// ....O#...O
// .......OOO
// #..OO###..
// #.OOO#...O
