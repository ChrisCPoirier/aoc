package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAoc(t *testing.T) {

	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 0,
			input:    `2-4,6-8`,
		},
		{
			expected: 1,
			input:    `6-6,4-6`,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, aoc(test.input), test.input)
	}
}

func TestAoc2(t *testing.T) {

	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 0,
			input:    `2-4,6-8`,
		},
		{
			expected: 1,
			input:    `6-6,4-6`,
		},
		{
			expected: 1,
			input:    `5-7,7-9`,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, aoc2(test.input), test.input)
	}
}

func TestTotalScore(t *testing.T) {

	tests := []struct {
		expected  int
		input     string
		scoreFunc func(string) int
	}{
		{
			expected: 2,
			input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`,
			scoreFunc: aoc,
		},
		{
			expected: 4,
			input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`,
			scoreFunc: aoc2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, totalScore(test.input, test.scoreFunc), test.input)
	}

}
