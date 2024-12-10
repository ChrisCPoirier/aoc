package day10

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day10",
	Long:  `day10`,
	Use:   "day10",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
}

type block struct {
	id   int
	size int
}

func part1(s []byte) int {
	m := matrix.New(s, ``).Ints()
	score := 0
	for _, ends := range trailEnds(trailHeads(m), m) {
		ends = common.Uniq(ends)
		score += len(ends)
	}

	return score
}

var directions = matrix.DIR_CROSS

func part2(s []byte) int {
	m := matrix.New(s, ``).Ints()
	score := 0
	for _, ends := range trailEnds(trailHeads(m), m) {
		// ends = common.Uniq(ends)
		score += len(ends)
	}

	return score
}

func trailHeads(m matrix.Ints) [][]int {
	trailheads := [][]int{}
	for r, row := range m {
		for c, v := range row {
			if v == 0 {
				trailheads = append(trailheads, []int{r, c})
			}
		}
	}
	return trailheads
}

func trailEnds(trailheads [][]int, m matrix.Ints) map[string][][]int {
	trailEnds := map[string][][]int{}
	for _, trailhead := range trailheads {
		ends := [][]int{}
		for _, direction := range matrix.DIR_CROSS {
			ends = append(ends, step(trailhead, direction, m)...)
		}
		trailEnds[fmt.Sprintf("%d:%d", trailhead[0], trailhead[1])] = ends
	}
	return trailEnds
}

func step(pos, dir []int, m matrix.Ints) [][]int {
	nr := pos[0] + dir[0]
	nc := pos[1] + dir[1]

	if !m.InBound(nr, nc) {
		return [][]int{}
	}

	if m[nr][nc]-m[pos[0]][pos[1]] != 1 {
		return [][]int{}
	}

	if m[nr][nc] == 9 {
		return [][]int{{nr, nc}}
	}
	ends := [][]int{}
	for _, direction := range matrix.DIR_CROSS {
		ends = append(ends, step([]int{nr, nc}, direction, m)...)
	}

	return ends

}
