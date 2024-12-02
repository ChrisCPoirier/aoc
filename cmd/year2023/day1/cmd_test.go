package day1

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
			expected: 142,
			input:    b,
			fn:       part1,
		},
		{
			expected: 281,
			input:    b2,
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}

func BenchmarkPart1(b *testing.B) {
	in, err := os.ReadFile(`test2.txt`)
	assert.NoError(b, err)

	tData := in
	for i := 0; i <= 10; i++ {
		tData = append(tData, in...)
	}

	for n := 0; n < b.N; n++ {
		part1(tData)
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

func BenchmarkMutate(b *testing.B) {
	in, err := os.ReadFile(`test2.txt`)
	assert.NoError(b, err)

	tData := in
	for i := 0; i <= 10; i++ {
		tData = append(tData, in...)
	}

	for n := 0; n < b.N; n++ {
		part1(mutate(tData))
	}
}
