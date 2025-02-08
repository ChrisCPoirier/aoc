package grid

import "fmt"

type Ints [][]int

func (g Ints) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%d\n", item)
	}
	return out

}

func (in Ints) InBound(i, j int) bool {
	return !(i < 0 || i >= len(in) || j < 0 || j >= len(in[0]))
}

func (g Ints) Fill(rl, cl int, def int) Ints {
	g = Ints{}
	for range rl {
		t := []int{}
		for range cl {
			t = append(t, def)
		}
		g = append(g, t)
	}
	return g
}

type Uint16s [][]uint16

func (g Uint16s) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%d\n", item)
	}
	return out

}

func (in Uint16s) InBound(i, j int) bool {
	return !(i < 0 || i >= len(in) || j < 0 || j >= len(in[0]))
}

func (g Uint16s) Fill(rl, cl int, def uint16) Uint16s {
	g = Uint16s{}
	for range rl {
		t := []uint16{}
		for range cl {
			t = append(t, def)
		}
		g = append(g, t)
	}
	return g
}
