package day6

import (
	"aoc/cmd/common"
	"aoc/cmd/display"
	"aoc/cmd/grid"
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day6",
	Long:  `day6`,
	Use:   "day6",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
	common.Run(parent, command, 1, visualizePart1)
	// common.Run(parent, command, 2, visualizePart2)

}

func part1(s []byte) int {
	g := grid.New(s, "")
	pos := g.FindCell(`^`)
	visited, _ := getVisited(pos, g)
	uniq := uniq(visited)
	return len(uniq)
}

var directions = [][]int{grid.DIR_UP, grid.DIR_RIGHT, grid.DIR_DOWN, grid.DIR_LEFT}

type loc struct {
	i, j int
}

func part2(s []byte) int {
	score := 0

	g := grid.New(s, "")
	pos := g.FindCell(`^`)

	visited, _ := getVisited(slices.Clone(pos), g)

	for _, v := range uniq(visited) {
		g[v.i][v.j] = `#`
		if _, err := getVisited(slices.Clone(pos), g); err != nil {
			score++
		}
		g[v.i][v.j] = `.`
	}

	return score
}

func uniq(in [][]int) map[string]loc {
	uniq := map[string]loc{}
	for _, r := range in {
		uniq[fmt.Sprintf("%d:%d", r[0], r[1])] = loc{i: r[0], j: r[1]}
	}
	return uniq
}

func getVisited(pos []int, m grid.Strings) ([][]int, error) {
	visited := [][]int{}
	tracer := map[string]loc{}

	visited = append(visited, slices.Clone(pos))

	dir := 0
	for {
		pos[0] += directions[dir][0]
		pos[1] += directions[dir][1]

		if !m.InBound(pos[0], pos[1]) {
			break
		}

		if m[pos[0]][pos[1]] != `#` {
			visited = append(visited, slices.Clone(pos))
			continue
		}

		if _, ok := tracer[fmt.Sprintf("%d:%d:%d", pos[0], pos[1], dir)]; ok {
			return visited, errors.New("infinite loop")
		}

		tracer[fmt.Sprintf("%d:%d:%d", pos[0], pos[1], dir)] = loc{i: pos[0], j: pos[1]}

		pos[0] -= directions[dir][0]
		pos[1] -= directions[dir][1]
		if dir == 3 {
			dir = 0
			continue
		}

		dir++
	}

	return visited, nil
}

func visualizePart1(s []byte) int {
	wg := &sync.WaitGroup{}
	g := grid.New(s, "")
	d := display.New(g)

	pos := g.FindCell(`^`)
	visited, _ := getVisited(pos, g)

	time.Sleep(time.Second * 3)
	go d.ColorCells(visited, display.GREEN)

	d.ShowAndRun()
	wg.Wait()
	return len(uniq(visited))
}

func visualizePart2(s []byte) int {
	score := 0
	g := grid.New(s, "")
	d := display.New(g)

	pos := g.FindCell(`^`)

	visited, _ := getVisited(slices.Clone(pos), g)

	time.Sleep(3 * time.Second)
	go func() {
		for i, v := range uniq(visited) {
			logrus.Infof("uniq visited %s", i)
			g[v.i][v.j] = `#`
			d.ColorCell(v.i, v.j, display.BLUE)
			newPath, err := getVisited(slices.Clone(pos), g)
			if err != nil {
				d.ColorCells(newPath, display.RED)
				score++
			} else {
				d.ColorCells(newPath, display.GREEN)
			}
			time.Sleep(time.Second * 2)
			g[v.i][v.j] = `.`
			d.ColorCell(v.i, v.j, display.BLACK)
			d.ColorCellsNoWait(newPath, display.BLACK)
		}
	}()

	d.ShowAndRun()
	return score
}
