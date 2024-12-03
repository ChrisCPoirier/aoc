package day3

import (
	"aoc/cmd/common"

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
	common.Run(parent, command, 1, part1)
	// run(parent, command, 2, part2)
}

func part1(s []byte) float64 {
	score := 0.0

	// m := matrix.
	// 	New(string(s), "   ").
	// 	Floats()

	return score
}

func part2(s []byte) float64 {
	score := 0.0

	return score
}
