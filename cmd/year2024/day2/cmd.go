package day2

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"math"
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
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
	common.Run(parent, command, 2, bruteForce, `bruteForce`)
}

func part1(s []byte) int {
	score := 0

	g := matrix.New(string(s), " ").Ints()

	for _, row := range g {

		if getFirstUnsafeIndex(row) == -1 {
			score++
		}
	}

	return score
}

func part2(s []byte) int {
	score := 0

	g := matrix.New(string(s), " ").Ints()

	for _, row := range g {
		firstUnsafe := getFirstUnsafeIndex(row)

		if firstUnsafe == -1 {
			score++
			continue
		}

		for _, i := range []int{0, 1, -1} {
			nrow := slices.Clone(row)
			if firstUnsafe+i < 0 || firstUnsafe+i > len(row) {
				continue
			}
			n := slices.Delete(nrow, firstUnsafe+i, firstUnsafe+i+1)

			if getFirstUnsafeIndex(n) == -1 {
				score++
				break
			}
		}
	}

	return score
}

func bruteForce(s []byte) int {
	score := 0

	g := matrix.New(string(s), " ").Ints()

	for _, row := range g {
		for i := 0; i < len(row); i++ {
			nrow := slices.Clone(row)
			nrow = slices.Delete(nrow, i, i+1)
			if getFirstUnsafeIndex(nrow) == -1 {
				score++
				break
			}
		}
	}

	return score
}

func getFirstUnsafeIndex(row []int) int {
	slope := -1
	if row[0]-row[1] < 0 {
		slope = 1
	}

	for i := 1; i < len(row); i++ {
		diff := row[i-1] - row[i]

		if diff == 0 {

			return i - 1
		}

		if math.Abs(float64(diff)) > 3.0 {

			return i - 1
		}

		if diff*slope > 0 {

			return i - 1
		}
	}

	return -1
}
