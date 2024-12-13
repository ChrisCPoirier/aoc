package day13

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day13",
	Long:  `day13`,
	Use:   "day13",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 1, part2, ".2")
}

func part1(s []byte) int {
	g := grid.New(s, ``)

	return len(g)
}

func part2(s []byte) int {
	g := grid.New(s, ``)

	return len(g)
}
