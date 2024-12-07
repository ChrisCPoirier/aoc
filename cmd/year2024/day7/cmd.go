package day7

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day7",
	Long:  `day7`,
	Use:   "day7",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
}

func part1(s []byte) int64 {
	var score int64 = 0
	m := matrix.New(s, ": ")

	for _, r := range m {
		t, _ := strconv.Atoi(r[0])
		total := int64(t)
		numbers := common.AsInts(strings.Split(r[1], ` `))

		if getScore(total, 0, numbers, 0, "+", "*") {
			score += total
		}

	}

	return score
}

func getScore(total int64, sum int64, numbers []int, i int, operators ...string) bool {
	if sum > total {
		return false
	}
	if i > len(numbers)-1 {
		return sum == total
	}

	for _, operator := range operators {
		switch operator {
		case "+":
			if getScore(total, sum+int64(numbers[i]), numbers, i+1, operators...) {
				return true
			}
		case "*":
			if i == 0 {
				continue
			}
			if getScore(total, sum*int64(numbers[i]), numbers, i+1, operators...) {
				return true
			}
		case "||":
			if i == 0 {
				continue
			}
			c, _ := strconv.Atoi(fmt.Sprintf("%d%d", sum, numbers[i]))
			if getScore(total, int64(c), numbers, i+1, operators...) {
				return true
			}

		}
	}

	return false
}

func part2(s []byte) int64 {
	var score int64 = 0
	m := matrix.New(s, ": ")

	for _, r := range m {
		t, _ := strconv.Atoi(r[0])
		total := int64(t)
		numbers := common.AsInts(strings.Split(r[1], ` `))

		if getScore(total, 0, numbers, 0, "+", "*", "||") {
			score += total
		}

	}

	return score
}
