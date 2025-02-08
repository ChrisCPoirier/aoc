package grid

type Bools [][]bool

func (g Bools) Fill(rl, cl int, def bool) Bools {
	g = Bools{}
	for range rl {
		t := []bool{}
		for range cl {
			t = append(t, def)
		}
		g = append(g, t)
	}
	return g
}
