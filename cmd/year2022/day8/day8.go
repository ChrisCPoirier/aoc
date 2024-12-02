package day8

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day8",
	Short: "day8",
	Long:  `day8`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func execute() {
	b, err := os.ReadFile(`data/day8-1.txt`)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Day 8 part 1: %d", aoc(string(b)))
}

func aoc(s string) int {
	return 0
}

func transform(s string) [][]int {

	lines := strings.Split(s, "\n")

	rows := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, r := range line {
			i, _ := strconv.Atoi(string(r))
			row = append(row, i)
		}
		rows = append(rows, row)
	}

	return rows
}

type point struct {
	x     int
	y     int
	val   int
	left  bool
	right bool
}

func getVisible(grid [][]int) map[string]int {

	for x, row := range grid {
		visible := []point{}
		for y, i := range row {
			if len(visible) == 0 {
				p := point{x: x, y: y, val: i, left: true, right: true}
				visible = append(visible, p)
				continue
			}

			vleft := false
			p := visible[x-1]
			if p.val < i || p.val == i {

				p.right = false
				visible[x-1] = p
			}

			if visible[x-1].val < i {
				vleft = true
			}

			visible = append(visible, point{x: x, y: y, val: i, left: vleft, right: true})

		}

	}

	return map[string]int{}
}

func getVisible2(grid [][]int) map[string]bool {

	visible := map[string]bool{}
	for x, row := range grid {
		for y, _ := range row {
			if isVisible(grid, x, y) {
				visible[fmt.Sprintf("%d:%d", x, y)] = true
			}
		}

	}

	return visible
}

func isVisible(grid [][]int, x, y int) bool {

	if x == 0 || y == 0 || x == len(grid)-1 || y == len(grid[0])-1 {
		return true
	}

	for i := x + 1; i < len(grid)-1; i++ {
		if grid[x][y] < grid[x][y] {
			return true
		}
	}
	return false
}
