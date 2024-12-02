package day19

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day19",
	Short: "day19",
	Long:  `day19`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(string(b)))
	logrus.Infof("score part2: %d", part2(string(b)))

}

func part1(s string) int64 {
	// var score int = 0

	sets := strings.Split(s, "\n\n")

	workflows := parseWorkflows(sets[0])
	parts := parseParts(sets[1])

	accepted := []int{}

	for i, part := range parts {
		w := workflows[`in`]

	WORKFLOW:
		for {
			for _, rule := range w.rules {
				if isMatch(rule, part) {
					if rule.destination == `R` {
						break WORKFLOW
					}

					if rule.destination == `A` {
						accepted = append(accepted, i)
						break WORKFLOW
					}

					w = workflows[rule.destination]
					break
				}
			}
		}
	}

	return int64(score(parts, accepted))
}

func score(parts []map[string]int, accepted []int) int {
	score := 0
	for _, i := range accepted {
		for _, val := range parts[i] {
			score += val
		}
	}
	return score
}

func isMatch(r rule, p map[string]int) bool {
	if r.identifier == `` {
		return true
	}

	if r.operator == `>` {
		return p[r.identifier] > r.value
	}

	return p[r.identifier] < r.value
}

type workflow struct {
	name  string
	rules []rule
}

type rule struct {
	identifier  string
	operator    string
	value       int
	destination string
}

func parseParts(s string) []map[string]int {
	parts := []map[string]int{}
	for _, line := range strings.Split(s, "\n") {
		part := map[string]int{}
		for _, rating := range strings.Split(line[1:len(line)-1], ",") {
			iandv := strings.Split(rating, "=")
			val, err := strconv.Atoi(iandv[1])

			if err != nil {
				panic(err)
			}

			part[iandv[0]] = val
		}
		parts = append(parts, part)
	}
	return parts
}

func parseWorkflows(s string) map[string]workflow {
	workflows := map[string]workflow{}

	for _, line := range strings.Split(s, "\n") {

		nai := strings.Split(line, "{")
		name := nai[0]
		w := workflow{name: name}

		for _, rules := range strings.Split(nai[1][0:len(nai[1])-1], ",") {
			items := strings.Split(rules, ":")

			if len(items) == 1 {
				w.rules = append(w.rules, rule{destination: items[0]})
				continue
			}

			identifier := items[0][0:1]
			operator := items[0][1:2]
			value, err := strconv.Atoi(items[0][2:])

			if err != nil {
				panic(err)
			}

			w.rules = append(w.rules,
				rule{
					identifier:  identifier,
					operator:    operator,
					value:       value,
					destination: items[1],
				},
			)
		}
		workflows[name] = w
	}

	return workflows

}

func part2(s string) int64 {
	// var score int = 0

	sets := strings.Split(s, "\n\n")

	workflows := parseWorkflows(sets[0])
	part := map[string][]int{
		`x`: {1, 4000},
		`m`: {1, 4000},
		`a`: {1, 4000},
		`s`: {1, 4000},
	}

	workflows[`A`] = workflow{name: `A`}
	workflows[`R`] = workflow{name: `R`}

	w := workflows[`in`]

	accepted := splitWorkflow(part, w, workflows)

	return int64(score2(accepted))
}

func score2(accepted []map[string][]int) int {
	score := 0

	for _, part := range accepted {
		combinations := 1
		for _, rating := range part {
			combinations *= (rating[1] - rating[0]) + 1
		}
		score += combinations
	}

	return score
}

func splitWorkflow(p map[string][]int, wf workflow, workflows map[string]workflow) []map[string][]int {
	accepted := []map[string][]int{}

	if wf.name == `R` {
		return nil
	}

	if wf.name == `A` {
		return append(accepted, p)
	}

	for _, rule := range wf.rules {

		if rule.identifier == `` {
			return append(accepted, splitWorkflow(p, workflows[rule.destination], workflows)...)
		}

		if rule.operator == `>` {
			if p[rule.identifier][0] > rule.value {
				return append(accepted, splitWorkflow(p, workflows[rule.destination], workflows)...)
			}
			if p[rule.identifier][1] > rule.value {
				newP := copyPart(p)
				newP[rule.identifier][0] = rule.value + 1
				p[rule.identifier][1] = rule.value

				accepted = append(accepted, splitWorkflow(newP, workflows[rule.destination], workflows)...)
			}
			continue
		}
		// <

		if p[rule.identifier][1] < rule.value {
			return append(accepted, splitWorkflow(p, workflows[rule.destination], workflows)...)
		}
		if p[rule.identifier][0] < rule.value {
			newP := copyPart(p)
			newP[rule.identifier][1] = rule.value - 1
			p[rule.identifier][0] = rule.value

			accepted = append(accepted, splitWorkflow(newP, workflows[rule.destination], workflows)...)
		}
		continue

	}

	return accepted
}

func copyPart(part map[string][]int) map[string][]int {
	n := map[string][]int{}
	for k, v := range part {
		n[k] = []int{v[0], v[1]}
	}
	return n
}
