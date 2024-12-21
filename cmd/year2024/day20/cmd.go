package day20

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"

	"github.com/sirupsen/logrus"
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
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

func part1(s []byte) int {
	g := grid.New(s, ``)

	start := g.FindCell(`S`)
	end := g.FindCell(`E`)

	path := g.BFS(start[0], start[1], end[0], end[1], 0)

	scores := map[int]int{0: len(path.Path)}

	queue := [][][]int{}
	seen := map[string]bool{}

	for r, row := range g {
		for c, v := range row {
			if v == `#` {
				if _, ok := seen[grid.Key(r, c)]; ok {
					continue
				}

				seen[grid.Key(r, c)] = true

				next := [][][]int{}

				for _, dir := range grid.DIR_CROSS {
					nr := r + dir[0]
					nc := r + dir[1]

					if g.InBound(nr, nc) && g[nr][nc] == `#` {
						n := sort([][]int{{r, c}, {nr, nc}})

						if _, ok := seen[grid.Key(n[0][0], n[0][1])]; ok {
							continue
						}

						seen[grid.Key(n[0][0], n[0][1])] = true

						next = append(next, n)
					}
				}

				if len(next) == 0 {
					next = append(next, [][]int{{r, c}})
				}

				queue = append(queue, next...)
			}
		}
	}
	logrus.Infof("queue is size %d", len(queue))

	for i, q := range queue {
		logrus.Infof("processing: %d", i)

		g[q[0][0]][q[0][1]] = `.`
		if len(q) > 1 {
			g[q[1][0]][q[1][1]] = `.`
		}

		np := g.BFS(start[0], start[1], end[0], end[1], scores[0]-100)

		if len(np.Path) < scores[0] {
			scores[scores[0]-len(np.Path)]++
		}

		g[q[0][0]][q[0][1]] = `#`

		if len(q) > 1 {
			g[q[1][0]][q[1][1]] = `#`
		}
	}

	score := 0

	for k, v := range scores {
		if k >= 100 {
			score += v
		}
	}

	logrus.Infof("%#v", scores)
	return score
}

func sort(in [][]int) [][]int {
	if in[1][0] < in[0][0] || (in[0][0] == in[1][0] && in[1][1] < in[1][0]) {
		in[0], in[1] = in[1], in[0]
	}

	return in
}

func part2(s []byte) int {

	return part1(s)
}
