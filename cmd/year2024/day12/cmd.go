package day12

import (
	"aoc/cmd/common"
	"aoc/cmd/display"
	"aoc/cmd/matrix"
	"fmt"
	"image/color"
	"sync"

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
	common.Run(parent, command, 1, vizPart1)
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

var colors = map[string]color.RGBA{
	"A": {255, 99, 71, 255},   // Tomato
	"B": {0, 123, 255, 255},   // Bright Blue
	"C": {50, 205, 50, 255},   // Lime Green
	"D": {255, 165, 0, 255},   // Orange
	"E": {138, 43, 226, 255},  // Blue Violet
	"F": {255, 20, 147, 255},  // Deep Pink
	"G": {64, 224, 208, 255},  // Turquoise
	"H": {255, 255, 0, 255},   // Yellow
	"I": {75, 0, 130, 255},    // Indigo
	"J": {165, 42, 42, 255},   // Brown
	"K": {220, 20, 60, 255},   // Crimson
	"L": {70, 130, 180, 255},  // Steel Blue
	"M": {173, 255, 47, 255},  // Green Yellow
	"N": {255, 182, 193, 255}, // Light Pink
	"O": {0, 250, 154, 255},   // Medium Spring Green
	"P": {123, 104, 238, 255}, // Medium Slate Blue
	"Q": {245, 222, 179, 255}, // Wheat
	"R": {240, 230, 140, 255}, // Khaki
	"S": {32, 178, 170, 255},  // Light Sea Green
	"T": {255, 69, 0, 255},    // Red Orange
	"U": {199, 21, 133, 255},  // Medium Violet Red
	"V": {176, 224, 230, 255}, // Powder Blue
	"W": {218, 165, 32, 255},  // Golden Rod
	"X": {106, 90, 205, 255},  // Slate Blue
	"Y": {144, 238, 144, 255}, // Light Green
	"Z": {245, 245, 220, 255}, // Beige
}

func vizPart1(s []byte) int {
	m := matrix.New(s, ``)

	d := display.New(m)
	score := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// time.Sleep(time.Second * 5)
		for _, group := range groups(m) {
			fr, fc := group[0][0], group[0][1]
			d.ColorCells(group, colors[m[fr][fc]])
			area := len(group)
			perm := perimeter(m, group)
			score += area * perm
		}
		wg.Done()
	}()

	d.ShowAndRun()
	wg.Wait()
	return score
}
