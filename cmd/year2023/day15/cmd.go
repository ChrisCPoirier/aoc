package day15

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day15",
	Short: "day15",
	Long:  `day15`,
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
	var score int = 0

	for _, seq := range strings.Split(s, `,`) {

		score += hash(seq)
	}

	return int64(score)
}

type lens struct {
	label string
	fl    int
}

func part2(s string) int64 {
	var score int = 0

	m := map[int][]lens{}
	for _, seq := range strings.Split(s, `,`) {
		minus := strings.Index(seq, `-`)
		if minus > -1 {
			label := seq[:minus]
			h := hash(label)

			if v, ok := m[h]; ok {
				for i, lens := range v {
					if lens.label == label {
						m[h] = slices.Delete(v, i, i+1)
						break
					}
				}
			}
			continue
		}

		items := strings.Split(seq, `=`)
		label := items[0]
		fl, err := strconv.Atoi(items[1])

		if err != nil {
			panic(err)
		}

		h := hash(label)
		if v, ok := m[h]; ok {
			found := false
			for i, l := range v {
				if l.label == label {
					found = true
					v[i] = lens{label: label, fl: fl}
					m[h] = v
					break
				}
			}
			if !found {
				m[h] = append(v, lens{label: label, fl: fl})
			}
		} else {
			m[h] = []lens{{label: label, fl: fl}}
		}
	}

	fmt.Printf("%#v\n", m)
	for k, lenses := range m {
		for i, lens := range lenses {
			score += (k + 1) * (i + 1) * lens.fl
		}
	}

	return int64(score)
}

func hash(s string) int {
	ht := 0
	for _, c := range s {
		ht = (ht + int(c)) * 17 % 256
	}
	return ht
}
