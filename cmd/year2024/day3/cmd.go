package day3

import (
	"aoc/cmd/common"
	"bytes"
	"regexp"
	"slices"
	"strconv"

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
	common.Run(parent, command, 2, part2)
}

var reMul = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func part1(s []byte) float64 {
	score := 0.0

	for _, submatch := range reMul.FindAllSubmatch(s, len(s)) {

		f1, _ := strconv.ParseFloat(string(submatch[1]), 64)
		f2, _ := strconv.ParseFloat(string(submatch[2]), 64)
		score += f1 * f2

	}

	return score
}

func part2(s []byte) float64 {
	return part1(clean(s))
}

func clean(in []byte) []byte {
	for {
		begin := bytes.Index(in, []byte(`don't()`))
		if begin == -1 {
			break
		}

		end := bytes.Index(in[begin:], []byte(`do()`))
		if end == -1 {
			return in[:begin]
		}
		in = slices.Delete(in, begin, begin+end)
	}
	return in
}
