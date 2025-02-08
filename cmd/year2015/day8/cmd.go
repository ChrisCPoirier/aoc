package day8

import (
	"aoc/cmd/common"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day8",
	Short: "day8",
	Long:  `day8`,
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
	for _, line := range strings.Split(string(s), "\n") {
		uq, err := strconv.Unquote(line)

		if err != nil {
			logrus.Error(err)
		}

		score += len(line) - len(uq)
	}

	return score
}

func part2(s []byte) int {
	score := 0

	for _, line := range strings.Split(string(s), "\n") {
		uq := strconv.Quote(line)

		score += len(uq) - len(line)
	}

	return score
}
