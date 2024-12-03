package day10

import (
	"aoc/cmd/matrix"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day10",
	Short: "day10",
	Long:  `day10`,
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
	// var score int64 = 0
	m := matrix.New(s, "")

	x, y := getStart(m)

	// fmt.Printf("%#v %d %d", m, y, x)
	visited := map[string]bool{}

	// visited[fmt.Sprintf("%d,%d", x, y)] = true

	winners := [][][]int{}
	winners = append(winners, move(m, x, y, x+1, y, visited))
	winners = append(winners, move(m, x, y, x-1, y, visited))
	winners = append(winners, move(m, x, y, x, y-1, visited))
	winners = append(winners, move(m, x, y, x, y+1, visited))

	loop := [][]int{}

	// fmt.Printf("%#v", winners)
	for _, winner := range winners {
		if len(winner) == 0 {
			continue
		}
		if winner[len(winner)-1][0] == x && winner[len(winner)-1][1] == y {
			if len(winner) > len(loop) {
				loop = winner
			}
		}
	}

	return int64(len(loop) / 2)
}

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

func move(m [][]string, fromX, fromY, toX, toY int, visited map[string]bool) [][]int {
	if visited[fmt.Sprintf("%d,%d", toX, toY)] {
		return nil
	}

	visited[fmt.Sprintf("%d,%d", toX, toY)] = true

	if toX == -1 || toY == -1 || toX > len(m)-1 || toY > len(m[0])-1 {
		return nil
	}

	dX := fromX - toX
	dY := fromY - toY
	c := m[toX][toY]

	if c == `.` {
		return nil
	}

	if c == `S` {
		return [][]int{{toX, toY}}
	}

	if dX == 1 && strings.Contains(`|7F`, c) {
		if c == `|` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX-1, toY, visited)...)
		}

		if c == `7` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX, toY-1, visited)...)
		}

		if c == `F` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX, toY+1, visited)...)
		}
	}

	if dX == -1 && strings.Contains(`|LJ`, c) {
		if c == `|` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX+1, toY, visited)...)
		}
		if c == `J` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX, toY-1, visited)...)
		}

		if c == `L` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX, toY+1, visited)...)
		}
	}

	if dY == -1 && strings.Contains(`-7J`, c) {
		if c == `-` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX, toY+1, visited)...)
		}

		if c == `7` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX+1, toY, visited)...)
		}

		if c == `J` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX-1, toY, visited)...)
		}
	}

	if dY == 1 && strings.Contains(`-FL`, c) {
		if c == `-` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX, toY-1, visited)...)
		}

		if c == `F` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX+1, toY, visited)...)
		}

		if c == `L` {
			return append([][]int{{toX, toY}}, move(m, toX, toY, toX-1, toY, visited)...)
		}
	}
	return nil
}

func getStart(m [][]string) (int, int) {
	for x, row := range m {
		for y, v := range row {
			if v == `S` {
				return x, y
			}
		}
	}
	return -1, -1
}

func part2(s string) int64 {
	// var score int64 = 0
	m := matrix.New(s, "")

	x, y := getStart(m)

	// fmt.Printf("%#v %d %d", m, y, x)
	visited := map[string]bool{}

	// visited[fmt.Sprintf("%d,%d", x, y)] = true

	winners := [][][]int{}
	winners = append(winners, move(m, x, y, x+1, y, visited))
	winners = append(winners, move(m, x, y, x-1, y, visited))
	winners = append(winners, move(m, x, y, x, y-1, visited))
	winners = append(winners, move(m, x, y, x, y+1, visited))

	loop := [][]int{}

	// fmt.Printf("%#v", winners)
	for _, winner := range winners {
		if len(winner) == 0 {
			continue
		}
		if winner[len(winner)-1][0] == x && winner[len(winner)-1][1] == y {
			if len(winner) > len(loop) {
				loop = winner
			}
		}
	}

	lMap := map[string]bool{}
	for _, l := range loop {
		lMap[fmt.Sprintf("%d,%d", l[0], l[1])] = true
	}

	for x, row := range m {
		for y := range row {
			if lMap[fmt.Sprintf("%d,%d", x, y)] {
				continue
			}
			m[x][y] = "."
		}
	}

	m = ReplaceS(m, lMap)

	// score := 0

	fmt.Printf("%#v\n\n", m)

	contained := map[string]bool{}
	for x, rows := range m {
		up := false
		inLoop := false
		for y := range rows {
			c := m[x][y]
			if c == `|` {
				up = false
				inLoop = !inLoop
			}
			if c == `F` || c == `L` {
				up = c == "L"
			}

			if c == `7` || c == `J` {
				compare := `J`
				if !up {
					compare = `7`
				}
				if c != compare {
					inLoop = !inLoop
				}
				up = false
			}

			if inLoop && !lMap[fmt.Sprintf("%d,%d", x, y)] {
				contained[fmt.Sprintf("%d,%d", x, y)] = true
			}
		}
	}

	fmt.Printf("%#v\n", lMap)
	fmt.Printf("%#v\n", contained)
	fmt.Printf("%d %d\n", len(m), len(m[0]))

	return int64(len(contained))
}

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
func ReplaceS(m [][]string, lMap map[string]bool) [][]string {
	for x, row := range m {
		for y, v := range row {
			if v == `S` {
				if lMap[fmt.Sprintf("%d,%d", x-1, y)] && strings.Contains("7F|", m[x-1][y]) &&
					lMap[fmt.Sprintf("%d,%d", x+1, y)] && strings.Contains("LJ|", m[x+1][y]) {
					m[x][y] = `|`
					return m
				}
				if lMap[fmt.Sprintf("%d,%d", x-1, y)] && strings.Contains("7F|", m[x-1][y]) &&
					lMap[fmt.Sprintf("%d,%d", x, y+1)] && strings.Contains("-J7", m[x][y+1]) {
					m[x][y] = `L`
					return m
				}
				if lMap[fmt.Sprintf("%d,%d", x-1, y)] && strings.Contains("7F|", m[x-1][y]) &&
					lMap[fmt.Sprintf("%d,%d", x, y-1)] && strings.Contains("-LF", m[x][y+1]) {
					m[x][y] = `J`
					return m
				}
				if lMap[fmt.Sprintf("%d,%d", x, y-1)] && strings.Contains("LF-", m[x][y-1]) &&
					lMap[fmt.Sprintf("%d,%d", x, y+1)] && strings.Contains("J7-", m[x][y+1]) {
					m[x][y] = `-`
					return m
				}
				if lMap[fmt.Sprintf("%d,%d", x-1, y)] && strings.Contains("7F|", m[x-1][y]) &&
					lMap[fmt.Sprintf("%d,%d", x, y+1)] && strings.Contains("J7-", m[x][y+1]) {
					m[x][y] = `F`
					return m
				}
				if lMap[fmt.Sprintf("%d,%d", x+1, y)] && strings.Contains("LJ|", m[x+1][y]) &&
					lMap[fmt.Sprintf("%d,%d", x, y-1)] && strings.Contains("-LF", m[x][y-1]) {
					m[x][y] = `7`
					return m
				}
			}
		}
	}
	return m
}
