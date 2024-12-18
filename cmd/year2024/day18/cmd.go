package day18

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day18",
	Long:  `day18`,
	Use:   "day18",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

func parse(s []byte) ([][]int, int, int) {
	lines := strings.Split(string(s), "\n")

	items := strings.Split(lines[0], `,`)
	sizeS := items[0]
	size, _ := strconv.Atoi(sizeS)
	runsS := items[1]
	length, _ := strconv.Atoi(runsS)

	corrupt := [][]int{}
	for _, line := range lines[2:] {
		items := strings.Split(line, `,`)
		c, _ := strconv.Atoi(items[0])
		r, _ := strconv.Atoi(items[1])
		corrupt = append(corrupt, []int{r, c})
	}

	return corrupt, size, length

}

type mem struct {
	r    int
	c    int
	path [][]int
}

func part1(s []byte) int {
	corrupted, size, length := parse(s)

	g := make(grid.Strings, size+1)

	for i := range g {
		g[i] = strings.Split(strings.Repeat(`.`, size+1), ``)
	}

	for i, corrupt := range corrupted {
		r := corrupt[0]
		c := corrupt[1]
		if i < length {
			g[r][c] = `#`
		}
	}

	mem := getPath(g, size)

	return len(mem.path) - 1
}

func part2(s []byte) string {
	corrupted, size, length := parse(s)
	g := make(grid.Strings, size+1)

	for i := range g {
		g[i] = strings.Split(strings.Repeat(`.`, size+1), ``)
	}

	for i, corrupt := range corrupted {
		r := corrupt[0]
		c := corrupt[1]

		if i < length {
			g[r][c] = `#`
		}
	}

	mem := getPath(g, size)
	for _, v := range corrupted[length:] {

		r := v[0]
		c := v[1]
		g[r][c] = `#`

		if exist(mem.path, r, c) {
			mem = getPath(g, size)
		}

		if len(mem.path) == 0 {
			return fmt.Sprintf("%d,%d", c, r)
		}
	}

	fmt.Println(g.Pretty())

	return `NONE`
}

func exist(path [][]int, r, c int) bool {
	for _, p := range path {
		if p[0] == r && p[1] == c {
			return true
		}
	}
	return false
}

func getPath(g grid.Strings, size int) mem {

	queue := []mem{{
		r:    0,
		c:    0,
		path: [][]int{{0, 0}},
	}}
	var p mem

	visited := map[string]bool{}
	for len(queue) > 0 {
		p, queue = queue[0], queue[1:]

		if _, ok := visited[grid.Key(p.r, p.c)]; ok {
			continue
		}

		visited[grid.Key(p.r, p.c)] = true

		if p.r == size && p.c == size {
			return p
		}

		for _, dir := range grid.DIR_CROSS {
			nr := p.r + dir[0]
			nc := p.c + dir[1]
			if !g.InBound(nr, nc) {
				continue
			}

			if g[nr][nc] == `#` {
				continue
			}

			queue = append(queue, mem{
				r:    nr,
				c:    nc,
				path: append(slices.Clone(p.path), []int{nr, nc}),
			})
		}
	}
	return mem{}
}
