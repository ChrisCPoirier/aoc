package day3

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day3",
	Long:  `day3`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func execute() {
	b, err := os.ReadFile(`data/day3-1.txt`)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score: %d", totalScore(string(b), aoc))
	logrus.Infof("example2: %d", totalScore2(string(b), aoc2))
}

func aoc(s string) int {
	return score(findMatch(split(s)))
}

func aoc2(s ...string) int {
	return score(findMatch(s...))
}

func score(r rune) int {
	if r < 91 {
		return int(r - 38)
	}
	return int(r - 96)
}

func findMatch(strings ...string) rune {
	counts := []map[rune]int{}

	for i := 0; i <= len(strings)-2; i++ {

		m1 := map[rune]int{}

		for _, r := range strings[i] {
			m1[r] += 1
		}
		counts = append(counts, m1)
	}

	ts := strings[len(strings)-1]
	i := 0

MAIN:
	for i < len(ts) {
		if len(counts) == 0 {
			return 0
		}

		for _, m := range counts {
			if _, exists := m[rune(ts[i])]; !exists {
				i++
				goto MAIN
			}
		}
		return rune(ts[i])
	}
	return 0
}

func split(s string) (string, string) {
	return s[:len(s)/2], s[len(s)/2:]
}

func totalScore(s string, aoc func(string) int) int {
	lines := strings.Split(s, "\n")

	totalScore := 0
	for _, line := range lines {
		totalScore += aoc(line)
	}

	return totalScore
}

func totalScore2(s string, aoc func(...string) int) int {
	lines := strings.Split(s, "\n")

	totalScore := 0
	i := 0
	for i < len(lines) {

		totalScore += aoc2(lines[i : i+3]...)
		i += 3
	}

	return totalScore
}
