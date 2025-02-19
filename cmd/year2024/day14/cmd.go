package day14

import (
	"aoc/cmd/common"
	"aoc/cmd/display"
	"aoc/cmd/grid"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day14",
	Long:  `day14`,
	Use:   "day14",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, func(b []byte) int { return part1(b, 101, 103) }, "part 1")
	common.Run(parent, command, 1, func(b []byte) int { return part2(b, 101, 103) }, "part 2")
}

var reLine = regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

type robot struct {
	px, py int
	vx, vy int
}

func part1(s []byte, bx, by int) int {
	logrus.Infof("bx:%d by:%d", bx, by)
	robots := parseRobots(s)
	return score(robots, bx, by)
}

func part2(s []byte, bx, by int) int {
	logrus.Infof("bx:%d by:%d", bx, by)
	robots := parseRobots(s)
	return score2(robots, bx, by)
}

func parseRobots(b []byte) []robot {
	robots := []robot{}
	for _, match := range reLine.FindAllStringSubmatch(string(b), -1) {
		px, _ := strconv.Atoi(match[1])
		py, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])

		robots = append(robots, robot{
			px: px, py: py, vx: vx, vy: vy,
		})
	}
	return robots
}

func key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func setInBounds(i, b int) int {
	if i < 0 {
		return i + b
	}
	return i
}

func score(robots []robot, bx, by int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, r := range robots {
		//velocity * 100 iterations + starting position Remainder of bounding x/y for final box position
		nx := (r.vx*100 + r.px) % bx
		ny := (r.vy*100 + r.py) % by

		// if we are negative we need to roll back to the previous box
		nx = setInBounds(nx, bx)
		ny = setInBounds(ny, by)

		switch {
		case nx == bx/2 || ny == by/2:
			continue
		case nx < bx/2 && ny < by/2:
			q1++
		case nx < bx/2 && ny > by/2:
			q2++
		case ny < by/2:
			q3++
		default:
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func score2(rb []robot, bx, by int) int {
	score := 0

	g := grid.Strings{}.Fill(bx, by, `.`)
	display.New(g)
	for i := range 100000 {
		g = grid.Strings{}.Fill(bx, by, `.`)
		for i, r := range rb {
			rb[i].px += r.vx
			rb[i].py += r.vy

			if rb[i].px < 0 {
				rb[i].px += bx
			}

			if rb[i].py < 0 {
				rb[i].py += by
			}

			if rb[i].px >= bx {
				rb[i].px -= bx
			}

			if rb[i].py >= by {
				rb[i].py -= by
			}

			g[rb[i].px][rb[i].py] = `#`
		}

		if i%10 == 0 || maybeTree(g) {

		}

		if maybeTree(g) {
			logrus.Infof("iter: %d", i)
			score = i + 1
			logrus.Info("maybe tree")
			fmt.Println(g.Rotate().Pretty())
			break
		}

	}

	return score
}

var partTree = `########`

func maybeTree(g grid.Strings) bool {
	for _, r := range g {
		if strings.Contains(strings.Join(r, ``), partTree) {
			return true
		}
	}
	return false
}
