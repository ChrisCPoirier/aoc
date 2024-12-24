package day23

import (
	"aoc/cmd/common"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day23",
	Long:  `day23`,
	Use:   "day23",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}

type node struct {
	name      string
	neighbors map[string]*node
}

func (n node) New(name string) *node {
	return &node{
		name:      name,
		neighbors: map[string]*node{},
	}
}

func part1(s []byte) int {
	nodes := map[string]*node{}

	lines := strings.Split(string(s), "\n")

	for _, line := range lines {
		nText := strings.Split(line, `-`)
		n1 := nText[0]
		n2 := nText[1]

		if _, ok := nodes[n1]; !ok {
			nodes[n1] = node{}.New(n1)
		}

		if _, ok := nodes[n2]; !ok {
			nodes[n2] = node{}.New(n2)
		}

		nodes[n1].neighbors[nodes[n2].name] = nodes[n2]
		nodes[n2].neighbors[nodes[n1].name] = nodes[n1]
	}

	networks := [][]*node{}
	for k, v := range nodes {
		if !strings.HasPrefix(k, `t`) {
			continue
		}

		for _, n1 := range v.neighbors {
			for _, n2 := range v.neighbors {
				for _, n3 := range n2.neighbors {
					if n3.name == v.name {
						continue
					}

					if n1.name == n3.name {
						networks = append(networks, []*node{v, n1, n2})
					}
				}
			}
		}
	}

	uniq := map[string][]*node{}
	for _, network := range networks {
		names := []string{network[0].name, network[1].name, network[2].name}
		slices.Sort(names)
		uniq[fmt.Sprintf("%#v", names)] = network
	}

	return len(uniq)
}

func part2(s []byte) string {
	nodes := map[string]*node{}

	lines := strings.Split(string(s), "\n")

	for _, line := range lines {
		nText := strings.Split(line, `-`)
		n1 := nText[0]
		n2 := nText[1]

		if _, ok := nodes[n1]; !ok {
			nodes[n1] = node{}.New(n1)
		}

		if _, ok := nodes[n2]; !ok {
			nodes[n2] = node{}.New(n2)
		}

		nodes[n1].neighbors[nodes[n2].name] = nodes[n2]
		nodes[n2].neighbors[nodes[n1].name] = nodes[n1]
	}

	networks := map[string]bool{}
	for _, v := range nodes {
		find(v, map[string]*node{v.name: v}, networks)
	}

	largest := ``

	for k := range networks {
		if len(k) > len(largest) {
			largest = k
		}
	}

	return largest
}

func find(n *node, matches map[string]*node, seen map[string]bool) {
	if _, ok := seen[key(matches)]; ok {
		return
	}

	seen[key(matches)] = true

	for _, neighbor := range n.neighbors {
		if _, ok := matches[neighbor.name]; ok {
			continue
		}

		if !hasAll(neighbor.neighbors, matches) {
			continue
		}

		copy := maps.Clone(matches)
		copy[neighbor.name] = neighbor
		find(neighbor, copy, seen)
	}
}

func key(nodes map[string]*node) string {
	names := []string{}

	for _, node := range nodes {
		names = append(names, node.name)
	}

	slices.Sort(names)

	return strings.Join(names, `,`)
}

func hasAll(neighbors, matches map[string]*node) bool {
	for k := range matches {
		if _, ok := neighbors[k]; !ok {
			return false
		}
	}
	return true
}
