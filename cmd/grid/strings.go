package grid

import (
	"aoc/cmd/common"
	"fmt"
	"slices"
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
		out += fmt.Sprintf("%s\n", strings.Join(item, ""))
	}
	return out

}

type step struct {
	R    int
	C    int
	Path [][]int
}

// sr (start r), sc (start c), er (end r), ec (end c)
// walk the grid starting at sr,sc and find the path to er,ec
func (s Strings) BFS(sr, sc, er, ec, maxDepth int) step {
	queue := []step{{
		R:    sr,
		C:    sc,
		Path: [][]int{{sr, sc}},
	}}
	var p step

	optimal := step{}

	visited := map[string]int{}
	for len(queue) > 0 {
		p, queue = queue[0], queue[1:]

		if v, ok := visited[Key(p.R, p.C)]; ok {
			if v < len(p.Path) {
				continue
			}
		}

		visited[Key(p.R, p.C)] = len(p.Path)

		if p.R == er && p.C == ec {
			if len(optimal.Path) == 0 || len(p.Path) < len(optimal.Path) {
				optimal = p
			}
		}

		for _, dir := range DIR_CROSS {
			nr := p.R + dir[0]
			nc := p.C + dir[1]
			if !s.InBound(nr, nc) {
				continue
			}

			if s[nr][nc] == `#` {
				continue
			}

			queue = append(queue, step{
				R:    nr,
				C:    nc,
				Path: append(slices.Clone(p.Path), []int{nr, nc}),
			})
		}
	}
	return optimal
}

func (g Strings) Fill(rl, cl int, s string) Strings {
	g = Strings{}
	for range rl {
		g = append(g, strings.Split(strings.Repeat(`.`, cl), ``))
	}
	return g
}

func Key(r, c int) string {
	return fmt.Sprintf("%d:%d", r, c)
}

func (s Strings) InBound(i, j int) bool {
	return !(i < 0 || i >= len(s) || j < 0 || j >= len(s[0]))
}

func (s Strings) FindCell(search string) []int {
	for i, r := range s {
		for j, c := range r {
			if c == search {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func (s Strings) Clone() Strings {
	n := Strings{}
	for _, r := range s {
		n = append(n, slices.Clone(r))
	}
	return n
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
