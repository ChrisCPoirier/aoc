package day12

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day12",
	Long:  `day12`,
	Use:   "day12",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 1, part2, ".2")
}

func part1(s []byte) int {
	m := matrix.New(s, ``)

	score := 0
	for _, group := range groups(m) {
		area := len(group)
		perm := perimeter(m, group)
		score += area * perm
	}

	return score
}

func part2(s []byte) int {
	m := matrix.New(s, ``)

	score := 0
	for _, group := range groups(m) {
		area := len(group)
		sides := sides(m, group)
		score += area * sides
	}

	return score
}

func groups(m matrix.Strings) [][][]int {
	v := map[string]bool{}

	groups := [][][]int{}

	for r, row := range m {
		for c := range row {
			if _, ok := v[key(r, c)]; ok {
				continue
			}

			v[key(r, c)] = true
			group := [][]int{{r, c}}

			for _, dir := range matrix.DIR_CROSS {
				group = append(group, step(m, r, c, dir, v)...)
			}
			groups = append(groups, group)
		}
	}

	return groups
}

func step(m matrix.Strings, r, c int, dir []int, v map[string]bool) [][]int {
	nr := r + dir[0]
	nc := c + dir[1]

	if !m.InBound(nr, nc) {
		return nil
	}

	if m[r][c] != m[nr][nc] {
		return nil
	}

	if _, ok := v[key(nr, nc)]; ok {
		return nil
	}

	v[key(nr, nc)] = true

	steps := [][]int{{nr, nc}}

	for _, dir := range matrix.DIR_CROSS {
		steps = append(steps, step(m, nr, nc, dir, v)...)
	}
	return steps
}

func perimeter(m matrix.Strings, g [][]int) int {
	perm := 0
	for _, p := range g {
		for _, dir := range matrix.DIR_CROSS {
			r := p[0] + dir[0]
			c := p[1] + dir[1]
			if !m.InBound(r, c) || m[p[0]][p[1]] != m[r][c] {
				perm++
				continue
			}
		}
	}
	return perm
}

func sides(m matrix.Strings, g [][]int) int {
	sides := 0
	v := map[string]bool{}
	for _, p := range g {
		r, c := p[0], p[1]
		v[key(r, c)] = true
		for _, dir := range [][]int{matrix.DIR_UP, matrix.DIR_DOWN} {
			if _, ok := v[sKey(r, c, dir)]; ok {
				continue
			}

			cr := p[0] + dir[0]
			cc := p[1] + dir[1]
			if !m.InBound(cr, cc) || m[r][c] != m[cr][cc] {
				sides++
				v[sKey(r, c, dir)] = true
				followEdge(m, v, r, c, matrix.DIR_LEFT, dir)
				followEdge(m, v, r, c, matrix.DIR_RIGHT, dir)
				continue
			}
		}

		for _, dir := range [][]int{matrix.DIR_LEFT, matrix.DIR_RIGHT} {
			if _, ok := v[sKey(r, c, dir)]; ok {
				continue
			}

			cr := p[0] + dir[0]
			cc := p[1] + dir[1]
			if !m.InBound(cr, cc) || m[r][c] != m[cr][cc] {
				sides++
				v[sKey(r, c, dir)] = true
				followEdge(m, v, r, c, matrix.DIR_UP, dir)
				followEdge(m, v, r, c, matrix.DIR_DOWN, dir)
				continue
			}
		}

	}
	return sides
}

func followEdge(m matrix.Strings, v map[string]bool, r, c int, dir []int, check []int) {
	nr := r + dir[0]
	nc := c + dir[1]
	cr := nr + check[0]
	cc := nc + check[1]

	if !m.InBound(nr, nc) {
		return
	}

	if m[r][c] != m[nr][nc] {
		return
	}

	if !m.InBound(cr, cc) || m[nr][nc] != m[cr][cc] {
		v[sKey(nr, nc, check)] = true
		followEdge(m, v, nr, nc, dir, check)
	}
}

func sKey(r, c int, check []int) string {
	return fmt.Sprintf("%s %#v", key(r, c), check)
}

// shorten function name for code read
func key(r, c int) string {
	return matrix.Key(r, c)
}
