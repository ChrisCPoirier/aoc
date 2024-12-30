package day24

import (
	"fmt"
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
			expected: 4,
			input:    b,
			fn:       part1,
		},
		{
			expected: 2024,
			input:    b2,
			fn:       part1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}

func TestPart2(t *testing.T) {

	b, err := os.ReadFile(`1.txt`)
	assert.NoError(t, err)

	tests := []struct {
		expected string
		input    []byte
		fn       func([]byte) string
	}{
		{
			expected: `bmn,jss,mvb,rds,wss,z08,z18,z23`,
			input:    b,
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}

func TestAddBinaryStrings(t *testing.T) {
	tests := []struct {
		name  string
		in    []string
		want  int
		wantS string
	}{
		{
			name:  `sample input from tests`,
			in:    []string{`1011`, `1101`},
			want:  24,
			wantS: `11000`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := AddBinaryStrings(tt.in...)
			assert.Equal(t, tt.want, v)
			assert.Equal(t, tt.wantS, fmt.Sprintf("%b", v))
		})
	}
}

func TestAndBinaryStrings(t *testing.T) {
	tests := []struct {
		name  string
		in    []string
		want  int
		wantS string
	}{
		{
			name:  `sample input from tests`,
			in:    []string{`1011`, `1101`},
			want:  9,
			wantS: `1001`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := AndBinaryStrings(tt.in...)
			assert.Equal(t, tt.want, v)
			assert.Equal(t, tt.wantS, fmt.Sprintf("%b", v))
		})
	}
}
