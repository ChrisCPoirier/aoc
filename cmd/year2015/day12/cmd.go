package day12

import (
	"aoc/cmd/common"
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day12",
	Short: "day12",
	Long:  `day12`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

type edge struct {
	cost int
	dest string
}

var reNums = regexp.MustCompile(`-?\d+`)

func part1(s []byte) int {
	sum := 0
	matches := reNums.FindAllString(string(s), len(s))

	for _, match := range matches {
		// logrus.Info(match)
		v, _ := strconv.Atoi(match)
		sum += v
	}

	return sum
}

func part2(s []byte) int {

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(s, &jsonMap)
	if err != nil {
		panic(err)
	}

	// logrus.Infof("%#v", jsonMap)

	return score(jsonMap)
}

func score(m map[string]interface{}) int {
	total := 0

	if _, ok := m["red"]; ok {
		logrus.Infof("found red: %#v", m)
		return total
	}

	for _, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			total += score(v.(map[string]interface{}))
		case float64:
			mv := v.(float64)
			total += int(mv)
		case string:
			mv := v.(string)
			if mv == "red" {
				return 0
			}
		case []interface{}:
			total += scoreArray(v.([]interface{}))
		}
	}

	return total
}

func scoreArray(items []interface{}) int {
	total := 0

	for _, item := range items {
		switch item.(type) {
		case float64:
			mv := item.(float64)
			total += int(mv)
		case []interface{}:
			total += scoreArray(item.([]interface{}))
		case map[string]interface{}:
			total += score(item.(map[string]interface{}))
		}
	}

	return total
}
