package day24

import (
	"aoc/cmd/common"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day24",
	Long:  `day24`,
	Use:   "day24",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

func part1(s []byte) int {
	score, _ := strconv.ParseInt(getZoutputs(string(s)), 2, 64)
	return int(score)
}

func part2(s []byte) string {
	gates := NewGates(s)

	zNames := []string{}
	for _, gate := range gates {
		if !strings.HasPrefix(gate.name, `z`) {
			continue
		}
		zNames = append(zNames, gate.name)
	}

	slices.Sort(zNames)

	fixed := []string{}
	for _, gateName := range zNames {
		zGate := gates[gateName]
		fixed = append(fixed, gates.Fix(zGate)...)
	}

	slices.Sort(fixed)
	return strings.Join(fixed, `,`)
}

func getZoutputs(s string) string {
	g := NewGates([]byte(s))

	g.Solve()

	zGates := []string{}

	for _, gate := range g {
		if !strings.HasPrefix(gate.name, `z`) {
			continue
		}
		zGates = append(zGates, gate.name)
	}

	slices.Sort(zGates)
	slices.Reverse(zGates)

	temp := ""
	for _, gate := range zGates {
		temp += fmt.Sprintf("%d", g[gate].value)
	}
	return temp
}
