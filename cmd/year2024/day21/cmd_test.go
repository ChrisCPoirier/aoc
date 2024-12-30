package day21

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
		name     string
		expected int
		input    []byte
		fn       func([]byte) int
	}{
		{
			name:     `part1`,
			expected: 126384,
			input:    b,
			fn:       part1,
		},
		{
			name:     `part2`,
			expected: 126384,
			input:    b,
			fn:       func(s []byte) int { return part2(s, 2) },
		},
		{
			name:     `part1 only 0`,
			expected: 0,
			input:    []byte(`0`),
			fn:       part1,
		},
		{
			name:     `part1 only 3`,
			expected: 36,
			input:    b2,
			fn:       part1,
		},
		{
			name:     `part1 only 9`,
			expected: 126,
			input:    []byte(`9`),
			fn:       part1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input), test.name)
	}
}

func Test_cartesian(t *testing.T) {
	tests := []struct {
		name string
		sets [][][]any
		want [][]any
	}{
		{
			name: `2 sets of 2`,
			sets: [][][]any{
				{{`1`, `2`}, {`7`, `8`, `9`}},
				{{`4`, `5`, `6`}, {`11`, `12`, `13`}},
			},
			want: [][]any{
				{`1`, `2`, `4`, `5`, `6`},
				{`7`, `8`, `9`, `4`, `5`, `6`},
				{`1`, `2`, `11`, `12`, `13`},
				{`7`, `8`, `9`, `11`, `12`, `13`},
			},
		},
		{
			name: `2 set of 1 and 1 set of 2`,
			sets: [][][]any{
				{{"^", "^", "^", "A", "v", "A", "<", "<", "A"}},
				{{"v", ">", "v", ">", "A"}, {"v", ">", ">", "v", "A"}, {">", "v", "v", ">", "A"}, {">", "v", ">", "v", "A"}, {">", ">", "v", "v", "A"}},
			},
			want: [][]any{
				{`^`, `^`, `^`, `A`, `v`, `A`, `<`, `<`, `A`, `v`, `>`, `v`, `>`, `A`},
				{`^`, `^`, `^`, `A`, `v`, `A`, `<`, `<`, `A`, `v`, `>`, `>`, `v`, `A`},
				{`^`, `^`, `^`, `A`, `v`, `A`, `<`, `<`, `A`, `>`, `v`, `v`, `>`, `A`},
				{`^`, `^`, `^`, `A`, `v`, `A`, `<`, `<`, `A`, `>`, `v`, `>`, `v`, `A`},
				{`^`, `^`, `^`, `A`, `v`, `A`, `<`, `<`, `A`, `>`, `>`, `v`, `v`, `A`},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cartesian(tt.sets[0], tt.sets[1])
			assert.Equal(t, tt.want, got)
		})
	}
}

// func Test_computeKeyPress(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		start []int
// 		want  int
// 		depth int
// 	}{
// 		{
// 			name:  `press 0 depth 0`,
// 			start: dmp[`<`],
// 			depth: 0,
// 			want:  1,
// 		},
// 		{
// 			name:  `pres 0 depth 1`,
// 			start: dmp[`<`],
// 			depth: 1,
// 			want:  2,
// 		},
// 		{
// 			name:  `pres 0 depth 2`,
// 			start: dmp[`<`],
// 			depth: 2,
// 			want:  8,
// 		},
// 		{
// 			name:  `press 3 depth 0`,
// 			start: dmp[`^`],
// 			depth: 0,
// 			want:  2,
// 		},
// 		{
// 			name:  `press 3 depth 1`,
// 			start: dmp[`^`],
// 			depth: 1,
// 			want:  4,
// 		},
// 		{
// 			name:  `press 3 depth 2`,
// 			start: dmp[`^`],
// 			depth: 2,
// 			want:  12,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := computeKeyPress(tt.start, tt.depth)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }

func Test_computeKeyPresses(t *testing.T) {
	tests := []struct {
		name  string
		in    []string
		want  int
		depth int
	}{
		// {
		// 	name:  `0 depth 0`,
		// 	in:    []string{`<`, `A`},
		// 	depth: 0,
		// 	want:  2,
		// },
		{
			name:  `0 depth 1`,
			in:    []string{`<`, `A`},
			depth: 0,
			want:  4,
		},
		{
			name:  `0 depth 2`,
			in:    []string{`<`, `A`},
			depth: 2,
			want:  18,
		},
		{
			name:  `depth 1 `,
			in:    []string{`<`, `A`, `^`, `A`, `^`, `^`, `>`, `A`, `v`, `v`, `v`, `A`},
			depth: 1,
			want:  28,
		},

		// {
		// 	name:  `9 depth 0`,
		// 	in:    []string{`^`, `^`, `^`, `A`},
		// 	depth: 0,
		// 	want:  4,
		// },
		// {
		// 	name:  `9 depth 1`,
		// 	in:    []string{`^`, `^`, `^`, `A`},
		// 	depth: 1,
		// 	want:  6,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := computeKeyPresses(tt.in, tt.depth)
			assert.Equal(t, tt.want, got)
		})
	}
}
