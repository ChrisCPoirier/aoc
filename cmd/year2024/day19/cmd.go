package day19

import (
	"aoc/cmd/common"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day19",
	Long:  `day19`,
	Use:   "day19",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

func part1(s []byte) int {
	sections := strings.Split(string(s), "\n\n")

	towels := strings.Split(sections[0], `, `)
	m := map[string]bool{}

	for _, towel := range towels {
		m[towel] = true
	}

	patterns := strings.Split(sections[1], "\n")

	score := 0

	for _, pattern := range patterns {

		cache := map[string]int{}
		if find(pattern, m, cache) > 0 {
			score++
		}
	}

	return score
}

func part2(s []byte) int {
	sections := strings.Split(string(s), "\n\n")

	towels := strings.Split(sections[0], `, `)
	m := map[string]bool{}

	for _, towel := range towels {
		m[towel] = true
	}

	patterns := strings.Split(sections[1], "\n")
	score := 0

	for _, pattern := range patterns {
		cache := map[string]int{}
		score += find(pattern, m, cache)
	}

	return score
}

func find(s string, t map[string]bool, c map[string]int) int {
	if len(s) == 0 {
		return 1
	}

	if _, ok := c[s]; ok {
		return c[s]
	}

	c[s] = 0

	score := 0
	for k := range t {
		n := strings.TrimPrefix(s, k)
		score += find(n, t, c)
	}
	c[s] = score

	return score
}

func findByIndex(s string, start int, t map[string]bool, c map[int]int) int {
	if start >= len(s) {
		return 1
	}

	if _, ok := c[start]; ok {
		return c[start]
	}

	c[start] = 0

	score := 0

	for k := range t {
		end := start + len(k)

		if end > len(s) {
			continue
		}

		if !(s[start:end] == k) {
			continue
		}

		score += findByIndex(s, end, t, c)
	}

	c[start] = score

	return score
}
