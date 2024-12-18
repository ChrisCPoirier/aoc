package day18

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day18",
	Long:  `day18`,
	Use:   "day18",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

func part1(s []byte) int {
	g := grid.New(s, ``)

	return len(g)
}

func part2(s []byte) int {
	return (part1(s))
}
