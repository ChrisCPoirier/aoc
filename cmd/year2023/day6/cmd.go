package day6

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day6",
	Short: "day6",
	Long:  `day6`,
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

var spaceRE = regexp.MustCompile(`\s+`)

func part1(s string) int64 {
	var score int64 = 1

	lines := strings.Split(s, "\n")

	times := strings.Split(strings.Split(spaceRE.ReplaceAllString(lines[0], ` `), ": ")[1], " ")
	distances := strings.Split(strings.Split(spaceRE.ReplaceAllString(lines[1], ` `), ": ")[1], " ")

	for i, ts := range times {
		t, _ := strconv.Atoi(ts)
		d, _ := strconv.Atoi(distances[i])
		var count int64 = 0
		for holdTime := 1; holdTime < t; holdTime++ {
			if holdTime*(t-holdTime) > d {
				count++
			}

		}
		score = score * count
	}

	fmt.Printf("%#v", times)
	fmt.Printf("%#v", distances)

	return score
}

func part2(s string) int64 {
	var score int64 = 1

	lines := strings.Split(s, "\n")

	times := strings.Split(strings.Split(spaceRE.ReplaceAllString(lines[0], ` `), ": ")[1], " ")
	distances := strings.Split(strings.Split(spaceRE.ReplaceAllString(lines[1], ` `), ": ")[1], " ")

	ts := strings.Join(times, "")
	ds := strings.Join(distances, "")

	t, _ := strconv.Atoi(ts)
	d, _ := strconv.Atoi(ds)
	var count int64 = 0
	for holdTime := 1; holdTime < t; holdTime++ {
		if holdTime*(t-holdTime) > d {
			count++
		}

		score = count
	}

	return score
}
