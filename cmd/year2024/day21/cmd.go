package day21

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"slices"

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
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
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
	num := slices.Clone(numPad)
	sr := 3
	sc := 2

	// input := []string{`0`, `2`, `9`, `A`}
	input := []string{`0`, `2`, `9`, `A`}

	a1sr := 0
	a1sc := 2
	a2sr := 0
	a2sc := 2
	a1string := []string{}
	final := []string{}
	for _, in := range input {
		des := num.FindCell(in)
		p := num.BFS(sr, sc, des[0], des[1], 0)
		sr = p.Path[len(p.Path)-1][0]
		sc = p.Path[len(p.Path)-1][1]

		logrus.Infof("num pad %#v", p.Path)

		a1 := slices.Clone(arrowPad)

		a1Path := [][][]int{}
		a1sr, a1sc, a1Path = sequence(a1, a1sr, a1sc, [][][]int{p.Path})
		logrus.Infof("a1path %#v", a1Path)

		// println(arrowPad.Pretty())
		// newA1 := [][][]int{}
		for _, group := range a1Path {
			p := group[len(group)-1]
			// newA1 = append(newA1, p)
			a1string = append(a1string, arrowPad[p[0]][p[1]])
		}

		a2 := slices.Clone(arrowPad)

		a2Path := [][][]int{}
		logrus.Infof("%#v", a1Path)
		a2sr, a2sc, a2Path = sequence(a2, a2sr, a2sc, a1Path)

		logrus.Infof("a2path %#v", a2Path)

		for _, group := range a2Path {
			p := group[len(group)-1]
			logrus.Infof("group: %#v", group)
			final = append(final, arrowPad[p[0]][p[1]])
		}
	}
	logrus.Infof("%#v", final)
	return len(`a1string`)
}

var dm = map[string]string{
	grid.Key(grid.DIR_LEFT[0], grid.DIR_LEFT[1]):   `<`,
	grid.Key(grid.DIR_RIGHT[0], grid.DIR_RIGHT[1]): `>`,
	grid.Key(grid.DIR_UP[0], grid.DIR_UP[1]):       `^`,
	grid.Key(grid.DIR_DOWN[0], grid.DIR_DOWN[1]):   `v`,
}

func sequence(g grid.Strings, sr, sc int, groups [][][]int) (int, int, [][][]int) {
	n := [][][]int{}
	logrus.Infof("sequence %#v", groups)
	for _, in := range groups {
		for i := 0; i < len(in)-1; i++ {
			dr := in[i+1][0] - in[i][0]
			dc := in[i+1][1] - in[i][1]
			k := dm[grid.Key(dr, dc)]

			d := g.FindCell(k)

			logrus.Infof("sr: %d sc: %d d0: %d d1: %d dr: %d dc: %d %s %#v", sr, sc, d[0], d[1], dr, dc, k, d)

			p := g.BFS(sr, sc, d[0], d[1], 0)
			sr = p.Path[len(p.Path)-1][0]
			sc = p.Path[len(p.Path)-1][1]
			n = append(n, p.Path)
		}
		n = append(n, [][]int{{0, 2}})
	}
	return sr, sc, n
}

func part2(s []byte) int {

	return part1(s)
}
