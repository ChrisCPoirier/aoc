package day3

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
			expected: 16,
			input:    `vJrwpWtwJgWrhcsFMMfFFhFp`,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, aoc(test.input), test.input)
	}
}

func TestFindMatch(t *testing.T) {

	tests := []struct {
		expected rune
		input    string
	}{
		{
			expected: 112,
			input:    `vJrwpWtwJgWrhcsFMMfFFhFp`,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, findMatch(split(test.input)), test.input)
	}
}

func TestSplit(t *testing.T) {

	tests := []struct {
		expected []string
		input    string
	}{
		{
			expected: []string{`vJrwpWtwJgWr`, `hcsFMMfFFhFp`},
			input:    `vJrwpWtwJgWrhcsFMMfFFhFp`,
		},
	}

	for _, test := range tests {
		s1, s2 := split(test.input)
		assert.Equal(t, test.expected[0], s1, test.input)
		assert.Equal(t, test.expected[1], s2, test.input)
	}
}

func TestTotalScore(t *testing.T) {

	tests := []struct {
		expected  int
		input     string
		scoreFunc func(string) int
	}{
		{
			expected:  157,
			input:     "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw",
			scoreFunc: aoc,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, totalScore(test.input, test.scoreFunc), test.input)
	}

}

func TestScore(t *testing.T) {

	tests := []struct {
		expected int
		input    rune
	}{
		{
			expected: 16,
			input:    'p',
		},
		{
			expected: 42,
			input:    'P',
		},
		{
			expected: 26,
			input:    'z',
		},
		{
			expected: 52,
			input:    'Z',
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, score(test.input), test.input)
	}
}

func TestTotalScore2(t *testing.T) {

	tests := []struct {
		expected  int
		input     string
		scoreFunc func(...string) int
	}{
		{
			expected:  70,
			input:     "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw",
			scoreFunc: aoc2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, totalScore2(test.input, test.scoreFunc), test.input)
	}

}
