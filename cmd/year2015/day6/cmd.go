package day6

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day6",
	Short: "day6",
	Long:  `day6`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

var re = regexp.MustCompile(`(.*) (\d+),(\d+).* (\d+),(\d+)`)

func part1(s []byte) int {

	g := grid.Bools{}.Fill(1000, 1000, false)

	for _, line := range strings.Split(string(s), "\n") {

		if !re.MatchString(line) {
			logrus.Errorf("%s did not match regex", line)
		}

		matches := re.FindStringSubmatch(line)

		instruction := matches[1]
		sc, _ := strconv.Atoi(matches[2])
		sr, _ := strconv.Atoi(matches[3])
		ec, _ := strconv.Atoi(matches[4])
		er, _ := strconv.Atoi(matches[5])

		switch instruction {
		case `turn on`:
			g = on(g, sr, sc, er, ec)
		case `turn off`:
			g = off(g, sr, sc, er, ec)
		default:
			g = toggle(g, sr, sc, er, ec)
		}
	}

	return score(g)
}

func part2(s []byte) int {

	g := grid.Ints{}.Fill(1000, 1000, 0)

	for _, line := range strings.Split(string(s), "\n") {

		if !re.MatchString(line) {
			logrus.Errorf("%s did not match regex", line)
		}

		matches := re.FindStringSubmatch(line)

		instruction := matches[1]
		sc, _ := strconv.Atoi(matches[2])
		sr, _ := strconv.Atoi(matches[3])
		ec, _ := strconv.Atoi(matches[4])
		er, _ := strconv.Atoi(matches[5])

		switch instruction {
		case `turn on`:
			g = onV2(g, sr, sc, er, ec)
		case `turn off`:
			g = offV2(g, sr, sc, er, ec)
		default:
			g = toggleV2(g, sr, sc, er, ec)
		}
	}

	return score2(g)
}

func score(g grid.Bools) int {
	score := 0
	for _, row := range g {
		for _, lit := range row {
			if lit {
				score++
			}
		}
	}
	return score
}

func score2(g grid.Ints) int {
	score := 0
	for _, row := range g {
		for _, brigntess := range row {
			score += brigntess
		}
	}
	return score
}

func toggle(g grid.Bools, sr, sc, er, ec int) grid.Bools {
	for i := sr; i <= er; i++ {
		for j := sc; j <= ec; j++ {
			g[i][j] = !g[i][j]
		}
	}
	return g
}

func off(g grid.Bools, sr, sc, er, ec int) grid.Bools {
	for i := sr; i <= er; i++ {
		for j := sc; j <= ec; j++ {
			g[i][j] = false
		}
	}
	return g
}

func on(g grid.Bools, sr, sc, er, ec int) grid.Bools {
	for i := sr; i <= er; i++ {
		for j := sc; j <= ec; j++ {
			g[i][j] = true
		}
	}
	return g
}

func toggleV2(g grid.Ints, sr, sc, er, ec int) grid.Ints {
	for i := sr; i <= er; i++ {
		for j := sc; j <= ec; j++ {
			g[i][j] = g[i][j] + 2
		}
	}
	return g
}

func offV2(g grid.Ints, sr, sc, er, ec int) grid.Ints {
	for i := sr; i <= er; i++ {
		for j := sc; j <= ec; j++ {
			if g[i][j] == 0 {
				continue
			}

			g[i][j] = g[i][j] - 1
		}
	}
	return g
}

func onV2(g grid.Ints, sr, sc, er, ec int) grid.Ints {
	for i := sr; i <= er; i++ {
		for j := sc; j <= ec; j++ {
			g[i][j] = g[i][j] + 1
		}
	}
	return g
}
