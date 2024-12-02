package common

import (
	"strconv"
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
