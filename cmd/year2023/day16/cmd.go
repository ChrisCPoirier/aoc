package day16

import (
	"aoc/cmd/matrix"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day16",
	Short: "day16",
	Long:  `day16`,
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

	g := matrix.New(s, "")

	m := map[string]bool{}
	beam(g, 0, 0, `right`, m)

	return int64(len(m))
}

func part2(s string) int64 {
	var score int = 0

	g := matrix.New(s, "")

	//left
	for x := range g {
		m := map[string]bool{}
		beam(g, x, 0, `right`, m)

		nScore := len(m)
		if nScore > score {
			score = nScore
		}
	}

	//right
	for x := range g {
		m := map[string]bool{}
		beam(g, x, len(g[0])-1, `left`, m)

		nScore := len(m)
		if nScore > score {
			score = nScore
		}
	}

	//up
	for y := range g[0] {
		m := map[string]bool{}
		beam(g, 0, y, `down`, m)

		nScore := len(m)
		if nScore > score {
			score = nScore
		}
	}

	// down
	for y := range g[0] {
		m := map[string]bool{}
		beam(g, len(g)-1, y, `up`, m)

		nScore := len(m)
		if nScore > score {
			score = nScore
		}
	}

	return int64(score)
}

func beam(g matrix.Strings, x, y int, direction string, m map[string]bool) {
	if x < 0 || x > len(g)-1 || y < 0 || y > len(g[0])-1 {
		return
	}

	for {
		if _, ok := m[fmt.Sprintf("%d,%d", x, y)]; !ok {
			m[fmt.Sprintf("%d,%d", x, y)] = false
		}

		c := g[x][y]

		switch direction {
		case `right`:
			if c == `|` {
				if hasSplit(m, x, y) {
					return
				}

				beam(g, x-1, y, "up", m)
				beam(g, x+1, y, "down", m)
				return
			}

			if c == `\` {
				beam(g, x+1, y, "down", m)
				return
			}
			if c == `/` {
				beam(g, x-1, y, "up", m)
				return
			}

			y += 1

		case `left`:
			if c == `|` {
				if hasSplit(m, x, y) {
					return
				}

				beam(g, x-1, y, "up", m)
				beam(g, x+1, y, "down", m)
				return
			}
			if c == `\` {
				beam(g, x-1, y, "up", m)
				return
			}
			if c == `/` {
				beam(g, x+1, y, "down", m)
				return
			}

			y -= 1
		case `up`:
			if c == `-` {
				if hasSplit(m, x, y) {
					return
				}

				beam(g, x, y+1, "right", m)
				beam(g, x, y-1, "left", m)
				return
			}
			if c == `\` {
				beam(g, x, y-1, "left", m)
				return
			}
			if c == `/` {
				beam(g, x, y+1, "right", m)
				return
			}

			x -= 1
		case `down`:
			if c == `-` {
				if hasSplit(m, x, y) {
					return
				}

				beam(g, x, y+1, "right", m)
				beam(g, x, y-1, "left", m)
				return
			}
			if c == `\` {
				beam(g, x, y+1, "right", m)
				return
			}
			if c == `/` {
				beam(g, x, y-1, "left", m)
				return
			}

			x += 1
		default:
			panic(`you messed up`)
		}
		if x < 0 || x > len(g)-1 || y < 0 || y > len(g[0])-1 {
			return
		}
	}
}

func hasSplit(m map[string]bool, x, y int) bool {
	if m[fmt.Sprintf("%d,%d", x, y)] {
		return true
	}

	m[fmt.Sprintf("%d,%d", x, y)] = true
	return false
}
