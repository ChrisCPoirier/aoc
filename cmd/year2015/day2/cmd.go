package day2

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"slices"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day2",
	Long:  `day2`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

func part1(s []byte) int {
	dimensions := grid.New(s, `x`).Ints()

	total := 0
	for _, d := range dimensions {

		a := d[0] * d[1]
		b := d[1] * d[2]
		c := d[0] * d[2]

		slack := a
		for _, v := range []int{a, b, c} {
			if v < slack {
				slack = v
			}
			total += v * 2
		}

		total += slack
	}

	return total
}

func part2(s []byte) int {
	dimensions := grid.New(s, `x`).Ints()

	total := 0
	for _, d := range dimensions {
		total += d[0] * d[1] * d[2]

		ds := slices.Clone(d)
		slices.Sort(ds)

		total += ds[0]*2 + ds[1]*2
	}

	return total
}
