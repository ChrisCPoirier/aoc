package day15

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"fmt"
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
		d := dm[mv]
		if !move(g, rr, rc, d) {
			continue
		}

		rr = rr + d[0]
		rc = rc + d[1]
	}

	// fmt.Println(g.Pretty())
	return score(g)
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
			if v == `O` {
				score += 100*r + c
			}
		}
	}
	return score
}

func move(g grid.Strings, r, c int, d []int) bool {
	nr := r + d[0]
	nc := c + d[1]
	v := g[nr][nc]

	if v == `#` {
		return false
	}

	if v == `O` {
		if !move(g, nr, nc, d) {
			return false
		}
	}

	g[r][c], g[nr][nc] = g[nr][nc], g[r][c]

	return true
}

func movable(g grid.Strings, points [][]int, d []int) (bool, [][]int) {
	if len(points) == 0 {
		return true, [][]int{}
	}

	mv := [][]int{}
	for _, p := range points {
		nr := p[0] + d[0]
		nc := p[1] + d[1]
		v := g[nr][nc]

		if v == `#` {
			return false, [][]int{}
		}

		if slices.Contains([]string{`[`, `]`}, v) {
			if d[0] == grid.DIR_LEFT[0] && d[1] == grid.DIR_LEFT[1] ||
				d[0] == grid.DIR_RIGHT[0] && d[1] == grid.DIR_RIGHT[1] {
				mv = append(mv, []int{nr, nc})
			}
			if v == `[` {
				mv = append(mv, []int{nr, nc})
				mv = append(mv, []int{nr + grid.DIR_RIGHT[0], nc + grid.DIR_RIGHT[1]})
				continue
			}
			mv = append(mv, []int{nr, nc})
			mv = append(mv, []int{nr + grid.DIR_LEFT[0], nc + grid.DIR_LEFT[1]})
			continue
		}
	}

	ok, next := movable(g, mv, d)
	if !ok {
		return false, [][]int{}
	}
	return true, append(mv, next...)

}

func uniq(in [][]int) [][]int {
	out := [][]int{}
	u := map[string][]int{}

	for _, r := range in {
		if _, ok := u[fmt.Sprintf("%#v", r)]; ok {
			continue
		}
		u[fmt.Sprintf("%#v", r)] = r
		out = append(out, r)
	}

	return out
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
		d := dm[mv]
		queue := [][]int{{rr, rc}}

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
			continue
		}

		//need so we do not double move an item by it being in the list twice
		// We could also create a uniq set of visited locations for a better solution
		movable = uniq(movable)

		for i := len(movable) - 1; i >= 0; i-- {
			m := movable[i]
			g[m[0]][m[1]], g[m[0]+d[0]][m[1]+d[1]] = g[m[0]+d[0]][m[1]+d[1]], g[m[0]][m[1]]
		}
		movable = [][]int{}

		rr += d[0]
		rc += d[1]
	}

	return score2(g)
}

func score2(g grid.Strings) int {
	score := 0
	for r, row := range g {
		for c, v := range row {
			if v == `[` {
				score += 100*r + c
			}
		}
	}
	return score
}
