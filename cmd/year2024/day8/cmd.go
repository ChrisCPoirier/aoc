package day8

import (
	"aoc/cmd/common"
	"aoc/cmd/matrix"
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day8",
	Long:  `day8`,
	Use:   "day8",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
}

type block struct {
	id   int
	size int
}

type antiNode struct {
	pos  []int
	name string
}

func part1(s []byte) int {
	in := matrix.New(s, ``)
	antennaGroup := getAntennas(in)

	antinodes := map[string][]int{}
	for _, antennas := range antennaGroup {
		for _, a := range antennas {
			for _, b := range antennas {
				if a[0] == b[0] && a[1] == b[1] {
					continue
				}

				antinode := []int{a[0] + (a[0] - b[0]), a[1] + (a[1] - b[1])}

				if !in.InBound(antinode[0], antinode[1]) {
					continue
				}
				antinodes[fmt.Sprintf("%#v", antinode)] = antinode
			}

		}
	}

	return len(antinodes)
}

func part2(s []byte) int {
	m := matrix.New(s, ``)
	antennaGroup := getAntennas(m)

	antinodes := map[string][]int{}

	//every antenna is also an antinode
	for _, group := range antennaGroup {
		for _, a := range group {
			antinode := []int{a[0], a[1]}
			antinodes[fmt.Sprintf("%#v", antinode)] = antinode
		}
	}

	for _, antennas := range antennaGroup {
		for _, a := range antennas {
			for _, b := range antennas {
				if a[0] == b[0] && a[1] == b[1] {
					continue
				}

				dr := (a[0] - b[0])
				dc := a[1] - b[1]
				antinode := []int{a[0] + dr, a[1] + dc}
				for m.InBound(antinode[0], antinode[1]) {
					antinodes[fmt.Sprintf("%#v", antinode)] = antinode
					antinode = []int{antinode[0] + dr, antinode[1] + dc}
				}
			}
		}
	}

	return len(antinodes)
}

func getAntennas(in matrix.Strings) map[string][][]int {
	antennas := map[string][][]int{}
	for i, r := range in {
		for j, c := range r {
			if c != "." {
				antennas[c] = append(antennas[c], []int{i, j})
			}
		}
	}
	return antennas
}
