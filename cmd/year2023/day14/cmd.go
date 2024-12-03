package day14

import (
	"aoc/cmd/common"
	"aoc/cmd/grid"
	"fmt"
	"os"
	"slices"
	"strings"

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
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(string(b), 0))
	logrus.Infof("score part2: %d", part1(string(b), 1000000000))

}

func part1(s string, cycles int) int64 {
	// var score int = 0

	g := grid.AsGrid(s, "")

	fmt.Printf("init:\n%s\n", g.Pretty())

	g.Rotate()

	fmt.Printf("first rotate:\n%s\n", g.Pretty())

	if cycles == 0 {
		g = slide(g)
	}

	cache := map[string]int{}

	cycleFound := false
	for i := 0; i < cycles; i++ {
		for range []string{`north`, `west`, `south`, `east`} {
			g = slide(g)
			if !cycleFound {
				if c, ok := cache[strings.ReplaceAll(g.Pretty(), "\n", "")]; ok {
					fmt.Printf("cycle found: first :%d, second: %d\n", c, i)
					cycleFound = true
					i = (cycles - (cycles-c)%(i-c))
					fmt.Printf("cycle found: jumping to: %d\n", i)
				}
			}
			cache[strings.ReplaceAll(g.Pretty(), "\n", "")] = i

			g.Rotate()
		}

	}

	g.Rotate()
	g.Rotate()
	g.Rotate()

	// fmt.Printf("North:\n%s\n", g.Pretty())

	return int64(score(g))
}

func slide(g grid.Grid) grid.Grid {
	for i, row := range g {

		chunks := common.Chunk(row, `#`)

		for j, c := range chunks {
			slices.Sort(c)
			chunks[j] = c
		}
		g[i] = common.Stitch(chunks, `#`)
	}
	return g
}

func score(g grid.Grid) int {
	score := 0
	multipler := len(g)
	for i, row := range g {
		for _, v := range row {
			if v == `O` {
				score += 1 * (multipler - i)
			}
		}
	}
	return score
}
