package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day9",
	Short: "day9",
	Long:  `day9`,
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

func part1(s string) int64 {
	var score int64 = 0
	for _, line := range strings.Split(s, "\n") {
		// fmt.Printf("%s\n", line)
		items := strings.Fields(line)

		n := []int64{}
		for _, item := range items {
			i, err := strconv.Atoi(item)

			if err != nil {
				panic(err)
			}

			n = append(n, int64(i))
		}

		score += nextVal(n)

	}
	return score
}

func part2(s string) int64 {
	var score int64 = 0
	for _, line := range strings.Split(s, "\n") {
		// fmt.Printf("%s\n", line)
		items := strings.Fields(line)

		n := []int64{}
		for _, item := range items {
			i, err := strconv.Atoi(item)

			if err != nil {
				panic(err)
			}

			n = append(n, int64(i))
		}

		n = lo.Reverse(n)
		score += nextVal(n)

	}
	return score
}

func nextVal(items []int64) int64 {

	diff := true
	stack := [][]int64{items}

	for diff {
		bucket := []int64{}
		work := stack[len(stack)-1]
		diff = false
		for i := 1; i < len(work); i++ {
			next := work[i] - work[i-1]
			bucket = append(bucket, next)
			if next != 0 {
				diff = true
			}
		}
		if len(bucket) > 0 {
			stack = append(stack, bucket)
		}
	}

	var val int64 = 0
	for i := len(stack) - 1; i >= 0; i-- {
		val += stack[i][len(stack[i])-1]
	}

	return val
}
