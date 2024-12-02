package day12

import (
	"aoc/cmd/common"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day12",
	Short: "day12",
	Long:  `day12`,
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
	var score int = 0
	// m := common.AsStringMatrix(s, "")

	springs := []string{}
	records := [][]int{}

	for _, line := range strings.Split(s, "\n") {
		items := strings.Split(line, " ")
		springs = append(springs, items[0])
		records = append(records, common.AsInts(strings.Split(items[1], ",")))
	}

	for i, spring := range springs {
		score += arrangements(spring, records[i])
	}

	fmt.Printf("%#v", springs)
	fmt.Printf("%#v", records)

	return int64(score)
}
func part2(s string) int64 {
	var score int = 0
	// m := common.AsStringMatrix(s, "")

	springs := []string{}
	records := [][]int{}

	for _, line := range strings.Split(s, "\n") {
		items := strings.Split(line, " ")
		springs = append(springs, items[0]+`?`+items[0]+`?`+items[0]+`?`+items[0]+`?`+items[0])
		ints := common.AsInts(strings.Split(items[1], ","))
		intsUnfolded := []int{}
		for i := 0; i < 5; i++ {
			intsUnfolded = append(intsUnfolded, ints...)
		}
		records = append(records, intsUnfolded)
	}

	//unfold

	for i, spring := range springs {
		score += arrangements(spring, records[i])
		fmt.Printf("score at %d: %d\n", i, score)
	}

	return int64(score)
}

var cache = map[string]int{}

func arrangements(s string, record []int) int {
	var totalCount int

	if len(record) == 0 {
		if strings.Contains(s, `#`) {
			return 0
		}
		return 1
	}

	key := fmt.Sprintf("%s,%d", s, record)

	if v, ok := cache[key]; ok {
		return v
	}

	firstSpring := strings.Index(s, `#`)

	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			j = 0
			continue
		}

		if firstSpring != -1 && i-j > firstSpring {
			return totalCount
		}

		if s[i] == '?' && (i == len(s)-1 || s[i+1] == '.' || s[i+1] == '?' || s[i+1] == '#') {
			j++
		}

		if s[i] == '#' {
			j++
		}

		if j == record[0] {
			if i < len(s)-1 && s[i+1] == '#' {
				j -= 1
				continue
			}
			mod := 0
			if i < len(s)-1 && s[i+1] == '?' {
				mod = 1
			}
			totalCount += arrangements(s[i+mod+1:], record[1:])
			cache[key] = totalCount
			i -= j - 1

			for s[i] == '#' {
				if i < len(s)-1 && (s[i+1] == '#' || s[i+1] == '?') && (s[i] == '#') {
					i++
				} else {
					break
				}
			}
			j = 0
		}
	}

	return totalCount
}
