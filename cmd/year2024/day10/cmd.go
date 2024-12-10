package day10

import (
	"aoc/cmd/common"
	"strings"

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
	in := strings.Split(string(s), ``)

	return len(in)
}

func part2(s []byte) int {
	return part1(s)
}
