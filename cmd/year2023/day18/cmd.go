package day18

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day18",
	Short: "day18",
	Long:  `day18`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(string(b)))
	logrus.Infof("score part2: %d", part2(string(b)))

}

var dirs = map[string][]int{
	`U`: {-1, 0},
	`D`: {1, 0},
	`L`: {0, -1},
	`R`: {0, 1},
	`3`: {-1, 0},
	`1`: {1, 0},
	`2`: {0, -1},
	`0`: {0, 1},
}

func part1(s string) int64 {
	// var score int = 0

	points := [][]int{{0, 0}}

	boundary := 0
	//R 6 (#70c710)
	for _, line := range strings.Split(s, "\n") {
		item := strings.Split(line, " ")
		dir := dirs[item[0]]
		moves, err := strconv.Atoi(item[1])

		if err != nil {
			panic(err)
		}

		boundary += moves

		row := points[len(points)-1][0]
		column := points[len(points)-1][1]

		points = append(points, []int{row + (dir[0] * moves), column + (dir[1] * moves)})
	}
	area := polgonArea(points)

	interior := area - (float64(boundary) / 2) + 1

	return int64(interior + float64(boundary))
}

func polgonArea(points [][]int) float64 {
	area := 0.0
	j := len(points) - 1

	for i := range points {
		area += (float64(points[j][0]) + float64(points[i][0])) *
			(float64(points[j][1]) - float64(points[i][1]))
		j = i
	}

	return math.Abs(area / 2.0)
}

func part2(s string) int64 {
	// var score int = 0

	points := [][]int{{0, 0}}

	boundary := 0
	for _, line := range strings.Split(s, "\n") {
		item := strings.Split(line, " ")

		inst := item[2]
		inst = inst[2 : len(inst)-1]

		dir := dirs[inst[len(inst)-1:]]

		m, err := strconv.ParseInt(inst[0:len(inst)-1], 16, 0)
		moves := int(m)

		if err != nil {
			panic(err)
		}

		boundary += moves

		row := points[len(points)-1][0]
		column := points[len(points)-1][1]

		points = append(points, []int{row + (dir[0] * moves), column + (dir[1] * moves)})
	}
	area := polgonArea(points)

	interior := area - (float64(boundary) / 2) + 1

	return int64(interior + float64(boundary))
}
