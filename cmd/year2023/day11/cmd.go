package day11

import (
	"aoc/cmd/common"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day11",
	Short: "day11",
	Long:  `day11`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(string(b), 1))
	logrus.Infof("score part1: %d", part1(string(b), 1000000))
	// logrus.Infof("score part2: %d", part2(string(b)))

}

func part1(s string, scale int) int64 {
	// var score int64 = 0
	m := common.AsGrid(s, "")

	galaxies := [][]int{}

	for x, row := range m {
		for y, v := range row {
			if v == `#` {
				galaxies = append(galaxies, []int{x, y})
			}
		}
	}

	expandedSpace := map[string]bool{}

	for x, row := range m {
		expanded := true
		for _, v := range row {
			if v == `#` {
				expanded = false
				break
			}
		}
		if expanded {
			expandedSpace[fmt.Sprintf("x:%d", x)] = true
		}
	}

	for y := range m[0] {
		expanded := true
		for x := range m {
			if m[x][y] == `#` {
				expanded = false
				break
			}
		}
		if expanded {
			expandedSpace[fmt.Sprintf("y:%d", y)] = true
		}
	}

	pairLen := (len(galaxies) * (len(galaxies) - 1)) / 2

	paths := map[string]int{}

	for len(paths) < pairLen {
		for i, galaxy := range galaxies {
			for j, toGalaxy := range galaxies {
				if i == j {
					continue
				}

				min := min(i, j)
				max := max(i, j)
				key := fmt.Sprintf("%d -> %d", min, max)

				if _, ok := paths[key]; !ok {
					paths[key] = calcDistance(galaxy, toGalaxy, expandedSpace, scale)
				}
			}
		}
	}

	var score = 0
	for _, v := range paths {
		score += v
	}

	return int64(score)
}

func calcDistance(from, to []int, expandedSpace map[string]bool, scale int) int {
	distance := 0
	for x := min(from[0], to[0]); x < max(from[0], to[0]); x++ {
		if _, ok := expandedSpace[fmt.Sprintf("x:%d", x)]; ok {
			distance += scale
			continue
		}
		distance += 1
	}

	for y := min(from[1], to[1]); y < max(from[1], to[1]); y++ {
		if _, ok := expandedSpace[fmt.Sprintf("y:%d", y)]; ok {
			distance += scale
			continue
		}
		distance += 1
	}

	return distance
}
