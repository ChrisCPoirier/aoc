package day13

import (
	"aoc/cmd/common"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day13",
	Long:  `day13`,
	Use:   "day13",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

type machine struct {
	ax, ay, bx, by, px, py float64
}

var reButton = regexp.MustCompile(`Button \w: \w\+(\d+), \w\+(\d+)`)
var rePrize = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

func part1(s []byte) int {

	machines := getMachines(string(s))

	score := 0

	for _, m := range machines {
		score += getMachineScore2(&m)
	}

	return score
}

func part2(s []byte) int {

	machines := getMachines(string(s))

	score := 0

	for _, m := range machines {
		m.px += 10000000000000
		m.py += 10000000000000
		score += getMachineScore2(&m)
	}

	return score
}

func getMachines(s string) []machine {
	sections := strings.Split(s, "\n\n")
	machines := []machine{}
	for _, section := range sections {
		lines := strings.Split(section, "\n")

		match := reButton.FindStringSubmatch(lines[0])
		ax, _ := strconv.ParseFloat(match[1], 64)
		ay, _ := strconv.ParseFloat(match[2], 64)

		match = reButton.FindStringSubmatch(lines[1])
		bx, _ := strconv.ParseFloat(match[1], 64)
		by, _ := strconv.ParseFloat(match[2], 64)

		match = rePrize.FindStringSubmatch(lines[2])
		px, _ := strconv.ParseFloat(match[1], 64)
		py, _ := strconv.ParseFloat(match[2], 64)

		machines = append(machines, machine{ax: ax, ay: ay, bx: bx, by: by, px: px, py: py})
	}
	return machines
}

func getMachineScore2(i *machine) int {
	n := (i.px*i.by - i.py*i.bx) / (i.ax*i.by - i.ay*i.bx)
	m := (i.py - n*i.ay) / i.by

	if n == math.Trunc(n) && m == math.Trunc(m) {
		return int(3*n + m)
	}

	return 0
}
