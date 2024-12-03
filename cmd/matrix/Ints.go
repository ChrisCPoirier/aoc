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
