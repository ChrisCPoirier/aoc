package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day3",
	Long:  `day3`,
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

func part1(s string) int {
	score := 0

	lines := strings.Split(s, "\n")

	ns := -1
	ne := -1

	for x, line := range lines {

		for y, r := range line {

			if r >= '0' && r <= '9' {
				if ns == -1 {
					ns = y
					// continue
				}
				ne = y
			}

			if (r == '.' || isSymbol(byte(r))) && ns != -1 {
				//handle number check
				i, err := strconv.Atoi(lines[x][ns : ne+1])

				if err != nil {
					panic(err)
				}

				if found, _, _ := hasAdjacent(lines, x, ns, ne, isSymbol); found {

					fmt.Printf("Number for: %d,%d,%d (%s) is %d\n", x, ns, ne+1, lines[x][ns:ne+1], i)
					score += i
				}
				// score += i
				ns = -1
				ne = -1
			}
		}
		if ns != -1 {
			i, err := strconv.Atoi(lines[x][ns : ne+1])

			if err != nil {
				panic(err)
			}

			if found, _, _ := hasAdjacent(lines, x, ns, ne, isSymbol); found {
				fmt.Printf("Number for: %d,%d,%d (%s) is %d\n", x, ns, ne+1, lines[x][ns:ne+1], i)
				score += i
			}

			//handle number check
			ns = -1
			ne = -1
		}
	}

	return score
}

func hasAdjacent(input []string, x, ys, ye int, check func(byte) bool) (bool, int, int) {
	fmt.Printf("x: %d, ys: %d, ye: %d, check: %s line: %s\n", x, ys, ye, input[x][ys:ye+1], input[x])

	for i := ys; i <= ye; i++ {
		if x > 0 {
			if check(input[x-1][i]) {
				return true, x - 1, i
			}
		}

		if x < len(input)-1 {
			if check(input[x+1][i]) {
				return true, x + 1, i
			}
		}
	}

	if ys > 0 {
		if check(input[x][ys-1]) {
			return true, x, ys - 1
		}
		if x > 0 {
			if check(input[x-1][ys-1]) {
				return true, x - 1, ys - 1
			}
		}
		if x < len(input)-1 {
			if check(input[x+1][ys-1]) {
				return true, x + 1, ys - 1
			}
		}
	}

	if ye < len(input[x])-1 {
		if check(input[x][ye+1]) {
			return true, x, ye + 1
		}
		if x > 0 {
			if check(input[x-1][ye+1]) {
				return true, x - 1, ye + 1
			}
		}
		if x < len(input)-1 {
			if check(input[x+1][ye+1]) {
				return true, x + 1, ye + 1
			}
		}
	}

	return false, -1, -1
}

func isSymbol(b byte) bool {
	// return lo.Contains([]byte("$*#+"), b)
	if b >= '0' && b <= '9' ||
		b == '.' {
		return false
	}
	return true
}

func isAterik(b byte) bool {
	return b == '*'
}

func part2(s string) int {
	score := 0

	scorable := map[string][]int{}

	lines := strings.Split(s, "\n")

	ns := -1
	ne := -1

	for x, line := range lines {

		for y, r := range line {

			if r >= '0' && r <= '9' {
				if ns == -1 {
					ns = y
					ne = y
					// continue
				}
				ne = y
			}

			if (r == '.' || isSymbol(byte(r))) && ns != -1 {
				//handle number check
				i, err := strconv.Atoi(lines[x][ns : ne+1])

				if err != nil {
					panic(err)
				}

				if found, fx, fy := hasAdjacent(lines, x, ns, ne, isAterik); found {

					fmt.Printf("Number for: %d,%d,%d (%s) is %d\n", x, ns, ne+1, lines[x][ns:ne+1], i)
					scorable[fmt.Sprintf("%d,%d", fx, fy)] = append(scorable[fmt.Sprintf("%d,%d", fx, fy)], i)
				}
				// score += i
				ns = -1
				ne = -1
			}
		}
		if ns != -1 {
			i, err := strconv.Atoi(lines[x][ns : ne+1])

			if err != nil {
				panic(err)
			}

			if found, fx, fy := hasAdjacent(lines, x, ns, ne, isAterik); found {
				fmt.Printf("Number for: %d,%d,%d (%s) is %d\n", x, ns, ne+1, lines[x][ns:ne+1], i)
				scorable[fmt.Sprintf("%d,%d", fx, fy)] = append(scorable[fmt.Sprintf("%d,%d", fx, fy)], i)
			}

			//handle number check
			ns = -1
			ne = -1
		}
	}

	for _, v := range scorable {
		if len(v) != 2 {
			continue
		}

		score += v[0] * v[1]
	}

	return score
}
