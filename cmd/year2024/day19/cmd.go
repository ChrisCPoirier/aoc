package day19

import (
	"aoc/cmd/common"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day19",
	Long:  `day19`,
	Use:   "day19",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

func part1(s []byte) int {
	return len(s)
}

func part2(s []byte) int {
	return part1(s)
}
