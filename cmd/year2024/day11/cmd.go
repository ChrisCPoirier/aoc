package day11

import (
	"aoc/cmd/common"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
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

type block struct {
	id   int
	size int
}

func part1(s []byte) int {
	m := strings.Split(string(s), ` `)
	logrus.Infof("%#v", m)
	score := 0

	cache := map[string]int{}
	for _, s := range m {
		score += next(s, 0, 24, cache)
	}

	return score
}

func next(s string, step, limit int, cache map[string]int) int {
	if step > limit {
		return 1
	}
	step++

	if v, ok := cache[fmt.Sprintf("%s:%d", s, step)]; ok {
		return v
	}

	if s == `0` {
		n := next(`1`, step, limit, cache)
		cache[fmt.Sprintf("%s:%d", s, step)] = n
		return n
	}

	if len(s)%2 == 0 {
		left := s[0 : len(s)/2]

		right := s[len(s)/2:]

		left = removeLeading(left)
		right = removeLeading(right)

		n := next(left, step, limit, cache) + next(right, step, limit, cache)
		cache[fmt.Sprintf("%s:%d", s, step)] = n
		return n
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		logrus.Error(err)
	}
	v = v * 2024

	n := next(fmt.Sprintf("%d", v), step, limit, cache)
	cache[fmt.Sprintf("%s:%d", s, step)] = n
	return n
}

func removeLeading(s string) string {
	for i, v := range s {
		if v != '0' {
			return s[i:]
		}
	}
	return `0`
}

func part2(s []byte) int {
	m := strings.Split(string(s), ` `)
	logrus.Infof("%#v", m)
	score := 0

	cache := map[string]int{}
	for i, s := range m {
		logrus.Infof("processing %d %s", i, s)
		score += next(s, 0, 74, cache)
	}

	return score
}
