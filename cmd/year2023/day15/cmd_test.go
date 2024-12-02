package day15

import (
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
			expected: 52,
			input:    `HASH`,
			fn:       part1,
		},
		{
			expected: 1320,
			input:    `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
			fn:       part1,
		},
		{
			expected: 0,
			input:    `cm`,
			fn:       part1,
		},
		{
			expected: 3,
			input:    `ot`,
			fn:       part1,
		},
		{
			expected: 145,
			input:    `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
			fn:       part2,
		},
		{
			expected: 89,
			input:    `rn=1,cm-,qp=3,cm=2,qp-,rn-,cm-,rn=1,cm=2,pc=4,ot=9,abc=5,pc-,pc=6,ot=7,pc=7,abc-`,
			fn:       part2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.fn(test.input))
	}
}

// After "ot=7":
// Box 0: [rn 1] [cm 2]
// Box 3: [ot 7] [ab 5] [pc 6]
