package day5

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"4d63.com/strrev"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day5",
	Short: "day5",
	Long:  `day5`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func execute() {
	b, err := os.ReadFile(`data/day5-1.txt`)

	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("CrateMover 9000: %s", aoc(string(b), makeMoves))
	logrus.Infof("CrateMover 9001: %s", aoc(string(b), makeMovesNoReverse))

}

func aoc(s string, makeMoves func([]string, []move) []string) string {
	return getTop(makeMoves(parse(s)))
}

type move struct {
	from  int
	to    int
	count int
}

// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2

func parse(input string) ([]string, []move) {
	stacks := []string{}
	moves := []move{}

	stacksComplete := false
	for _, line := range strings.Split(input, "\n") {
		if line == `` {
			stacksComplete = true
			continue
		}

		if !stacksComplete {
			//process line as stack
			stacks = parseStack(line, stacks)
			continue
		}

		//process line as move
		moves = parseMove(line, moves)

	}

	return stacks, moves

}

var stacksRE = regexp.MustCompile(`[\[\s]([\w\s])[\]\s]\s`)

func parseStack(s string, stacks []string) []string {
	matches := stacksRE.FindAllStringSubmatch(s+` `, len(s))

	if len(stacks) == 0 {
		stacks = make([]string, len(matches))
	}

	if matches[0][1] == `1` {
		return stacks
	}

	for i, match := range matches {
		if match[1] != ` ` {
			stacks[i] += match[1]
		}

	}

	return stacks
}

var movesRE = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func parseMove(s string, moves []move) []move {
	matches := movesRE.FindStringSubmatch(s)

	count, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])

	return append(moves, move{
		count: count,
		from:  from,
		to:    to,
	})
}

func makeMoves(stacks []string, moves []move) []string {
	for _, move := range moves {
		d := stacks[move.from-1][0:move.count]
		stacks[move.from-1] = stacks[move.from-1][move.count:]
		stacks[move.to-1] = strrev.Reverse(d) + stacks[move.to-1]
	}

	return stacks

}

func makeMovesNoReverse(stacks []string, moves []move) []string {
	for _, move := range moves {
		d := stacks[move.from-1][0:move.count]
		stacks[move.from-1] = stacks[move.from-1][move.count:]
		stacks[move.to-1] = d + stacks[move.to-1]
	}

	return stacks

}

func getTop(stacks []string) string {
	s := ``

	for _, stack := range stacks {
		s += string(stack[0])
	}

	return s
}
