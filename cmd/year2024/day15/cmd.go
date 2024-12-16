package day15

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day15",
	Long:  `day15`,
	Use:   "day15",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

var dm = map[string][]int{
	">": grid.DIR_RIGHT,
	"<": grid.DIR_LEFT,
	"^": grid.DIR_UP,
	"v": grid.DIR_DOWN,
}

func part1(s []byte) int {
	sections := strings.Split(string(s), "\n\n")

	g := grid.New(sections[0], ``)

	moves := []string{}
	for _, line := range strings.Split(sections[1], "\n") {
		moves = append(moves, strings.Split(line, ``)...)
	}

	rb := g.FindCell(`@`)
	rr := rb[0]
	rc := rb[1]

	for _, mv := range moves {
		rr, rc = move(g, mv, rr, rc)
	}

	return score(g)
}

func part2(s []byte) int {
	sections := strings.Split(string(s), "\n\n")

	sections[0] = expand(sections[0])

	g := grid.New(sections[0], ``)

	moves := []string{}
	for _, line := range strings.Split(sections[1], "\n") {
		moves = append(moves, strings.Split(line, ``)...)
	}

	rb := g.FindCell(`@`)
	rr := rb[0]
	rc := rb[1]

	for _, mv := range moves {
		rr, rc = move(g, mv, rr, rc)
	}

	return score(g)
}

func move(g grid.Strings, dir string, r, c int) (int, int) {
	d := dm[dir]
	queue := [][]int{{r, c}}

	var mv []int
	movable := [][]int{}
	canMove := true

	for len(queue) > 0 {
		mv, queue = queue[0], queue[1:]
		movable = append(movable, mv)
		nr := mv[0] + d[0]
		nc := mv[1] + d[1]
		v := g[nr][nc]

		if v == `#` {
			canMove = false
			break
		}

		if v == `O` {
			queue = append(queue, []int{nr, nc})
			continue
		}

		if slices.Contains([]string{`[`, `]`}, v) {
			if d[0] == grid.DIR_LEFT[0] && d[1] == grid.DIR_LEFT[1] ||
				d[0] == grid.DIR_RIGHT[0] && d[1] == grid.DIR_RIGHT[1] {
				queue = append(queue, []int{nr, nc})
				continue
			}
			if v == `[` {
				queue = append(queue, []int{nr, nc})
				queue = append(queue, []int{nr + grid.DIR_RIGHT[0], nc + grid.DIR_RIGHT[1]})
				continue
			}
			queue = append(queue, []int{nr, nc})
			queue = append(queue, []int{nr + grid.DIR_LEFT[0], nc + grid.DIR_LEFT[1]})
			continue
		}
	}

	if !canMove {
		movable = [][]int{}
		return r, c
	}

	//need so we do not double move an item by it being in the list twice
	// We could also create a uniq set of visited locations for a better solution
	movable = common.Uniq(movable)

	for i := len(movable) - 1; i >= 0; i-- {
		m := movable[i]
		g[m[0]][m[1]], g[m[0]+d[0]][m[1]+d[1]] = g[m[0]+d[0]][m[1]+d[1]], g[m[0]][m[1]]
	}
	movable = [][]int{}

	return r + d[0], c + d[1]
}

func expand(s string) string {
	s = strings.ReplaceAll(s, `O`, `[]`)
	s = strings.ReplaceAll(s, `#`, `##`)
	s = strings.ReplaceAll(s, `.`, `..`)
	s = strings.ReplaceAll(s, `@`, `@.`)
	return s
}

func score(g grid.Strings) int {
	score := 0
	for r, row := range g {
		for c, v := range row {
			if v == `O` || v == `[` {
				score += 100*r + c
			}
		}
	}
	return score
}
