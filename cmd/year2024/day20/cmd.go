package day20

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day20",
	Long:  `day20`,
	Use:   "day20",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, func(s []byte) int { return part1(s, 2, 100) }, "part 1")
	common.Run(parent, command, 1, func(s []byte) int { return part1(s, 20, 100) }, "part 2")
}

type cheat struct {
	R, C    int
	SR, SC  int // startR, startC
	visited map[string]bool
}

func part1(s []byte, cheatLength, minimum int) int {
	g := grid.New(s, ``)

	start := g.FindCell(`S`)
	end := g.FindCell(`E`)

	lastStep := g.BFS(start[0], start[1], end[0], end[1], 0)
	track := map[string]int{}

	scores := map[int]int{0: len(lastStep.Path)}

	queue := []cheat{}
	for i, step := range lastStep.Path {
		r := step[0]
		c := step[1]
		track[grid.Key(r, c)] = i
		queue = append(queue, cheat{R: r, C: c, SR: r, SC: c, visited: map[string]bool{}})
	}

	var p cheat
	cheats := map[string]int{}
	for len(queue) > 0 {
		p, queue = queue[0], queue[1:]
		if p.visited[grid.Key(p.R, p.C)] {
			continue
		}
		p.visited[grid.Key(p.R, p.C)] = true

		for _, dir := range grid.DIR_CROSS {
			nr := p.R + dir[0]
			nc := p.C + dir[1]

			d := grid.Distance(p.SR, p.SC, nr, nc)

			if d > cheatLength {
				continue
			}

			if !g.InBound(nr, nc) {
				continue
			}

			if _, ok := track[grid.Key(nr, nc)]; ok {
				score := int(math.Abs(float64(track[grid.Key(nr, nc)]-track[grid.Key(p.SR, p.SC)]))) - d
				key := sort([][]int{{nr, nc}, {p.SR, p.SC}})
				cheats[fmt.Sprintf("%#v", key)] = score
			}

			queue = append(queue, cheat{
				R:       nr,
				C:       nc,
				SR:      p.SR,
				SC:      p.SC,
				visited: p.visited,
			})
		}

	}

	score := 0

	for _, v := range cheats {
		scores[v]++
		continue
	}

	for k, v := range scores {
		if k >= minimum {
			score += v
		}
	}
	return score
}

func part2(s []byte, cheatLength, minimum int) int {
	return part1(s, cheatLength, minimum)
}

func sort(in [][]int) [][]int {
	if in[1][0] < in[0][0] || (in[1][0] == in[0][0] && in[1][1] < in[0][1]) {
		in[0], in[1] = in[1], in[0]
	}

	return in
}
