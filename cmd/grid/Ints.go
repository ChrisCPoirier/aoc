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

type Floats [][]float64

func (g Floats) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%.0f\n", item)
	}
	return out

}
