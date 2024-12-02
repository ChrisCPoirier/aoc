package day1

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1",
	Long:  `day1`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

var words = map[string]byte{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	b2, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/2.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	b3 := bytes.Clone(b2)

	for k, v := range words {
		b3 = bytes.ReplaceAll(b3, []byte(k), []byte{v})
	}

	logrus.Infof("score part1: %d", part1(b))
	logrus.Infof("score part2: %d", part2(b2))
	// logrus.Infof("score part3: %d", part1(b3))

	// println(string(b3))
}

func part1(s []byte) int {
	score := 0
	for _, line := range bytes.Split(s, []byte("\n")) {
		score += getScoreInt(line)
	}

	return score
}

func part2(s []byte) int {

	score := 0
	for _, line := range bytes.Split(s, []byte("\n")) {
		score += getScoreIntAndWord(line)
	}

	return score
}

func getScoreIntAndWord(s []byte) int {
	var left, right byte
	l := len(s)

	for i, r := range s {
		if r >= '0' && r <= '9' {
			if left == 0 {
				left = r
			}
			right = r
			continue
		}

		switch r {
		case 'o', 't', 'f', 's', 'e', 'n':
			for k, v := range words {
				//one
				if r == k[0] {
					if i+len(k)-1 > l-1 {
						continue
					}

					if !bytes.Equal(s[i:i+len(k)], []byte(k)) {
						continue
					}

					if left == 0 {
						left = v
					}

					right = v
					break
				}
			}
		}
	}

	i, _ := strconv.Atoi(string(left) + string(right))

	return i
}

func getScoreInt(s []byte) int {
	var left, right byte

	for _, r := range s {
		if r >= '0' && r <= '9' {
			if left == 0 {
				left = r
			}
			right = r
		}
	}

	i, _ := strconv.Atoi(string(left) + string(right))

	return i
}

var replace = [][]string{
	{"one", "o1e"},
	{"two", "t2o"},
	{"three", "t3e"},
	{"four", "f4r"},
	{"five", "f5e"},
	{"six", "s6x"},
	{"seven", "s7n"},
	{"eight", "e8t"},
	{"nine", "n9e"},
}

func mutate(in []byte) []byte {
	for _, r := range replace {
		in = bytes.ReplaceAll(in, []byte(r[0]), []byte(r[1]))
	}

	in = onlyNum(in)

	return in
}

func onlyNum(s []byte) []byte {
	j := 0
	for _, b := range s {
		if '0' <= b && b <= '9' {
			s[j] = b
			j++
		}
	}
	return s[:j]
}
