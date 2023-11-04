package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info(`hello world`)
	b, err := os.ReadFile(`test.txt`)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score: %d", totalScore(string(b), score))
	logrus.Infof("scoreForced: %d", totalScore(string(b), scoreForced))
}

var sMap = map[string]int{
	`A`: 1,
	`B`: 2,
	`C`: 3,
	`X`: 1,
	`Y`: 2,
	`Z`: 3,
}

func score(s string) int {
	items := strings.Split(s, ` `)

	if len(items) < 2 {
		return 0
	}

	p1 := sMap[items[0]]
	p2 := sMap[items[1]]

	nScore := p1 - p2

	if nScore == 0 {
		return 3 + p2
	}

	if nScore == -1 || nScore == 2 {
		return p2 + 6
	}

	return p2
}

func totalScore(s string, score func(string) int) int {
	lines := strings.Split(s, "\n")

	totalScore := 0
	for _, line := range lines {
		totalScore += score(line)
	}

	return totalScore
}

func scoreForced(s string) int {
	items := strings.Split(s, ` `)

	if len(items) < 2 {
		return 0
	}

	p1 := sMap[items[0]]
	p2 := p1

	if items[1] == `Z` {
		p2 = p1 + 1
	}
	if items[1] == `X` {
		p2 = p1 + 2
	}
	if items[1] == `Y` {
		p2 = p1
	}

	if p2 > 3 {
		p2 = p2 - 3
	}

	nScore := p1 - p2

	if nScore == 0 {
		return 3 + p2
	}

	if nScore == -1 || nScore == 2 {
		return p2 + 6
	}

	return p2

}
