package day7

import (
	"aoc/cmd/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day7",
	Short: "day7",
	Long:  `day7`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

type operation struct {
	re *regexp.Regexp
	fn func(map[string]uint16, []string) error
}

var reMap = map[string]operation{
	"assign": {
		re: regexp.MustCompile(`^(\w+) \-\>\s(.*)`),
		fn: assign,
	},
	"NOT": {
		re: regexp.MustCompile(`^NOT (\w+) \-\>\s(.*)`),
		fn: NOT,
	},
	"AND": {
		re: regexp.MustCompile(`^(\w+) AND (\w+) \-\>\s(.*)`),
		fn: AND,
	},
	"OR": {
		re: regexp.MustCompile(`^(\w+) OR (\w+) \-\>\s(.*)`),
		fn: OR,
	},
	"LSHIFT": {
		re: regexp.MustCompile(`^(\w+) LSHIFT (\w+) \-\>\s(.*)`),
		fn: LSHIFT,
	},
	"RSHIFT": {
		re: regexp.MustCompile(`^(\w+) RSHIFT (\w+) \-\>\s(.*)`),
		fn: RSHIFT,
	},
}

func part1(s []byte) int {
	values := map[string]uint16{}

	queue := []string{}
	for _, line := range strings.Split(string(s), "\n") {
		queue = append(queue, line)
	}

	line := ``

	for len(queue) > 0 {

		line, queue = queue[0], queue[1:]

		found := false
		for _, op := range reMap {
			if op.re.MatchString(line) {
				found = true

				matches := op.re.FindStringSubmatch(line)[1:]

				if !ready(values, matches[:len(matches)-1]) {
					queue = append(queue, line)
					break
				}

				err := op.fn(values, matches)

				if err != nil {
					logrus.Errorf("err: %s", err)
				}
			}
		}

		if !found {
			logrus.Fatalf("no matching op for %s", line)
		}
	}

	return int(values[`a`])
}

var allNums = regexp.MustCompile(`^\d+$`)

func ready(m map[string]uint16, s []string) bool {
	for _, v := range s {
		if !hasValue(m, v) {
			return false
		}
	}
	return true
}

func hasValue(v map[string]uint16, s string) bool {
	if allNums.MatchString(s) {
		return true
	}

	if _, ok := v[s]; ok {
		return true
	}

	return false
}

var reBVal = regexp.MustCompile(`\n\d+ -> b`)

func part2(s []byte) int {

	v := part1(s)

	new := reBVal.ReplaceAll(s, []byte(fmt.Sprintf("\n%d -> b", v)))

	return part1(new)
}

func score(m map[string]int) int {
	score := 0
	for _, v := range m {
		score += v
	}
	return score
}

func assign(v map[string]uint16, items []string) error {
	if len(items) > 2 {
		return fmt.Errorf("item count of %d exceeds limit of 1 for assign", len(items))
	}

	v[items[1]] = getValOrLiteral(v, items[0])

	return nil
}

func NOT(v map[string]uint16, items []string) error {
	if len(items) > 2 {
		return fmt.Errorf("item count of %d exceeds limit of 1 for NOT", len(items))
	}

	v[items[1]] = ^getValOrLiteral(v, items[0])

	return nil
}

func AND(v map[string]uint16, items []string) error {
	if len(items) > 3 {
		return fmt.Errorf("item count of %d exceeds limit of 1 for AND", len(items))
	}

	v[items[2]] = getValOrLiteral(v, items[0]) & getValOrLiteral(v, items[1])

	return nil
}

func OR(v map[string]uint16, items []string) error {
	if len(items) > 3 {
		return fmt.Errorf("item count of %d exceeds limit of 1 for OR", len(items))
	}

	v[items[2]] = getValOrLiteral(v, items[0]) | getValOrLiteral(v, items[1])

	return nil
}

func LSHIFT(v map[string]uint16, items []string) error {
	if len(items) > 3 {
		return fmt.Errorf("item count of %d exceeds limit of 1 for OR", len(items))
	}

	v[items[2]] = getValOrLiteral(v, items[0]) << getValOrLiteral(v, items[1])

	return nil
}

func RSHIFT(v map[string]uint16, items []string) error {
	if len(items) > 3 {
		return fmt.Errorf("item count of %d exceeds limit of 1 for OR", len(items))
	}

	v[items[2]] = getValOrLiteral(v, items[0]) >> getValOrLiteral(v, items[1])

	return nil
}

func getValOrLiteral(v map[string]uint16, key string) uint16 {
	if i, err := strconv.Atoi(key); err == nil {
		return uint16(i)
	}

	return v[key]
}
