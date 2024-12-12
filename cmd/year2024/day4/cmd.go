package day4

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
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

var XMAS = []string{`X`, `M`, `A`, `S`}
var MAS = []string{`M`, `A`, `S`}

func part1(s []byte) int {
	g := grid.New(s, ``)
	found := find(g, XMAS, grid.DIR_ALL...)
	return len(found)
}

func part2(s []byte) int {
	score := 0
	g := grid.New(s, ``)
	found := find(g, MAS, grid.DIR_X...)

	xs := map[string]int{}

	for _, stack := range found {
		aPos := stack[1]
		if _, ok := xs[fmt.Sprintf("%d:%d", aPos[0], aPos[1])]; ok {
			score++
		}
		xs[fmt.Sprintf("%d:%d", aPos[0], aPos[1])]++
	}

	return score
}

func find(m grid.Strings, word []string, allowedDirections ...[][]int) [][][]int {
	found := [][][]int{}
	for i, row := range m {
		for j, c := range row {
			if c != word[0] {
				continue
			}

			for _, directions := range allowedDirections {
				if ok, location := findNext(m, word, 1, i, j, [][]int{}, directions...); ok {
					found = append(found, location)
				}
			}
		}
	}
	return found
}

func findNext(m grid.Strings, word []string, pos, i, j int, stack [][]int, directions ...[]int) (bool, [][]int) {
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
		return findNext(m, word, pos+1, ni, nj, stack, directions...)
	}

	return false, stack
}
