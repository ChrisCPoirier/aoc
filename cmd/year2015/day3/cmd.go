package day3

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day3",
	Long:  `day3`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

var dirs = map[string][]int{
	"v": grid.DIR_DOWN,
	"^": grid.DIR_UP,
	">": grid.DIR_RIGHT,
	"<": grid.DIR_LEFT,
}

func part1(s []byte) int {
	r, c := 0, 0
	houses := map[string]int{grid.Key(r, c): 1}

	for _, d := range strings.Split(string(s), ``) {
		r += dirs[d][0]
		c += dirs[d][1]

		houses[grid.Key(r, c)]++
	}

	return len(houses)
}

func part2(s []byte) int {

	//pos[0] = santa
	//pos[1] = robo-santa
	pos := [][]int{{0, 0}, {0, 0}}
	houses := map[string]int{}

	houses[grid.Key(pos[0][0], pos[0][1])]++
	houses[grid.Key(pos[1][0], pos[1][1])]++

	for i, d := range strings.Split(string(s), ``) {
		//bot changes between even and odd numbers
		bot := i % 2
		pos[bot][0] += dirs[d][0]
		pos[bot][1] += dirs[d][1]

		houses[grid.Key(pos[bot][0], pos[bot][1])]++
	}

	return len(houses)
}
