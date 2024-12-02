package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartOfPacketMarker(t *testing.T) {

	tests := []struct {
		expected int
		input    string
	}{
		{
			expected: 7,
			input:    `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
		},
		{
			expected: 5,
			input:    `bvwbjplbgvbhsrlpgdmjqwftvncz`,
		},
		{
			expected: 6,
			input:    `nppdvjthqldpwncqszvftbrmjlhg`,
		},
		{
			expected: 10,
			input:    `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
		},
		{
			expected: 11,
			input:    `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, startOfPacketMarker(test.input), test.input)
	}
}

func TestStartOfMessageMarker(t *testing.T) {

	tests := []struct {
		expected int
		input    []byte
	}{
		{
			expected: 19,
			input:    []byte(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`),
		},
		{
			expected: 23,
			input:    []byte(`bvwbjplbgvbhsrlpgdmjqwftvncz`),
		},
		{
			expected: 23,
			input:    []byte(`nppdvjthqldpwncqszvftbrmjlhg`),
		},
		{
			expected: 29,
			input:    []byte(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`),
		},
		{
			expected: 26,
			input:    []byte(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, startOfMessageMarker(test.input), test.input)
	}
}
