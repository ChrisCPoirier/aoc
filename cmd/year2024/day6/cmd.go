package day6

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"fmt"
	"slices"
	"strconv"
	"strings"

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
	m := matrix.New(s, "")

	pos := []int{}
CENTER:
	for i, r := range m {
		for j, c := range r {
			if c == "^" {
				pos = []int{i, j}
				m[i][j] = `.`
				break CENTER
			}
		}
	}

	return len(getVisited(pos, m))
}

func getVisited(pos []int, m matrix.Strings) map[string]int {
	uniq := map[string]int{}
	uniq[fmt.Sprintf("%d:%d", pos[0], pos[1])]++

	directions := [][]int{matrix.DIR_UP, matrix.DIR_RIGHT, matrix.DIR_DOWN, matrix.DIR_LEFT}
	dir := 0
	for {
		pos[0] += directions[dir][0]
		pos[1] += directions[dir][1]

		if pos[0] < 0 || pos[0] >= len(m) || pos[1] < 0 || pos[1] >= len(m[0]) {
			break
		}

		if m[pos[0]][pos[1]] == `#` {

			pos[0] -= directions[dir][0]
			pos[1] -= directions[dir][1]
			if dir == 3 {
				dir = 0
			} else {
				dir++
			}
			continue
		}

		uniq[fmt.Sprintf("%d:%d", pos[0], pos[1])]++
	}

	return uniq
}

func part2(s []byte) int {
	score := 0
	m := matrix.New(s, "")

	pos := []int{}
CENTER:
	for i, r := range m {
		for j, c := range r {
			if c == "^" {
				pos = []int{i, j}
				break CENTER
			}
		}
	}

	visited := getVisited(slices.Clone(pos), m)

	tests := []matrix.Strings{}
	for k := range visited {
		parts := strings.Split(k, `:`)

		i, _ := strconv.Atoi(parts[0])
		j, _ := strconv.Atoi(parts[1])
		n := m.Clone()
		n[i][j] = `#`
		tests = append(tests, n)
	}

	for _, test := range tests {
		if hasLoop(slices.Clone(pos), test) {
			score++
		}
	}

	return score
}

func hasLoop(pos []int, m matrix.Strings) bool {
	uniq := map[string]int{}
	uniq[fmt.Sprintf("%d:%d", pos[0], pos[1])]++

	directions := [][]int{matrix.DIR_UP, matrix.DIR_RIGHT, matrix.DIR_DOWN, matrix.DIR_LEFT}
	dir := 0
	for {
		pos[0] += directions[dir][0]
		pos[1] += directions[dir][1]

		if pos[0] < 0 || pos[0] >= len(m) || pos[1] < 0 || pos[1] >= len(m[0]) {
			return false
		}

		if m[pos[0]][pos[1]] == `#` {
			if _, ok := uniq[fmt.Sprintf("%d:%d:%d", pos[0], pos[1], dir)]; ok {
				return true
			}

			uniq[fmt.Sprintf("%d:%d:%d", pos[0], pos[1], dir)]++
			pos[0] -= directions[dir][0]
			pos[1] -= directions[dir][1]
			if dir == 3 {
				dir = 0
			} else {
				dir++
			}
			continue
		}
	}
	return false
}
