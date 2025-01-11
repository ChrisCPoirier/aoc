package day5

import (
	"aoc/cmd/common"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day5",
	Short: "day5",
	Long:  `day5`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

func part1(s []byte) int {

	nice := 0
	for _, line := range strings.Split(string(s), "\n") {
		if isNice(line) {
			nice++
		}
	}

	return nice

}

func part2(s []byte) int {
	nice := 0
	for _, line := range strings.Split(string(s), "\n") {
		if isNiceV2(line) {
			nice++
		}
	}

	return nice
}

var disallowed = []string{"ab", "cd", "pq", "xy"}

func isNice(s string) bool {

	for _, item := range disallowed {
		if strings.Contains(s, item) {
			return false
		}
	}

	items := strings.Split(s, ``)

	cnts := common.Counts(items)
	vowles := cnts[`a`] + cnts[`e`] + cnts[`i`] + cnts[`o`] + cnts[`u`]

	if vowles < 3 {
		return false
	}

	for i := 0; i < len(items)-1; i++ {
		if items[i] == items[i+1] {
			return true
		}
	}

	return false
}

func isNiceV2(s string) bool {
	if !hasRepeateWithGap(s) {
		return false
	}

	for _, locs := range getDoubleCharLocations(s) {
		if len(locs) > 1 {
			return true
		}
	}

	return false
}

func getDoubleCharLocations(s string) map[string][]int {
	items := strings.Split(s, ``)

	pos := map[string][]int{}
	for i := 0; i < len(items)-1; i++ {
		if slices.Contains(pos[items[i]+items[i+1]], i) {
			continue
		}
		pos[items[i]+items[i+1]] = append(pos[items[i]+items[i+1]], i+1)
	}

	return pos
}

func hasRepeateWithGap(s string) bool {
	items := strings.Split(s, ``)

	for i := 0; i < len(items)-2; i++ {
		if items[i] == items[i+2] {
			return true
		}
	}

	return false
}
