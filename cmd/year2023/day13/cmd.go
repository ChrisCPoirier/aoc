package day13

import (
	"aoc/cmd/common"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day13",
	Short: "day13",
	Long:  `day13`,
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
	logrus.Infof("score part2: %d", part1(string(b), 1))

}

func part1(s string, allowedImperfections int) int64 {
	var score int = 0

	mirrors := [][][]string{}

	for _, lines := range strings.Split(s, "\n\n") {
		g := common.AsGrid(lines, "")
		mirrors = append(mirrors, g)
	}

	for _, mirror := range mirrors {
		t, a, _ := findReflection(mirror, allowedImperfections)

		switch t {
		case `col`:
			score += a + 1
		case `row`:
			score += (a + 1) * 100
		default:
			println(`No relfection found`)
		}
	}

	return int64(score)
}

func findReflection(mirror [][]string, allowedImperfections int) (string, int, int) {
	for i := 0; i < len(mirror)-1; i++ {
		match := true
		imperfections := 0
		for j := range mirror[i] {
			if mirror[i][j] != mirror[i+1][j] {
				imperfections++
				if imperfections > allowedImperfections {
					match = false
					break
				}
			}
		}

		if match && validRowReflection(mirror, i, i+1, allowedImperfections) {
			return `row`, i, i + 1
		}
	}

	for j := 0; j < len(mirror[0])-1; j++ {
		match := true
		imperfections := 0
		for i := range mirror {
			if mirror[i][j] != mirror[i][j+1] {
				imperfections++
				if imperfections > allowedImperfections {
					match = false
					break
				}
			}
		}

		if match && validColReflection(mirror, j, j+1, allowedImperfections) {
			return `col`, j, j + 1
		}
	}

	return ``, -1, -1
}

func validRowReflection(mirror [][]string, above, below, allowedImperfections int) bool {
	checks := min(above, len(mirror)-1-below)
	imperfections := 0
	for i := 0; i <= checks; i++ {
		for j := range mirror[0] {
			if mirror[above-i][j] != mirror[below+i][j] {
				imperfections++
				if imperfections > allowedImperfections {
					return false
				}
			}
		}
	}
	return imperfections == allowedImperfections
}

func validColReflection(mirror [][]string, above, below, allowedImperfections int) bool {
	checks := min(above, len(mirror[0])-1-below)
	imperfections := 0
	for i := 0; i <= checks; i++ {
		for j := range mirror {
			if mirror[j][above-i] != mirror[j][below+i] {
				imperfections++
				if imperfections > allowedImperfections {
					return false
				}
			}
		}
	}
	return imperfections == allowedImperfections
}
