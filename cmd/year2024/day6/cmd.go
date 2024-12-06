package day6

import (
	"aoc/cmd/common"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day6",
	Long:  `day6`,
	Use:   "day6",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
}

func part1(s []byte) int {
	score := 0

	// items := bytes.Split(s, []byte("\n\n"))

	// rm := matrix.New(items[0], "|").Ints()

	return score
}

func part2(s []byte) int {
	return part1(s)
}
