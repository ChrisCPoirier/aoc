package day4

import (
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day4",
	Short: "day4",
	Long:  `day4`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func execute() {
	b, err := os.ReadFile(`data/day4-1.txt`)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("part 1: %d", totalScore(string(b), aoc))
	logrus.Infof("part 2: %d", totalScore(string(b), aoc2))
}

func aoc(s string) int {
	e := strings.Split(s, `,`)

	e1 := strings.Split(e[0], `-`)
	e2 := strings.Split(e[1], `-`)

	e1min, _ := strconv.Atoi(e1[0])
	e1max, _ := strconv.Atoi(e1[1])

	e2min, _ := strconv.Atoi(e2[0])
	e2max, _ := strconv.Atoi(e2[1])

	if e1min >= e2min && e1max <= e2max {
		return 1
	}

	if e2min >= e1min && e2max <= e1max {
		return 1
	}

	return 0
}

func aoc2(s string) int {
	e := strings.Split(s, `,`)

	e1 := strings.Split(e[0], `-`)
	e2 := strings.Split(e[1], `-`)

	e1min, _ := strconv.Atoi(e1[0])
	e1max, _ := strconv.Atoi(e1[1])

	e2min, _ := strconv.Atoi(e2[0])
	e2max, _ := strconv.Atoi(e2[1])

	if e1min >= e2min && e1min <= e2max {
		return 1
	}

	if e2min >= e1min && e2min <= e1max {
		return 1
	}

	return 0
}

func totalScore(s string, aoc func(string) int) int {
	lines := strings.Split(s, "\n")

	totalScore := 0
	for _, line := range lines {
		totalScore += aoc(line)
	}

	return totalScore
}
