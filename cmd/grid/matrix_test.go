package grid

import (
	"aoc/cmd/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridRotate(t *testing.T) {

	input := `ABC
EFG
IJK`
	initExpected := Strings{
		{`A`, `B`, `C`},
		{`E`, `F`, `G`},
		{`I`, `J`, `K`},
	}

	g := New(input, ``)

	assert.Equal(t, initExpected, g)

	g.Rotate()

	expected := Strings{
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
		chunks := common.Chunk(test.input, `#`)
		assert.Equal(t, test.expected, chunks)
		assert.Equal(t, test.input, common.Stitch(chunks, `#`))
	}

}

func TestGridRotateUnequal(t *testing.T) {

	input := `AB
EF
IJ
ZY`
	initExpected := Strings{
		{`A`, `B`},
		{`E`, `F`},
		{`I`, `J`},
		{`Z`, `Y`},
	}

	g := New(input, ``)

	assert.Equal(t, initExpected, g)

	g = g.Rotate()

	expected := Strings{
		{`Z`, `I`, `E`, `A`},
		{`Y`, `J`, `F`, `B`},
	}

	assert.Equal(t, expected, g)

}
