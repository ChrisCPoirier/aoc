package day21

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day21",
	Long:  `day21`,
	Use:   "day21",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	// common.Run(parent, command, 1, part1, "part 1 not so optimal")
	common.Run(parent, command, 1, func(s []byte) int { return part2(s, 2) }, "part 1 using part 2 code")
	common.Run(parent, command, 1, func(s []byte) int { return part2(s, 25) }, "part 2")
}

var numPad = grid.Strings{
	{`7`, `8`, `9`},
	{`4`, `5`, `6`},
	{`1`, `2`, `3`},
	{`#`, `0`, `A`},
}

var arrowPad = grid.Strings{
	{`#`, `^`, `A`},
	{`<`, `v`, `>`},
}

func part1(s []byte) int {
	score := 0
	for _, code := range strings.Split(string(s), "\n") {
		input := strings.Split(code, ``)

		num := slices.Clone(numPad)
		out := sequence(num, input)
		// cmp := complexity(code, len(out[0]))
		// logrus.Infof("code %s depth %d optimal is %s len: %d complexity: %d %b", code, 0, strings.Join(out[0], ``), len(out[0]), cmp, cmp)

		for range 2 {
			new := [][]string{}
			for _, item := range out {
				ap := slices.Clone(arrowPad)
				new = append(new, sequence(ap, item)...)
			}

			min := 0
			optimal := [][]string{}
			for _, item := range new {
				if min == 0 || len(item) < min {
					min = len(item)
				}
			}

			for _, item := range new {
				if len(item) == min {
					optimal = append(optimal, item)
				}
			}

			out = optimal
			// cmp := complexity(code, len(out[0]))
			// logrus.Infof("code %s depth %d optimal is %s len: %d complexity: %d %b", code, i+1, strings.Join(out[0], ``), len(out[0]), cmp, cmp)
		}

		comp := complexity(code, len(out[0]))
		score += comp
	}
	return score
}

func complexity(code string, size int) int {
	if len(code) < 4 {
		v, _ := strconv.Atoi(code)
		return v * size
	}

	v, err := strconv.Atoi(code[0:3])
	if err != nil {
		logrus.Fatal(err)
	}

	return v * size
}

var dm = map[string]string{
	grid.Key(grid.DIR_LEFT[0], grid.DIR_LEFT[1]):   `<`,
	grid.Key(grid.DIR_RIGHT[0], grid.DIR_RIGHT[1]): `>`,
	grid.Key(grid.DIR_UP[0], grid.DIR_UP[1]):       `^`,
	grid.Key(grid.DIR_DOWN[0], grid.DIR_DOWN[1]):   `v`,
}

var dmp = map[string][]int{
	`<`: {1, 0},
	`>`: {1, 2},
	`^`: {0, 1},
	`v`: {1, 1},
	`A`: {0, 2},
}

func sequence(g grid.Strings, in []string) [][]string {
	outs := [][]string{}

	cell := g.FindCell(`A`)
	sr, sc := cell[0], cell[1]

	for _, r := range in {
		d := g.FindCell(r)

		//if we are already on our destination just hit the activate button again
		if sr == d[0] && sc == d[1] {
			outs = cartesian(outs, [][]string{{`A`}})
			continue
		}

		paths := g.BFSAll(sr, sc, d[0], d[1], 0, false)
		sr, sc = d[0], d[1]

		nouts := [][]string{}
		for _, p := range paths {
			out := []string{}

			for i := 0; i < len(p.Path)-1; i++ {
				dr := p.Path[i+1][0] - p.Path[i][0]
				dc := p.Path[i+1][1] - p.Path[i][1]
				k := dm[grid.Key(dr, dc)]
				out = append(out, k)
			}
			out = append(out, `A`)
			nouts = append(nouts, out)
		}
		// logrus.Infof("outs: %#v", outs)
		// logrus.Infof("bfs result: %#v", nouts)
		outs = cartesian(outs, nouts)
		// logrus.Infof("catesian: %#v", outs)
	}

	return outs
}

