package day11

import (
	"aoc/cmd/common"
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day11",
	Short: "day11",
	Long:  `day11`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

type edge struct {
	cost int
	dest string
}

func part1(s []byte) string {

	san := sanatize(string(s))
	v := strings.Split(san, ``)

	for {
		v = increment(v, len(v)-1)

		if !hasDoubleOverlappingPairs(strings.Join(v, ``)) {
			continue
		}

		if !hasIncreasingStraight(strings.Join(v, ``)) {
			continue
		}

		return strings.Join(v, ``)
	}
}

func increment(v []string, i int) []string {
	if v[i] == `z` {
		v[i] = `a`
		return increment(v, i-1)
	}

	v[i] = string(v[i][0] + 1)

	for slices.Contains(forbiddens, v[i]) {
		v[i] = string(v[i][0] + 1)
	}

	return v
}

var forbiddens = []string{`o`, `i`, `l`}

func sanatize(s string) string {
	m := -1
	for _, forbidden := range forbiddens {
		i := strings.Index(string(s), forbidden)
		if m == -1 && i > -1 || i < m {
			m = i
		}
	}

	if m == -1 {
		return s
	}

	return fmt.Sprintf("%s%s%s", s[0:m], string(s[m]+1), strings.Repeat(`a`, len(s)-1-m))
}

func part2(s []byte) string {
	return part1([]byte(part1(s)))
}

func hasIncreasingStraight(s string) bool {
	prev := rune(s[0])
	count := 1

	for _, r := range s[1:] {
		count++
		if r-prev != 1 {
			count = 1
		}
		if count >= 3 {
			return true
		}
		prev = r
	}
	return false
}

func hasDoubleOverlappingPairs(s string) bool {
	count := 0

	for i := 0; i < len(s)-1; i++ {
		r1, r2 := rune(s[i]), rune(s[i+1])

		if r1-r2 != 0 {
			continue
		}

		count += 1

		if count >= 2 {
			return true
		}

		i++
	}

	return false
}
