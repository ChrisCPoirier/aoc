package day6

import (
	"aoc/cmd/matrix"
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
		expected int
		input    []byte
		fn       func([]byte) int
	}{
		{
			expected: 41,
			input:    b,
			fn:       part1,
		},
		{
			expected: 6,
			input:    b2,
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}

func Test_hasLoop(t *testing.T) {
	type args struct {
		pos []int
		m   matrix.Strings
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasLoop(tt.args.pos, tt.args.m); got != tt.want {
				t.Errorf("hasLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
