package common

import (
	"fmt"
	"strings"
)

func AsGrid(s, sep string) Grid {
	grid := Grid{}

	for _, line := range strings.Split(s, "\n") {
		println(line)
		grid = append(grid, strings.Split(line, sep))
		fmt.Printf("%#v\n", grid)
	}

	return grid
}

func FieldsAsGrid(s string) Grid {
	grid := Grid{}

	for _, line := range strings.Split(s, "\n") {
		grid = append(grid, strings.Fields(line))
	}

	return grid
}

type Grid [][]string

func (g Grid) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%s\n", strings.Join(item, " "))
	}
	return out

}

func (g GridI) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%d\n", item)
	}
	return out

}

func (g GridF) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%.0f\n", item)
	}
	return out

}

func (g Grid) AsGridI() GridI {
	n := GridI{}
	for _, row := range g {
		n = append(n, AsInts(row))
	}
	return n
}

func (g Grid) AsGridF() GridF {
	n := GridF{}
	for _, row := range g {
		n = append(n, AsFloats(row))
	}
	return n
}

func (g Grid) Rotate() Grid {

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
