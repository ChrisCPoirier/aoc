package day4

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day4",
	Short: "day4",
	Long:  `day4`,
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

func part1(s string) int {
	score := 0
	for _, line := range strings.Split(s, "\n") {

		score += Score(line)
	}

	return score
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func Score(s string) int {

	sc := strings.Split(s, ": ")[1]

	wp := strings.Split(sc, " | ")

	wins := map[string]bool{}
	plays := map[string]bool{}

	//get wins
	for _, win := range strings.Split(wp[0], " ") {
		win = strings.Trim(win, " ")
		if win == "" {
			continue
		}

		wins[win] = true

	}

	//get play
	for _, play := range strings.Split(wp[1], " ") {
		play = strings.Trim(play, " ")
		if play == "" {
			continue
		}

		plays[play] = true
	}

	score := 0
	for k, _ := range plays {
		if ok := wins[k]; ok {
			if score == 0 {
				score = 1
				continue
			}
			score = score * 2
		}
	}

	return score
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
func part2(s string) int {
	t := map[int]int{1: 1}
	score := 0
	for _, line := range strings.Split(s, "\n") {
		id, score := Score2(line)

		if _, ok := t[id]; !ok {
			t[id] = 1
		}

		if score == 0 {
			continue
		}

		for j := id + 1; j <= id+score; j++ {
			if _, ok := t[j]; !ok {
				t[j] = 1
			}

			t[j] += t[id]
		}

	}

	for _, v := range t {
		score += v
	}

	return score
}

var spaceRE = regexp.MustCompile(`\s+`)

func Score2(s string) (int, int) {

	IdSc := strings.Split(s, ": ")

	// Card 1
	id, err := strconv.Atoi(strings.Split(spaceRE.ReplaceAllString(IdSc[0], " "), " ")[1])
	if err != nil {
		println(s)
		panic(err)
	}
	sc := IdSc[1]

	wp := strings.Split(sc, " | ")

	wins := map[string]bool{}
	plays := map[string]bool{}

	//get wins
	for _, win := range strings.Split(wp[0], " ") {
		win = strings.Trim(win, " ")
		if win == "" {
			continue
		}

		wins[win] = true

	}

	//get play
	for _, play := range strings.Split(wp[1], " ") {
		play = strings.Trim(play, " ")
		if play == "" {
			continue
		}

		plays[play] = true
	}

	score := 0
	for k, _ := range plays {
		if ok := wins[k]; ok {
			score += 1
		}
	}

	return id, score
}
