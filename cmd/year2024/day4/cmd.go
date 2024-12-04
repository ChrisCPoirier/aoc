package day4

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day4",
	Long:  `day4`,
	Use:   "day4",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
}

var word = []string{`X`, `M`, `A`, `S`}
var mas = []string{`M`, `A`, `S`}

var UP = []int{0, -1}
var DOWN = []int{0, 1}
var LEFT = []int{-1, 0}
var RIGHT = []int{1, 0}

var ALL = [][][]int{
	{UP},
	{UP, LEFT},
	{UP, RIGHT},
	{DOWN},
	{DOWN, LEFT},
	{DOWN, RIGHT},
	{LEFT},
	{RIGHT},
}

var X = [][][]int{
	{UP, LEFT},
	{UP, RIGHT},
	{DOWN, LEFT},
	{DOWN, RIGHT},
}

func part1(s []byte) float64 {
	score := 0.0

	m := matrix.New(s, ``)

	for i, row := range m {
		for j, c := range row {
			if c == word[0] {
				for _, directions := range ALL {
					if ok, _ := crawl(m, word, 1, i, j, [][]int{}, directions...); ok {
						score++
					}
				}
			}
		}
	}
	return score
}

func crawl(m matrix.Strings, word []string, pos, i, j int, stack [][]int, directions ...[]int) (bool, [][]int) {
	if pos > len(word)-1 {
		return true, stack
	}

	stack = append(stack, []int{i, j})

	ni := i
	nj := j

	for _, direction := range directions {
		ni += direction[0]
		nj += direction[1]
	}

	if ni < 0 || ni >= len(m) {
		return false, stack
	}

	if nj < 0 || nj >= len(m) {
		return false, stack
	}

	if m[ni][nj] == word[pos] {
		return crawl(m, word, pos+1, ni, nj, stack, directions...)
	}

	return false, stack
}

func part2(s []byte) float64 {
	score := 0.0

	m := matrix.New(s, ``)

	masStacks := [][][]int{}
	for i, row := range m {
		for j, c := range row {
			if c == mas[0] {
				for _, directions := range X {
					if ok, stack := crawl(m, mas, 1, i, j, [][]int{}, directions...); ok {
						masStacks = append(masStacks, stack)
					}
				}
			}
		}
	}

	xs := map[string]int{}

	for _, stack := range masStacks {
		aPos := stack[1]
		if _, ok := xs[fmt.Sprintf("%d:%d", aPos[0], aPos[1])]; ok {
			score++
		}
		xs[fmt.Sprintf("%d:%d", aPos[0], aPos[1])]++
	}

	return score
}
