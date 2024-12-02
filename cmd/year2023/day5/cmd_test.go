package day5

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		expected int64
		input    string
		fn       func(string) int64
	}{
		{
			expected: 35,
			input:    `test.txt`,
			fn:       part1,
		},
		{
			expected: 46,
			input:    `test.txt`,
			fn:       part2,
		},
		{
			expected: 46,
			input:    `test.txt`,
			fn:       part3,
		},
		// {
		// 	expected: 11554135,
		// 	input:    `1.txt`,
		// 	fn:       part3,
		// },
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)))
	}
}

func TestSortMaps(t *testing.T) {
	tests := []struct {
		expected map[string][][]int64
		input    map[string][][]int64
		fn       func(string) int64
	}{
		{
			input: map[string][][]int64{
				`seeds-to-soil`: {
					{3, 10, 20},
					{4, 0, 10},
				},
				`soil-to-water`: {
					{3, 20, 20},
					{3, 10, 20},
					{4, 10, 10},
				},
			},
			expected: map[string][][]int64{
				`seeds-to-soil`: {
					{4, 0, 10},
					{3, 10, 20},
				},
				`soil-to-water`: {
					{4, 10, 10},
					{3, 10, 20},
					{3, 20, 20},
				},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, sortMaps(test.input))
	}

}

func TestGetDestRanges(t *testing.T) {
	tests := []struct {
		expected [][]int64
		sources  [][]int64
		mapping  [][]int64
	}{
		{
			sources: [][]int64{{0, 9}},
			mapping: [][]int64{
				{4, 0, 10},
				{3, 10, 20},
			},
			expected: [][]int64{{4, 13}},
		},
		{
			sources: [][]int64{{11, 17}},
			mapping: [][]int64{
				{3, 10, 20},
				{4, 0, 10},
			},
			expected: [][]int64{{4, 10}},
		},
		{
			sources: [][]int64{{29, 35}},
			mapping: [][]int64{
				{3, 10, 20},
				{4, 0, 10},
			},
			expected: [][]int64{{22, 22}, {30, 35}},
		},
		{
			sources: [][]int64{{29, 35}},
			mapping: [][]int64{
				{4, 0, 10},
				{3, 10, 20},
			},
			expected: [][]int64{{22, 22}, {30, 35}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, getDestRange(test.sources, test.mapping))
	}

}

func BenchmarkPart1(b *testing.B) {
	in, err := os.ReadFile(`test2.txt`)
	assert.NoError(b, err)

	for n := 0; n < b.N; n++ {
		part1(string(in))
	}
}

func BenchmarkPart2(b *testing.B) {
	in, err := os.ReadFile(`test2.txt`)
	assert.NoError(b, err)

	for n := 0; n < b.N; n++ {
		part2(string(in))
	}
}

func BenchmarkPart3(b *testing.B) {
	in, err := os.ReadFile(`test2.txt`)
	assert.NoError(b, err)

	for n := 0; n < b.N; n++ {
		part3(string(in))
	}
}
