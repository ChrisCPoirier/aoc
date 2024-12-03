package day2

import (
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
			expected: 2,
			input:    b,
			fn:       part1,
		},
		{
			expected: 11,
			input:    b2,
			fn:       part2,
		},
		{
			expected: 11,
			input:    b2,
			fn:       bruteForce,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}

func BenchmarkPart2(b *testing.B) {
	in, err := os.ReadFile(`test2.txt`)
	assert.NoError(b, err)

	tData := in
	for i := 0; i <= 10; i++ {
		tData = append(tData, in...)
	}

	for n := 0; n < b.N; n++ {
		part2(tData)
	}
}

func BenchmarkPart3(b *testing.B) {
	in, err := os.ReadFile(`test2.txt`)
	assert.NoError(b, err)

	tData := in
	for i := 0; i <= 10; i++ {
		tData = append(tData, in...)
	}

	for n := 0; n < b.N; n++ {
		bruteForce(tData)
	}
}
