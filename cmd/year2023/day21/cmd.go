package day21

import (
	"aoc/cmd/matrix"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var UP = []int{-1, 0}
var DOWN = []int{1, 0}
var LEFT = []int{0, -1}
var RIGHT = []int{0, 1}

var Cmd = &cobra.Command{
	Use:   "day21",
	Short: "day21",
	Long:  `day21`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(string(b), 64, false))
	// logrus.Infof("score part2: %d", part2(string(b)))

}

func part1(s string, steps int, infinite bool) int64 {
	//steps to up, down, left, right
	// m := map[string][]int{}
	// var score int = 0

	g := matrix.New(s, ``)

	start := getStart(g)

	score := walk(g, start, nil, steps, infinite)

	// fmt.Printf("%#v", locations)

	return int64(score)
}

func walk(g matrix.Strings, start []int, genDir []int, max int, infinite bool) int {
	var u, d, l, r []int

	if max < 0 {
		return 0
	}

	locations := [][]int{start}
	for i := 0; i < max; i++ {
		new := [][]int{}
		for _, location := range locations {
			for _, dir := range [][]int{UP, DOWN, LEFT, RIGHT} {
				n := []int{location[0] + dir[0], location[1] + dir[1]}
				if validStep(g, n) {
					new = append(new, n)
				}

				if !infinite {
					continue
				}

				if n[0] > len(g)-1 && !IsOposite(genDir, UP) {
					if d == nil {
						d = append(getStartFromOffset(g, n), i)
						d = append(d, DOWN...)
					}
				} else if n[0] < 0 && !IsOposite(genDir, DOWN) {
					if u == nil {
						u = append(getStartFromOffset(g, n), i)
						u = append(u, UP...)
					}
				} else if n[1] > len(g[0])-1 && !IsOposite(genDir, LEFT) {
					if r == nil {
						r = append(getStartFromOffset(g, n), i)
						r = append(r, RIGHT...)
					}
				} else if n[1] < 0 && !IsOposite(genDir, RIGHT) {
					if l == nil {
						l = append(getStartFromOffset(g, n), i)
						l = append(l, LEFT...)
					}
				}
			}
		}
		locations = uniq(new)
	}

	count := len(locations)

	for _, newWalk := range [][]int{u, d, l, r} {
		if newWalk == nil {
			continue
		}
		count += walk(g, []int{newWalk[0], newWalk[1]}, []int{newWalk[3], newWalk[4]}, max-newWalk[2]-1, infinite)
	}

	return count

}

func IsOposite(dir1 []int, dir2 []int) bool {
	if dir1 == nil {
		return false
	}

	if dir1[0] != 0 && dir2[0] != 0 && dir1[0]-dir2[0] == 0 {
		return true
	}

	if dir1[1] != 0 && dir2[1] != 0 && dir1[1]-dir2[1] == 0 {
		return true
	}
	return false
}

func getStartFromOffset(g matrix.Strings, s []int) []int {
	newRow := s[0]
	newColumn := s[1]

	if newRow >= len(g) {
		newRow = newRow % len(g)
	} else if newRow < 0 {
		newRow = (newRow * -1) % len(g)
		if newRow > 0 {
			newRow = len(g) - newRow
		}
	}

	if newColumn >= len(g[0]) {
		newColumn = newColumn % len(g[0])
	} else if newColumn < 0 {
		newColumn = (newColumn * -1) % len(g[0])
		if newColumn > 0 {
			newColumn = len(g[0]) - newColumn
		}
	}
	return []int{newRow, newColumn}
}

func validStep(g matrix.Strings, s []int) bool {

	row := s[0]
	column := s[1]

	if row > len(g)-1 || row < 0 || column > len(g[0])-1 || column < 0 {
		return false
	}

	return g[row][column] != `#`
}

func getStart(g matrix.Strings) []int {
	for x, row := range g {
		for y, c := range row {
			if c == `S` {
				return []int{x, y}
			}
		}
	}
	return nil
}

func uniq(input [][]int) [][]int {
	m := map[string][]int{}

	for _, in := range input {
		m[fmt.Sprintf("%d", in)] = in
	}

	out := [][]int{}

	for _, v := range m {
		out = append(out, v)
	}

	return out

}
