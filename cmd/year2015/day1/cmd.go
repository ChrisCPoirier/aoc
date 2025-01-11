package day1

import (
	"aoc/cmd/common"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1",
	Long:  `day1`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

func part1(s []byte) int {
	score := 0

	for _, v := range strings.Split(string(s), ``) {
		if v == `(` {
			score++
			continue
		}

		score--
	}

	return score
}

func part2(s []byte) int {
	score := 0

	for i, v := range strings.Split(string(s), ``) {
		if v == `(` {
			score++

		} else {
			score--
		}

		if score == -1 {
			return i + 1
		}
	}

	return len(s) + 1
}
