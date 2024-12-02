package day5

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day5",
	Short: "day5",
	Long:  `day5`,
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
	logrus.Infof("score part3: %d", part3(string(b)))
}

func buildMaps(s string) ([]int64, map[string][][]int64) {

	seeds := []int64{}
	maps := map[string][][]int64{}

	for _, line := range strings.Split(s, "\n\n") {
		// fmt.Printf("split: %s\n", line)

		if strings.Contains(line, "seeds:") {
			seeds = getInts(strings.ReplaceAll(line, "seeds: ", ""))
			// fmt.Printf("seeds: %#v", seeds)
			continue
		}

		mappings := strings.Split(line, "\n")

		name := strings.Split(mappings[0], " ")[0]

		for i := 1; i < len(mappings); i++ {
			maps[name] = append(maps[name], getInts(mappings[i]))
		}

		// score += Score(line)
	}
	return seeds, maps
}

func part1(s string) int64 {
	var score int64

	seeds, maps := buildMaps(s)
	for _, seed := range seeds {
		ts := Score(seed, maps)
		if score == 0 {
			score = ts
			continue
		}

		if ts < score {
			score = ts
		}

	}

	return score
}

func part2(s string) int64 {
	var score int64 = -1

	seeds, maps := buildMaps(s)

	for i := 0; i < len(seeds); i += 2 {
		// fmt.Printf("Processing seed range: %d,%d\n", seeds[i], seeds[i]+seeds[i+1]-1)
		for j := seeds[i]; j < (seeds[i] + seeds[i+1]); j++ {
			ts := Score(j, maps)
			if score == -1 {
				score = ts
				continue
			}

			if ts < score {
				score = ts
			}
		}
	}

	return score
}

func getInts(s string) []int64 {
	// 79 14 55 13
	ints := []int64{}
	for _, v := range strings.Split(s, " ") {
		i, err := strconv.ParseInt(v, 10, 64)

		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}

	return ints
}

// Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
// soil-to-fertilizer
// fertilizer-to-water
// water-to-light
// light-to-temperature
// temperature-to-humidity
// humidity-to-location
// 1-10,  3-4,7-9
func Score(seed int64, maps map[string][][]int64) int64 {
	soil := getDest(seed, maps[`seed-to-soil`])
	fertilizer := getDest(soil, maps[`soil-to-fertilizer`])
	water := getDest(fertilizer, maps[`fertilizer-to-water`])
	light := getDest(water, maps[`water-to-light`])
	temp := getDest(light, maps[`light-to-temperature`])
	humidity := getDest(temp, maps[`temperature-to-humidity`])
	location := getDest(humidity, maps[`humidity-to-location`])
	return location
}

func getDest(source int64, mapping [][]int64) int64 {
	for _, m := range mapping {
		if m[1] <= source && source < m[1]+m[2] {
			return m[0] + (source - m[1])
		}
	}
	return source
}

// SortMaps in decending order
func sortMaps(maps map[string][][]int64) map[string][][]int64 {
	for k, m := range maps {
		sort.Slice(m,
			func(i, j int) bool {

				if m[i][1] == m[j][1] {
					return m[i][2] < m[j][2]
				}

				return m[i][1] < m[j][1]
			},
		)
		maps[k] = m
	}

	return maps
}

func getDestRange(sources [][]int64, mapping [][]int64) [][]int64 {
	destinations := [][]int64{}
	for _, m := range mapping {
		//work backwards since we are delete out of the slice
		for i := len(sources) - 1; i >= 0; i-- {
			source := sources[i]

			//Since the maps are in order, if the start of the source is less than we know all are straight mappings
			if source[1] < m[1] {
				destinations = append(destinations, source)
				//If we enter this block, since are mapping is ordered, we know we have nothing more to do, return the desitnations
				return destinations
			}

			//handle anything outside of left side - straight mappings
			if source[0] < m[1] {
				destinations = append(destinations, [][]int64{{source[0], m[1] - 1}}...)
				sources[i][0] = m[1]
			}

			// m[0] + (source - m[1])
			//Map whats left with offsets
			if m[1] <= source[0] && source[1] < m[1]+m[2] {
				var start int64 = m[0] + (source[0] - m[1])
				var end int64 = start + (source[1] - source[0])
				destinations = append(destinations, [][]int64{{start, end}}...)
				sources = slices.Delete(sources, i, i+1)
			}

			//Only some of the data is within bounds and some pushes outside our right bound
			if source[0] < m[1]+m[2] && source[1] > m[1]+m[2] {
				var start int64 = m[0] + (source[0] - m[1])
				var end int64 = (m[1] + m[2] - 1) - (m[1] - m[0])
				destinations = append(destinations, [][]int64{{start, end}}...)
				sources[i][0] = m[1] + m[2]
			}
		}
	}

	//anything left is above our ranges of our maps and is added to the stack matching the 1 for 1 rule if missing
	destinations = append(destinations, sources...)

	return destinations
}

func part3(s string) int64 {
	var score int64 = -1

	seeds, maps := buildMaps(s)

	maps = sortMaps(maps)

	locationRange := [][]int64{}

	for i := 0; i < len(seeds); i += 2 {
		// fmt.Printf("Processing seed range: %d,%d\n", seeds[i], seeds[i]+seeds[i+1]-1)

		nlocation := ScoreRange([][]int64{{seeds[i], seeds[i] + seeds[i+1] - 1}}, maps)
		locationRange = append(locationRange, nlocation...)
	}

	for _, locations := range locationRange {
		for _, location := range locations {
			if score == -1 {
				score = location
				continue
			}
			if location < score {
				score = location
			}
		}
	}

	return score
}

func ScoreRange(seeds [][]int64, maps map[string][][]int64) [][]int64 {
	soil := getDestRange(seeds, maps[`seed-to-soil`])
	fertilizer := getDestRange(soil, maps[`soil-to-fertilizer`])
	water := getDestRange(fertilizer, maps[`fertilizer-to-water`])
	light := getDestRange(water, maps[`water-to-light`])
	temp := getDestRange(light, maps[`light-to-temperature`])
	humidity := getDestRange(temp, maps[`temperature-to-humidity`])
	location := getDestRange(humidity, maps[`humidity-to-location`])
	return location
}
