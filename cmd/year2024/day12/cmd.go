package day12

import (
	"aoc/cmd/common"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day12",
	Long:  `day12`,
	Use:   "day12",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 1, part2, ".2")
}

func part1(s []byte) int {
	m := strings.Split(string(s), ` `)
	score := len(m)
	return score
}

func part2(s []byte) int {
	return part1(s)
}
