package matrix

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
