package day11

import (
	"aoc/cmd/common"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day11",
	Long:  `day11`,
	Use:   "day11",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
}

func part1(s []byte) int {
	m := strings.Split(string(s), ` `)
	score := 0

	cache := map[string]int{}
	for _, s := range m {
		score += next(s, 25, cache)
	}

	return score
}

func part2(s []byte) int {
	m := strings.Split(string(s), ` `)
	score := 0

	cache := map[string]int{}
	for _, s := range m {
		score += next(s, 75, cache)
	}

	return score
}

func Key(s string, r int) string {
	return fmt.Sprintf("%s:%d", s, r)
}

func next(s string, remaining int, c map[string]int) int {
	if remaining == 0 {
		return 1
	}
	remaining--

	if v, ok := c[Key(s, remaining)]; ok {
		return v
	}

	if s == `0` {
		n := next(`1`, remaining, c)
		c[Key(s, remaining)] = n
		return n
	}

	if len(s)%2 == 0 {
		left := Trim(s[0 : len(s)/2])
		right := Trim(s[len(s)/2:])

		n := next(left, remaining, c) + next(right, remaining, c)
		c[Key(s, remaining)] = n
		return n
	}

	v, _ := strconv.Atoi(s)
	v = v * 2024

	n := next(fmt.Sprintf("%d", v), remaining, c)
	c[Key(s, remaining)] = n
	return n
}

func Trim(s string) string {
	for i, v := range s {
		if v != '0' {
			return s[i:]
		}
	}
	return `0`
}
