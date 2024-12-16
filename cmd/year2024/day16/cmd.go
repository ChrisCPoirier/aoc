package day16

import (
	"aoc/cmd/common"
	"aoc/cmd/display"
	"aoc/cmd/grid"
	"container/heap"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day16",
	Long:  `day16`,
	Use:   "day16",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}
func key(r, c int) string {
	return fmt.Sprintf("%d:%d", r, c)
}

func part1(s []byte) int {
	g := grid.New(s, ``)

	start := g.FindCell(`S`)

	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &path{
		row:             start[0],
		column:          start[1],
		rowDirection:    grid.DIR_RIGHT[0],
		columnDirection: grid.DIR_RIGHT[1],
		cost:            0,
	})

	visited := map[string]path{}
	for pq.Len() > 0 {
		p := heap.Pop(pq).(*path)

		if v, ok := visited[key(p.row, p.column)]; ok {
			if v.cost < p.cost {
				continue
			}
		}

		visited[key(p.row, p.column)] = *p

		if g[p.row][p.column] == `E` {
			return p.cost
		}

		for _, dir := range grid.DIR_CROSS {
			if g[p.row+dir[0]][p.column+dir[1]] == `#` {
				continue
			}
			c := cost(p.rowDirection, p.columnDirection, dir)

			heap.Push(pq, &path{
				row:             p.row + dir[0],
				column:          p.column + dir[1],
				rowDirection:    dir[0],
				columnDirection: dir[1],
				cost:            p.cost + c,
				path:            append(slices.Clone(p.path), []int{p.row, p.column}),
			})
		}

	}
	return 0
}

func part2(s []byte) int {
	g := grid.New(s, ``)

	start := g.FindCell(`S`)

	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &path{
		row:             start[0],
		column:          start[1],
		rowDirection:    grid.DIR_RIGHT[0],
		columnDirection: grid.DIR_RIGHT[1],
		cost:            0,
	})

	visited := map[string]path{}
	bestPaths := []path{}
	for pq.Len() > 0 {
		p := heap.Pop(pq).(*path)

		if len(bestPaths) > 0 && p.cost > bestPaths[0].cost {
			continue
		}

		if g[p.row][p.column] == `E` {
			bestPaths = append(bestPaths, *p)
			continue
		}

		if v, ok := visited[key(p.row, p.column)+key(p.rowDirection, p.columnDirection)]; ok {
			if v.cost < p.cost {
				continue
			}
		}

		visited[key(p.row, p.column)+key(p.rowDirection, p.columnDirection)] = *p

		for _, dir := range grid.DIR_CROSS {
			if g[p.row+dir[0]][p.column+dir[1]] == `#` {
				continue
			}
			c := cost(p.rowDirection, p.columnDirection, dir)

			heap.Push(pq, &path{
				row:             p.row + dir[0],
				column:          p.column + dir[1],
				rowDirection:    dir[0],
				columnDirection: dir[1],
				cost:            p.cost + c,
				path:            append(slices.Clone(p.path), []int{p.row, p.column}),
			})
		}

	}
	seats := map[string][]int{}

	for _, path := range bestPaths {
		seats[key(path.row, path.column)] = []int{path.row, path.column}
		for _, pos := range path.path {
			g[pos[0]][pos[1]] = `O`
			seats[key(pos[0], pos[1])] = pos
		}
	}

	return len(seats)
}

func cost(r, c int, dir []int) int {
	if r == dir[0]*-1 && c == dir[1]*-1 {
		return 2000 + 1
	}
	if r != dir[0] || c != dir[1] {
		return 1000 + 1
	}
	return 1
}

func part2Vis(s []byte) int {
	g := grid.New(s, ``)
	d := display.New(g)

	start := g.FindCell(`S`)

	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &path{
		row:             start[0],
		column:          start[1],
		rowDirection:    grid.DIR_RIGHT[0],
		columnDirection: grid.DIR_RIGHT[1],
		cost:            0,
	})

	visited := map[string]path{}
	bestPaths := []path{}
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for pq.Len() > 0 {
			p := heap.Pop(pq).(*path)

			if len(bestPaths) > 0 && p.cost > bestPaths[0].cost {
				continue
			}

			if g[p.row][p.column] == `E` {
				bestPaths = append(bestPaths, *p)
				continue
			}

			if v, ok := visited[key(p.row, p.column)+key(p.rowDirection, p.columnDirection)]; ok {
				if v.cost < p.cost {
					continue
				}
			}

			time.Sleep(time.Millisecond * 10)
			d.ColorCell(p.row, p.column, display.RED)
			visited[key(p.row, p.column)+key(p.rowDirection, p.columnDirection)] = *p

			for _, dir := range grid.DIR_CROSS {
				if g[p.row+dir[0]][p.column+dir[1]] == `#` {
					continue
				}
				c := cost(p.rowDirection, p.columnDirection, dir)

				heap.Push(pq, &path{
					row:             p.row + dir[0],
					column:          p.column + dir[1],
					rowDirection:    dir[0],
					columnDirection: dir[1],
					cost:            p.cost + c,
					path:            append(slices.Clone(p.path), []int{p.row, p.column}),
				})
			}

		}
		wg.Done()
	}()

	go func() {
		for _, path := range bestPaths {
			d.ColorCells(path.path, display.GREEN)
		}
	}()

	d.ShowAndRun()

	wg.Wait()
	seats := map[string][]int{}

	for _, path := range bestPaths {
		seats[key(path.row, path.column)] = []int{path.row, path.column}
		for _, pos := range path.path {
			seats[key(pos[0], pos[1])] = pos
		}
	}

	return len(seats)
}

type path struct {
	row             int
	column          int
	rowDirection    int
	columnDirection int
	cost            int
	index           int
	path            [][]int
}

type PriorityQueue []*path

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	path := x.(*path)
	path.index = n
	*pq = append(*pq, path)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	path := old[n-1]
	old[n-1] = nil
	path.index = -1
	*pq = old[0 : n-1]
	return path
}
