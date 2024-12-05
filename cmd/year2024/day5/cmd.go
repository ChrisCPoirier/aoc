package day5

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"bytes"
	"slices"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day5",
	Long:  `day5`,
	Use:   "day5",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
}

func part1(s []byte) int {
	score := 0

	items := bytes.Split(s, []byte("\n\n"))

	rm := matrix.New(items[0], "|").Ints()

	m := matrix.New(items[1], ",").Ints()

	rules := map[int][]int{}

	for _, r := range rm {
		rules[r[0]] = append(rules[r[0]], r[1])
	}

	for _, r := range m {
		if isValid(r, common.Index(r), rules) {
			mid := (len(r) - 1) / 2
			score += r[mid]
		}

	}

	return score
}

func isValid(r []int, rIndex map[int]int, rules map[int][]int) bool {
	for i, c := range r {
		for _, rule := range rules[c] {
			if n, ok := rIndex[rule]; ok {
				if i > n {
					return false
				}
			}
		}
	}
	return true
}

func part2(s []byte) int {
	score := 0

	items := bytes.Split(s, []byte("\n\n"))

	rm := matrix.New(items[0], "|").Ints()
	m := matrix.New(items[1], ",").Ints()

	rules := map[int][]int{}

	for _, r := range rm {
		rules[r[0]] = append(rules[r[0]], r[1])
	}

	for _, r := range m {
		if !isValid(r, common.Index(r), rules) {
			r = reorder(r, rules)
			mid := (len(r) - 1) / 2
			score += r[mid]
		}
	}

	return score
}

func reorder(r []int, rules map[int][]int) []int {
	out := []int{}

	for _, c := range r {
		rIndex := common.Index(out)
		insert := len(out)

		for _, rule := range rules[c] {
			if n, ok := rIndex[rule]; ok {
				if insert > n {
					insert = n
				}
			}
		}

		out = slices.Insert(out, insert, c)
	}

	return out
}
