package grid

import (
	"aoc/cmd/common"
	"fmt"
	"strings"
)

type text interface {
	string | []byte
}

func New[T1, T2 text](in T1, sep T2) Strings {
	grid := Strings{}

	for _, line := range strings.Split(string(in), "\n") {
		grid = append(grid, strings.Split(line, string(sep)))
	}

	return grid
}

func FieldsAsGrid[T1 text](s T1) Strings {
	grid := Strings{}

	for _, line := range strings.Split(string(s), "\n") {
		grid = append(grid, strings.Fields(line))
	}

	return grid
}

type Strings [][]string

func (g Strings) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%s\n", strings.Join(item, " "))
	}
	return out

}

func (g Ints) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%d\n", item)
	}
	return out

}

func (g Floats) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%.0f\n", item)
	}
	return out

}

func (g Strings) Ints() Ints {
	n := Ints{}
	for _, row := range g {
		n = append(n, common.AsInts(row))
	}
	return n
}

func (g Strings) Floats() Floats {
	n := Floats{}
	for _, row := range g {
		n = append(n, common.AsFloats(row))
	}
	return n
}

func (g Strings) Rotate() Strings {
	if len(g[0]) != len(g) {
		return g.rotateUnequal()
	}

	// reverse the grid
	for i, j := 0, len(g)-1; i < j; i, j = i+1, j-1 {
		g[i], g[j] = g[j], g[i]
	}

	// transpose the grid
	for i := 0; i < len(g); i++ {
		for j := 0; j < i; j++ {
			g[i][j], g[j][i] = g[j][i], g[i][j]
		}
	}
	return g
}

func (g Strings) rotateUnequal() Strings {
	n := make(Strings, len(g[0]))
	for _, row := range g {
		for i, col := range row {
			n[i] = append([]string{col}, n[i]...)
		}
	}

	g = n
	return g
}
