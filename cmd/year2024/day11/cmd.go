package day10

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day11",
	Long:  `day11`,
	Use:   "day11",
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
	score := len(m)
	return score
}

func part2(s []byte) int {
	return part1(s)
}
