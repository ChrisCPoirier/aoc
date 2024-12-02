package day5

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAndMove(t *testing.T) {
	b, err := os.ReadFile(`../../data/day5-test.txt`)
	assert.NoError(t, err)
	tests := []struct {
		expected           []string
		expectedMoves      []move
		expectedAfterMoves []string
		expectedTop        string
		input              string
	}{
		{
			expected:           []string{`NZ`, `DCM`, `P`},
			expectedAfterMoves: []string{`C`, `M`, `ZNDP`},
			expectedTop:        `CMZ`,
			expectedMoves: []move{
				{
					count: 1,
					from:  2,
					to:    1,
				},
				{
					count: 3,
					from:  1,
					to:    3,
				},
				{
					count: 2,
					from:  2,
					to:    1,
				},
				{
					count: 1,
					from:  1,
					to:    2,
				},
			},
			input: string(b),
		},
	}

	for _, test := range tests {
		stacks, moves := parse(test.input)
		assert.Equal(t, test.expected, stacks, test.input)
		assert.Equal(t, test.expectedMoves, moves, test.input)
		moved := makeMoves(stacks, moves)
		assert.Equal(t, test.expectedAfterMoves, moved, test.input)
		assert.Equal(t, test.expectedTop, getTop(moved), test.input)
	}
}
