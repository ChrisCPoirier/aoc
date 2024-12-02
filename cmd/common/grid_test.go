package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridRotate(t *testing.T) {

	input := `ABC
EFG
IJK`
	initExpected := Grid{
		{`A`, `B`, `C`},
		{`E`, `F`, `G`},
		{`I`, `J`, `K`},
	}

	g := AsGrid(input, ``)

	assert.Equal(t, initExpected, g)

	g.Rotate()

	expected := Grid{
		{`I`, `E`, `A`},
		{`J`, `F`, `B`},
		{`K`, `G`, `C`},
	}

	assert.Equal(t, expected, g)

	g.Rotate()
	g.Rotate()
	g.Rotate()

	assert.Equal(t, initExpected, g)

}

func TestRandom(t *testing.T) {

	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			input: []string{`.`, `.`, `.`, `0`, `0`, `#`, `.`, `0`, `0`, `#`, `#`, `.`, `.`, `0`},
			expected: [][]string{
				{`.`, `.`, `.`, `0`, `0`},
				{`.`, `0`, `0`},
				{},
				{`.`, `.`, `0`},
			},
		},
		{
			input: []string{`.`, `#`},
			expected: [][]string{
				{`.`},
				{},
			},
		},
	}

	for _, test := range tests {
		chunks := Chunk(test.input, `#`)
		assert.Equal(t, test.expected, chunks)
		assert.Equal(t, test.input, Stitch(chunks, `#`))
	}

}
