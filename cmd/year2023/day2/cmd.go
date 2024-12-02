package day2

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day2",
	Long:  `day2`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(b))
	logrus.Infof("score part2: %d", part2(b))
}

func part1(s []byte) int {
	score := 0
	for _, line := range bytes.Split(s, []byte("\n")) {
		score += getScore(line)
	}

	return score
}

func part2(s []byte) int {
	score := 0
	for _, line := range bytes.Split(s, []byte("\n")) {
		score += getScore2(line)
	}

	return score
}

var blue = 14
var red = 12
var green = 13

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func getScore(s []byte) int {

	played := bytes.Split(s, []byte(`:`))
	id := bytes.Split(played[0], []byte(` `))[1]

	idInt, _ := strconv.Atoi(string(id))

	for _, pull := range bytes.Split(bytes.Trim(played[1], ` `), []byte(`;`)) {
		for _, dice := range bytes.Split(pull, []byte(`,`)) {
			dice = bytes.Trim(dice, ` `)
			v := bytes.Split(dice, []byte(` `))

			i, _ := strconv.Atoi(string(v[0]))

			switch string(v[1]) {
			case "blue":
				if i > blue {
					return 0
				}
			case "green":
				if i > green {
					return 0
				}
			case "red":
				if i > red {
					return 0
				}
			}
		}
	}

	return idInt
}

func getScore2(s []byte) int {

	played := bytes.Split(s, []byte(`:`))
	// id := bytes.Split(played[0], []byte(` `))[1]

	// idInt, _ := strconv.Atoi(string(id))

	var maxBlue, maxGreen, maxRed int

	for _, pull := range bytes.Split(bytes.Trim(played[1], ` `), []byte(`;`)) {
		for _, dice := range bytes.Split(pull, []byte(`,`)) {
			dice = bytes.Trim(dice, ` `)
			v := bytes.Split(dice, []byte(` `))

			i, _ := strconv.Atoi(string(v[0]))

			switch string(v[1]) {
			case "blue":
				if i > maxBlue {
					maxBlue = i
				}
			case "green":
				if i > maxGreen {
					maxGreen = i
				}
			case "red":
				if i > maxRed {
					maxRed = i
				}
			}
		}
	}

	return maxBlue * maxRed * maxGreen
}