func sequence2(g grid.Strings, start, end string) [][]string {
	outs := [][]string{}

	cellS := g.FindCell(start)
	sr, sc := cellS[0], cellS[1]

	cellE := g.FindCell(end)
	er, ec := cellE[0], cellE[1]

	//if we are already on our destination just hit the activate button again
	if sr == er && sc == ec {
		outs = cartesian(outs, [][]string{{`A`}})
		return outs
	}

	paths := g.BFSAll(sr, sc, er, ec, 0, false)

	nouts := [][]string{}
	for _, p := range paths {
		out := []string{}

		for i := 0; i < len(p.Path)-1; i++ {
			dr := p.Path[i+1][0] - p.Path[i][0]
			dc := p.Path[i+1][1] - p.Path[i][1]
			k := dm[grid.Key(dr, dc)]
			out = append(out, k)
		}
		out = append(out, `A`)
		nouts = append(nouts, out)
	}

	outs = cartesian(outs, nouts)
	return outs
}

func cartesian[T any](sets ...[][]T) [][]T {
	if len(sets) == 0 {
		return [][]T{}
	}

	if len(sets) == 1 {
		return sets[0]
	}

	result := sets[0]

	for _, set := range sets[1:] {
		temp := [][]T{}
		if len(result) == 0 {
			result = set
			continue
		}
		for _, element := range set {
			for _, combinations := range result {
				temp = append(temp, append(slices.Clone(combinations), slices.Clone(element)...))
			}
		}

		result = temp
	}
	return result
}

func part2(s []byte, depth int) int {
	score := 0
	for _, code := range strings.Split(string(s), "\n") {
		input := strings.Split(code, ``)

		num := slices.Clone(numPad)
		seqs := sequence(num, input)

		optimal := 0
		for _, seq := range seqs {
			v := computeKeyPresses(seq, depth)

			if optimal == 0 || v < optimal {
				optimal = v
			}
		}

		comp := complexity(code, optimal)
		logrus.Infof("code %s len: %d complexity: %d", code, optimal, comp)
		score += comp
	}
	// logrus.Fatal(`test`)
	return score
}

func computeKeyPress(s, e []int, depth int, cache map[string]int) int {
	optimal := 0
	// optimalSeq := []string{}
	sr, sc := s[0], s[1]
	er, ec := e[0], e[1]

	if v, ok := cache[cacheKey(arrowPad[sr][sc], arrowPad[er][ec], depth)]; ok {
		return v
	}

	if depth == 0 {
		return 2
	}

	if depth == 1 {
		for _, seq := range sequence2(arrowPad, arrowPad[sr][sc], arrowPad[er][ec]) {
			possible := len(seq)
			if optimal == 0 || possible < optimal {
				// optimalSeq = slices.Clone(seq)
				optimal = possible
			}
		}
		// logrus.Infof("optimal seq: depth %d, button %s, %#v", depth, arrowPad[sr][sc], optimalSeq)
		cache[cacheKey(arrowPad[sr][sc], arrowPad[er][ec], depth)] = optimal
		return optimal
	}

	for _, seq := range sequence2(arrowPad, arrowPad[sr][sc], arrowPad[er][ec]) {
		possible := 0
		seq = append([]string{`A`}, seq...)

		for i := 0; i < len(seq)-1; i++ {
			possible += computeKeyPress(dmp[seq[i]], dmp[seq[i+1]], depth-1, cache)
		}

		if optimal == 0 || possible < optimal {
			// optimalSeq = slices.Clone(seq)
			optimal = possible
		}
	}

	cache[cacheKey(arrowPad[sr][sc], arrowPad[er][ec], depth)] = optimal
	return optimal
}

func computeKeyPresses(in []string, depth int) int {
	comp := 0
	in = append([]string{`A`}, in...)
	//start,end,depth
	cache := map[string]int{}

	for i := 0; i < len(in)-1; i++ {
		comp += computeKeyPress(dmp[in[i]], dmp[in[i+1]], depth, cache)
	}

	return comp
}

func cacheKey(start, end string, depth int) string {
	return fmt.Sprintf("%s:%s:%d", start, end, depth)
}
