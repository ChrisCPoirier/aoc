package day25

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day25",
	Long:  `day25`,
	Use:   "day25",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

type key struct {
	d       grid.Strings
	Lengths []int
}

func (k key) Heights() []int {
	if len(k.Lengths) == 0 {
		// compute heights
		k.Lengths = make([]int, 5)
		for i := range k.Lengths {
			k.Lengths[i]--
		}

		for _, r := range k.d {
			for c, v := range r {
				if v == `#` {
					k.Lengths[c]++
				}
			}
		}
	}
	return k.Lengths
}

type lock struct {
	d       grid.Strings
	Lengths []int
}

func (k lock) Heights() []int {
	if len(k.Lengths) == 0 {
		// compute heights
		k.Lengths = make([]int, 5)
		for i, r := range k.d {
			for c, v := range r {
				if v == `#` {
					k.Lengths[c] = i
				}
			}
		}
	}
	return k.Lengths
}

func part1(s []byte) int {
	locks := []lock{}
	keys := []key{}

	for _, block := range strings.Split(string(s), "\n\n") {
		g := grid.New(block, ``)
		if g[0][0] == `#` {
			locks = append(locks, lock{d: g})
			continue
		}

		keys = append(keys, key{d: g})
	}

	score := 0

	for _, lock := range locks {
		for _, key := range keys {
			if fit(lock.Heights(), key.Heights()) {
				// println(lock.d.Pretty())
				// println(key.d.Pretty())
				// logrus.Infof("key/lock fit\n%#v\n%#v", lock.Heights(), key.Heights())
				score++
			}
		}
	}

	return score

}

func fit(h1, h2 []int) bool {
	for i, h := range h1 {
		if h+h2[i] > 5 {
			return false
		}
	}
	return true
}

func part2(s []byte) int {
	return part1(s)
}
