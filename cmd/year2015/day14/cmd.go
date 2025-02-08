package day14

import (
	"aoc/cmd/common"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day14",
	Short: "day14",
	Long:  `day14`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, func(s []byte) int { return part1(s, 2503) }, `part 1`)
	common.Run(parent, command, 1, func(s []byte) int { return part2(s, 2503) }, `part 2`)
}

type reindeer struct {
	name     string
	flySpeed int
	flyTime  int
	restTime int
	traveled int
	time     int
}

func (r *reindeer) Tick() {
	r.time++
	if r.time == r.flyTime+r.restTime {
		r.time = 0
		return
	}

	if r.time <= r.flyTime {
		r.traveled++
	}
}

func (r *reindeer) Traveled() int {
	return r.traveled * r.flySpeed
}

func part1(s []byte, time int) int {
	distance := 0

	for _, r := range parse(string(s)) {
		d := fly(r, time)
		if d > distance {
			distance = d
		}
	}

	return distance
}

func part2(s []byte, time int) int {
	scores := map[string]int{}
	reindeers := parse(string(s))

	for range time {
		furthest := 0
		for _, r := range reindeers {
			r.Tick()
			if r.Traveled() > furthest {
				furthest = r.Traveled()
			}
		}

		for _, r := range reindeers {
			if r.Traveled() == furthest {
				scores[r.name]++
			}
		}
	}

	furthest := 0
	for _, score := range scores {
		if score > furthest {
			furthest = score
		}
	}

	logrus.Infof("scores: %#v", scores)
	return furthest
}

var re = regexp.MustCompile(`(\w+) can fly (\d+) km\/s for (\d+) seconds, but then must rest for (\d+) seconds\.`)

func parse(s string) []*reindeer {
	reindeers := []*reindeer{}

	for _, match := range re.FindAllStringSubmatch(string(s), len(s)) {
		name := match[1]
		flySpeed, _ := strconv.Atoi(match[2])
		flyTime, _ := strconv.Atoi(match[3])
		rest, _ := strconv.Atoi(match[4])

		reindeers = append(reindeers, &reindeer{name: name, flySpeed: flySpeed, flyTime: flyTime, restTime: rest})
	}

	return reindeers
}

func fly(r *reindeer, t int) int {
	for range t {
		r.Tick()
	}

	return r.traveled * r.flySpeed
}

// Does not work for part 2
func flyWithMaths(r reindeer, t int) int {
	cycleTime := r.flyTime + r.restTime
	fullCycles := t / cycleTime

	//After calculating our full cylcles, calculate how much remaining flight time we have
	remainingFlightTime := t - (cycleTime * fullCycles)
	if remainingFlightTime > r.flyTime {
		remainingFlightTime = r.flyTime
	}

	return (fullCycles * r.flySpeed * r.flyTime) + (remainingFlightTime * r.flySpeed)

}
