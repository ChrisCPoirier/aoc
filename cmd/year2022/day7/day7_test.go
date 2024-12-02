package day7

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectoryWalk(t *testing.T) {

	tests := []struct {
		expected int
		input    Directory
	}{
		{
			expected: 2823,
			input: Directory{
				name: `/`,
				size: 0,
				children: []Node{
					&File{
						name: `fuzzy`,
						size: 1234,
					},
					&Directory{
						name: `b`,
						children: []Node{
							&File{
								name: `wuzzy`,
								size: 1589,
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.input.Size(), test.input)
		assert.Equal(t, 1589, test.input.Directories()[0].Size(), test.input)
	}
}

func TestAoc(t *testing.T) {
	b, err := os.ReadFile(`../../data/day7-test.txt`)
	assert.NoError(t, err)
	root := buildTree(string(b))
	assert.Equal(t, 95437, aoc(root))
	assert.Equal(t, 24933642, aoc2(root))
}
