package day20

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	search = `output`
	tests := []struct {
		expected int64
		input    string
		fn       func(string) int64
		min, max int
	}{
		{
			expected: 32000000,
			input:    `test.txt`,
			fn:       part1,
		},
		{
			expected: 11687500,
			input:    `test2.txt`,
			fn:       part1,
		},
		{
			expected: 1,
			input:    `test2.txt`,
			fn:       part2,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)))
	}
}

func TestModuleQueue(t *testing.T) {
	m := moduleQueue{}
	assert.Len(t, m, 0, `should be 0`)
	m = m.Push(&broadcaster{})
	assert.Len(t, m, 1, `should be 1`)
	m, _ = m.Pop()
	assert.Len(t, m, 0, `should be 0`)
}
