package day10

import (
	"aoc/cmd/common"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day10",
	Short: "day10",
	Long:  `day10`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	logrus.Infof("Im slow. Come back and make me fast")
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

type edge struct {
	cost int
	dest string
}

func part1(s []byte) int {
	return len(sequence(string(s), 40))
}

func part2(s []byte) int {
	return len(sequence(string(s), 50))
}

func sequence(s string, cnt int) string {
	logrus.Infof("seq: %s", s)
	prev := 0.0
	for i := range cnt {
		out := ``
		n := '0'
		count := 0
		for _, v := range s {
			count++
			if n == '0' {
				n = v
				continue
			}

			if v != n {
				out += fmt.Sprintf("%d%s", count-1, string(n))
				n = v
				count = 1
			}
		}

		out += fmt.Sprintf("%d%s", count, string(n))
		s = out

		logrus.Infof("loop: %d, ratio: %.15f, len: %d, predicted: %.15f", i, float64(len(s))/prev, len(s), prev*1.303577269034)
		prev = float64(len(s))
		// logrus.Infof("seq: %s", s)
	}

	return s
}
