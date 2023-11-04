package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {

	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 4,
			input:    `A X`,
		},
		{
			expected: 8,
			input:    `A Y`,
		},
		{
			expected: 3,
			input:    `A Z`,
		},
		{
			expected: 1,
			input:    `B X`,
		},
		{
			expected: 5,
			input:    `B Y`,
		},
		{
			expected: 9,
			input:    `B Z`,
		},
		{
			expected: 7,
			input:    `C X`,
		},
		{
			expected: 2,
			input:    `C Y`,
		},
		{
			expected: 6,
			input:    `C Z`,
		},
		{
			expected: 4,
			input:    `A X`,
		},
		{
			expected: 1,
			input:    `B X`,
		},
		{
			expected: 7,
			input:    `C X`,
		},
		{
			expected: 8,
			input:    `A Y`,
		},
		{
			expected: 5,
			input:    `B Y`,
		},
		{
			expected: 2,
			input:    `C Y`,
		},
		{
			expected: 3,
			input:    `A Z`,
		},
		{
			expected: 9,
			input:    `B Z`,
		},
		{
			expected: 6,
			input:    `C Z`,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, score(test.input), test.input)
	}
}

func TestTotalScore(t *testing.T) {

	tests := []struct {
		expected  int
		input     string
		scoreFunc func(string) int
	}{
		{
			expected:  15,
			input:     "A B\nB X\nC Z",
			scoreFunc: score,
		},
		{
			expected:  12,
			input:     "A B\nB X\nC Z",
			scoreFunc: scoreForced,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, totalScore(test.input, test.scoreFunc), test.input)
	}

}

func TestScoreForced(t *testing.T) {

	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 4,
			input:    `A Y`,
		},
		{
			expected: 1,
			input:    `B X`,
		},
		{
			expected: 2,
			input:    `C X`,
		},
		{
			expected: 7,
			input:    `C Z`,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, scoreForced(test.input), test.input)
	}
}
