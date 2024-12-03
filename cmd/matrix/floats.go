package matrix

import "fmt"

type Floats [][]float64

func (g Floats) Pretty() string {
	out := ``
	for _, item := range g {
		out += fmt.Sprintf("%.0f\n", item)
	}
	return out

}
