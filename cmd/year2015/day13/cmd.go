package day13

import (
	"aoc/cmd/common"
	"regexp"
	"slices"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day13",
	Short: "day13",
	Long:  `day13`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

type nodes map[string]map[string]int

func (n nodes) Add(n1, n2 string, h int) {
	if _, ok := n[n1]; !ok {
		n[n1] = map[string]int{}
	}

	n[n1][n2] += h
}

var re = regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)`)

type arrangement struct {
	people    []string
	happiness int
}

func part1(s []byte) int {
	n := parse(string(s))

	happiest := findHappiest(n)
	return happiest.happiness
}

func part2(s []byte) int {
	n := parse(string(s))

	for k := range n {
		n.Add(k, `self`, 0)
		n.Add(`self`, k, 0)
	}

	happiest := findHappiest(n)
	return happiest.happiness
}

func findHappiest(n nodes) arrangement {
	queue := []arrangement{}

	for k := range n {
		queue = append(queue, arrangement{people: []string{k}, happiness: 0})
	}

	happiest := arrangement{}
	var a arrangement

	for len(queue) > 0 {
		a, queue = queue[0], queue[1:]

		if len(a.people) == len(n) {
			a.happiness += n[a.people[len(a.people)-1]][a.people[0]]

			if len(happiest.people) == 0 || a.happiness > happiest.happiness {
				happiest = a
			}
			continue
		}

		for name, happiness := range n[a.people[len(a.people)-1]] {
			if slices.Contains(a.people, name) {
				continue
			}
			queue = append(queue, arrangement{people: slices.Concat(a.people, []string{name}), happiness: a.happiness + happiness})
		}
	}

	return happiest
}

func parse(s string) nodes {
	n := nodes{}
	for _, match := range re.FindAllStringSubmatch(string(s), len(s)) {
		happiness, err := strconv.Atoi(match[3])
		if err != nil {
			logrus.Panicf("could not convert %s to int %s", match[3], err)
		}

		if match[2] == `lose` {
			happiness *= -1

		}

		n.Add(match[1], match[4], happiness)
		n.Add(match[4], match[1], happiness)
	}
	return n
}
