package day22

import (
	"aoc/cmd/common"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day22",
	Long:  `day22`,
	Use:   "day22",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

func part1(s []byte) int {
	score := 0

	lines := strings.Split(string(s), "\n")
	numbers := []int{}

	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		numbers = append(numbers, i)
	}

	for _, num := range numbers {
		for range 2000 {
			num = prune(mix(num, num*64))
			num = prune(mix(num, num/32))
			num = prune(mix(num, num*2048))

		}
		score += num
	}

	return score
}

func mix(s, m int) int {
	return s ^ m
}

func prune(s int) int {
	return s % 16777216
}

func part2(s []byte) int {
	score := 0

	lines := strings.Split(string(s), "\n")
	numbers := []int{}

	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		numbers = append(numbers, i)
	}

	projections := map[int][]int{}
	prices := map[int][]int{}
	for i, num := range numbers {
		for range 2000 {
			cur := num % 10
			num = prune(mix(num, num*64))
			num = prune(mix(num, num/32))
			num = prune(mix(num, num*2048))

			new := num % 10

			projections[i] = append(projections[i], new-cur)
			prices[i] = append(prices[i], new)

		}
		score += num
	}

	uniq := map[string][]int{}
	for _, projection := range projections {
		for i := 0; i < len(projection)-3; i++ {
			uniq[fmt.Sprintf("%#v", projection[i:i+4])] = projection[i : i+4]
		}
	}

	highest := 0

	for _, seq := range uniq {
		score := 0
		for j, prj := range projections {
			for i := 0; i < len(prj)-3; i++ {
				if prj[i] == seq[0] && prj[i+1] == seq[1] &&
					prj[i+2] == seq[2] && prj[i+3] == seq[3] {
					score += prices[j][i+3]
					break
				}
			}
		}
		highest = max(highest, score)
	}

	return highest
}
