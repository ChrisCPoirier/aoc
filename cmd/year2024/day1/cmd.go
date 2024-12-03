package day1

import (
	"aoc/cmd/matrix"
	"fmt"
	"math"
	"os"
	"slices"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1",
	Long:  `day1`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	b2, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/2.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %.0f", part1(b))
	logrus.Infof("score part2: %.0f", part2(b2))
}

func part1(s []byte) float64 {
	score := 0.0

	g := matrix.
		New(s, "   ").
		Rotate().
		Floats()

	slices.Sort(g[0])
	slices.Sort(g[1])

	for i := range g[0] {
		score += math.Abs(g[0][i] - g[1][i])
	}

	return score
}

func part2(s []byte) float64 {
	score := 0.0

	g := matrix.
		New(string(s), "   ").
		Rotate().
		Floats()

	slices.Sort(g[0])
	slices.Sort(g[1])

	am := mapWithCount(g[0])
	bm := mapWithCount(g[1])

	for k := range am {
		score += (am[k] * k) * bm[k]
	}

	return score
}

func mapWithCount(in []float64) map[float64]float64 {
	out := map[float64]float64{}

	for _, a := range in {
		out[a] = out[a] + 1
	}

	return out
}
