package day6

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"errors"
	"fmt"
	"slices"

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
}

func part1(s []byte) int {
	myApp := app.New()
	myWindow := myApp.NewWindow("visualize")

	m := matrix.New(s, "")

	m.Show(myWindow)
	myWindow.ShowAndRun()
	pos := m.FindCell(`^`)
	uniq, _ := getVisited(pos, m)
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

	for _, v := range visited {
		m[v.i][v.j] = `#`
		if _, err := getVisited(slices.Clone(pos), m); err != nil {
			score++
		}
		m[v.i][v.j] = `.`
	}

	return score
}

func getVisited(pos []int, m matrix.Strings) (map[string]loc, error) {
	uniq := map[string]loc{}
	tracer := map[string]loc{}

	uniq[fmt.Sprintf("%d:%d", pos[0], pos[1])] = loc{i: pos[0], j: pos[1]}

	dir := 0
	for {
		pos[0] += directions[dir][0]
		pos[1] += directions[dir][1]

		if !m.InBound(pos[0], pos[1]) {
			break
		}

		if m[pos[0]][pos[1]] != `#` {
			uniq[fmt.Sprintf("%d:%d", pos[0], pos[1])] = loc{i: pos[0], j: pos[1]}
			continue
		}

		if _, ok := tracer[fmt.Sprintf("%d:%d:%d", pos[0], pos[1], dir)]; ok {
			return uniq, errors.New("infinite loop")
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

	return uniq, nil
}
