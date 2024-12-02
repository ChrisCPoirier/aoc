package day1

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"

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

func parseColumns(in []byte) ([]float64, []float64) {
	a := []float64{}
	b := []float64{}

	for _, line := range bytes.Split(in, []byte("\n")) {

		cols := bytes.Split(line, []byte("   "))

		fa, _ := strconv.ParseFloat(string(cols[0]), 64)
		fb, _ := strconv.ParseFloat(string(cols[1]), 64)
		a = append(a, fa)
		b = append(b, fb)
	}
	return a, b
}

func part1(s []byte) float64 {
	score := 0.0

	a, b := parseColumns(s)

	slices.Sort(a)
	slices.Sort(b)

	for i := range a {
		score += math.Abs(a[i] - b[i])
	}

	return score
}

func part2(s []byte) float64 {
	score := 0.0

	a, b := parseColumns(s)

	slices.Sort(a)
	slices.Sort(b)

	am := mapWithCount(a)
	bm := mapWithCount(b)

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
