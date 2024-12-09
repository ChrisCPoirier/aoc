package day6

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
	// common.Run(parent, command, 1, visualizePart1)
	// common.Run(parent, command, 2, visualizePart2)

}

func part1(s []byte) int {
	m := matrix.New(s, "")
	pos := m.FindCell(`^`)
	visited, _ := getVisited(pos, m)
	uniq := uniq(visited)
	return len(uniq)
}

var directions = [][]int{matrix.DIR_UP, matrix.DIR_RIGHT, matrix.DIR_DOWN, matrix.DIR_LEFT}

type loc struct {
	i, j int
}

func part2(s []byte) int {
	score := 0

	m := matrix.New(s, "")
	pos := m.FindCell(`^`)

	visited, _ := getVisited(slices.Clone(pos), m)

	for _, v := range uniq(visited) {
		m[v.i][v.j] = `#`
		if _, err := getVisited(slices.Clone(pos), m); err != nil {
			score++
		}
		m[v.i][v.j] = `.`
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

func getVisited(pos []int, m matrix.Strings) ([][]int, error) {
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
	m := matrix.New(s, "")

	myApp := app.New()
	myWindow := myApp.NewWindow("visualize")
	c := m.Fyne(myWindow)

	pos := m.FindCell(`^`)
	wg.Add(1)

	visited, _ := getVisited(pos, m)

	time.Sleep(time.Second * 3)
	go func() {
		for _, visit := range visited {
			time.Sleep(10 * time.Millisecond)
			c.Objects[visit[1]+(len(m[0])*visit[0])] = matrix.NewSquare(m[visit[0]][visit[1]], YELLOW)
		}
		wg.Done()
	}()

	myWindow.ShowAndRun()
	wg.Wait()
	return len(uniq(visited))
}

func visualizePart2(s []byte) int {
	score := 0
	m := matrix.New(s, "")
	myApp := app.New()
	myWindow := myApp.NewWindow("visualize")
	c := m.Fyne(myWindow)

	pos := m.FindCell(`^`)

	visited, _ := getVisited(slices.Clone(pos), m)

	time.Sleep(3 * time.Second)
	go func() {
		for _, v := range uniq(visited) {
			m[v.i][v.j] = `#`
			fillCell(c, m, v.i, v.j, BLUE)
			newPath, err := getVisited(slices.Clone(pos), m)
			if err != nil {
				fill(c, m, newPath, RED)
				score++
			} else {
				fill(c, m, newPath, YELLOW)
			}
			time.Sleep(time.Second * 2)
			m[v.i][v.j] = `.`
			fillNoWait(c, m, newPath, BLACK)
		}
	}()

	myWindow.ShowAndRun()
	return score
}

var BLACK = color.RGBA{0, 0, 0, 100}
var YELLOW = color.RGBA{234, 239, 44, 100}
var RED = color.RGBA{255, 0, 0, 100}
var BLUE = color.RGBA{0, 0, 255, 100}

func reset(cont *fyne.Container, m matrix.Strings, cl color.RGBA) {
	for r, items := range m {
		for c := range items {
			fillCell(cont, m, r, c, cl)
		}
	}
}

func fillCell(cont *fyne.Container, m matrix.Strings, r, c int, cl color.RGBA) {
	cont.Objects[c+(len(m[r])*r)] = matrix.NewSquare(m[r][c], cl)
}

func fill(cont *fyne.Container, m matrix.Strings, cells [][]int, cl color.RGBA) {
	for _, cell := range cells {
		time.Sleep(time.Millisecond * 10)
		fillCell(cont, m, cell[0], cell[1], cl)
	}

}

func fillNoWait(cont *fyne.Container, m matrix.Strings, cells [][]int, cl color.RGBA) {
	for _, cell := range cells {
		fillCell(cont, m, cell[0], cell[1], cl)
	}

}
