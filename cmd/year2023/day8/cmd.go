package day8

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day8",
	Short: "day8",
	Long:  `day8`,
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
	// logrus.Infof("score part2: %d", part2(string(b)))
	logrus.Infof("score part3: %d", part3(string(b)))
}

// LLR

// AAA = (BBB, BBB)
// BBB = (AAA, ZZZ)
// ZZZ = (ZZZ, ZZZ)

var items = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func part1(s string) int64 {
	var score int64 = 0

	lines := strings.Split(s, "\n\n")

	m := map[string][]string{}

	for _, line := range strings.Split(lines[1], "\n") {
		matches := items.FindStringSubmatch(line)

		m[matches[1]] = []string{matches[2], matches[3]}
	}

	local := "AAA"

	for local != `ZZZ` {
		for _, d := range lines[0] {
			score++
			if d == 'R' {
				local = m[local][1]
			} else {
				local = m[local][0]
			}

			if local == `ZZZ` {
				break
			}
		}
	}

	return score
}

func part2(s string) int64 {
	var score int64 = 0

	lines := strings.Split(s, "\n\n")

	m := map[string][]string{}

	for _, line := range strings.Split(lines[1], "\n") {
		matches := items.FindStringSubmatch(line)
		m[matches[1]] = []string{matches[2], matches[3]}
	}

	locals := []string{}

	for k := range m {
		if strings.Contains(k, `A`) {
			locals = append(locals, k)
		}
	}

	for {
		for _, d := range lines[0] {
			score++
			for i, local := range locals {

				if d == 'R' {
					locals[i] = m[local][1]
				} else {
					locals[i] = m[local][0]
				}
			}

			fmt.Printf("current score: %d\n", score)

			if allOnZ(locals) {
				return score
			}
		}
	}

	// return score
}

type node struct {
	Name      string
	leftNode  *node
	rightNode *node
	z         bool
	a         bool
}

func part3(s string) int64 {
	lines := strings.Split(s, "\n\n")

	m := map[string][]string{}

	for _, line := range strings.Split(lines[1], "\n") {
		matches := items.FindStringSubmatch(line)
		m[matches[1]] = []string{matches[2], matches[3]}
	}

	nodes := map[string]*node{}

	for k := range m {
		nodes[k] = &node{Name: k, z: strings.Contains(k, `Z`), a: strings.Contains(k, `A`)}
	}

	for k := range nodes {
		nodes[k].leftNode = nodes[m[k][0]]
		nodes[k].rightNode = nodes[m[k][1]]
	}

	activeNodes := []*node{}

	for k, v := range nodes {
		if strings.Contains(k, `A`) {
			activeNodes = append(activeNodes, v)
		}
	}

	iterations := [][]int{}
	for _, n := range activeNodes {
		dir := lines[0]
		iteration := []int{}
		count := 0

		var first *node

		for {

			for count == 0 || !strings.Contains(n.Name, `Z`) {
				count += 1

				if dir[0] == 'R' {
					n = n.rightNode
				} else {
					n = n.leftNode
				}

				dir = dir[1:] + string(dir[0])

			}

			iteration = append(iteration, count)

			if first == nil {
				first = n
				count = 0
			} else if n == first {
				break
			}

		}
		iterations = append(iterations, iteration)

	}

	nums := []int{}

	for _, iteration := range iterations {
		nums = append(nums, iteration[0])
	}

	lcm := nums[0]

	for _, n := range nums {
		lcm = lcm * n / gcd(lcm, n)
	}

	return int64(lcm)
}

func allOnZ(locals []string) bool {
	for _, local := range locals {
		if !strings.Contains(local, `Z`) {
			return false
		}
	}

	return true
}

func allOnZNode(nodes []*node) bool {
	for _, n := range nodes {
		if !n.z {
			return false
		}
	}

	return true
}

func allOnANode(nodes []*node) bool {
	for _, n := range nodes {
		if !n.a {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
