package day6

import (
	"os"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day6",
	Short: "day6",
	Long:  `day6`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func execute() {
	b, err := os.ReadFile(`data/day6-1.txt`)

	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Day 6 part 1: %d", startOfPacketMarker(string(b)))
	logrus.Infof("Day 6 part 2: %d", startOfMessageMarker(b))
}

func startOfPacketMarker(s string) int {

	for i := 3; i < len(s)-1; i++ {
		if s[i] == s[i-1] || s[i] == s[i-2] || s[i] == s[i-3] {
			continue
		}
		if s[i-1] == s[i-2] || s[i-1] == s[i-3] {
			continue
		}

		if s[i-2] == s[i-3] {
			continue
		}
		return i + 1
	}
	return 0
}

const mml = 14

func startOfMessageMarker(b []byte) int {
	for i := mml - 1; i < len(b)-1; i++ {
		u := lo.Uniq(b[i-(mml-1) : i+1])
		if len(u) == mml {
			return i + 1
		}
	}
	return 0
}
