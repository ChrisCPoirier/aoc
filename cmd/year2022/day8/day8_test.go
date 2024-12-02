package day8

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAoc(t *testing.T) {
	b, err := os.ReadFile(`../../data/day8-test.txt`)
	assert.NoError(t, err)

	assert.Equal(t, 21, aoc(string(b)))
}
