package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func AsInts(items []string) []int {
	ints := []int{}
	for _, s := range items {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}

	return ints
}

func AsFloats(items []string) []float64 {
	floats := []float64{}
	for _, s := range items {
		i, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		floats = append(floats, i)
	}

	return floats
}

func Chunk(items []string, d string) [][]string {
	chunks := [][]string{}
	start := 0

	for i, v := range items {
		if v == d {
			chunks = append(chunks, items[start:i])
			start = i + 1
		} else if i == len(items)-1 {
			chunks = append(chunks, items[start:i+1])
		}
	}

	if items[len(items)-1] == d {
		chunks = append(chunks, []string{})
	}

	return chunks
}

func Stitch(items [][]string, d string) []string {
	stiched := []string{}

	for i, item := range items {
		stiched = append(stiched, item...)

		if i < len(items)-1 {
			stiched = append(stiched, `#`)
		}
	}

	return stiched
}

func Counts[E comparable](in []E) map[E]int {
	out := map[E]int{}

	for _, a := range in {
		out[a] = out[a] + 1
	}

	return out
}

func Index[T comparable](in []T) map[T]int {
	index := map[T]int{}
	for i, c := range in {
		index[c] = i
	}
	return index
}

func Uniq[T comparable](in [][]T) [][]T {
	k := map[string][]T{}
	for _, items := range in {
		k[fmt.Sprintf("%#v", items)] = items
	}

	u := [][]T{}

	for _, v := range k {
		u = append(u, v)
	}

	return u
}

func Run[T comparable](year, day string, part int, fn func([]byte) T, message ...string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/%d.txt`, year, day, part))

	if err != nil {
		logrus.Fatal(err)
	}

	m := strings.Join(message, ``)
	result := fn(b)
	switch v := any(result).(type) {
	case string:
		logrus.Infof("score part %d%s: %s", part, m, v)
	case float64, float32:
		logrus.Infof("score part %d%s: %.0f", part, m, v)
	case int:
		logrus.Infof("score part %d%s: %d", part, m, v)
	default:
		logrus.Infof("score part %d%s: %#v", part, m, v)
	}

}
