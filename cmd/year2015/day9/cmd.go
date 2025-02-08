package day9

import (
	"aoc/cmd/common"
	"container/heap"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day9",
	Short: "day9",
	Long:  `day9`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

type edge struct {
	cost int
	dest string
}

func part1(s []byte) int {
	nodes := map[string][]edge{}

	score := 0
	for _, line := range strings.Split(string(s), "\n") {
		items := strings.Split(line, " ")

		v, _ := strconv.Atoi(items[4])

		nodes[items[0]] = append(nodes[items[0]], edge{dest: items[2], cost: v})

		//ensure the dest node is in our list of nodes
		nodes[items[2]] = append(nodes[items[2]], edge{dest: items[0], cost: v})
	}

	pq := &shortestPath{}
	heap.Init(pq)

	for node := range nodes {
		heap.Push(pq, &path{
			name: node,
			cost: 0,
			path: []string{node},
		})
	}

	for pq.Len() > 0 {
		p := heap.Pop(pq).(*path)

		//return the first path that traversed all nodes.
		// Since we are traversing in least cost order, this is also shortest path
		if len(p.path) == len(nodes) {
			return p.cost
		}

		for _, edge := range nodes[p.name] {
			//do not visist a node we have visisted before
			if slices.Contains(p.path, edge.dest) {
				continue
			}

			heap.Push(pq, &path{
				name: edge.dest,
				cost: p.cost + edge.cost,
				path: append(slices.Clone(p.path), edge.dest),
			})
		}
	}

	return score
}

func part2(s []byte) int {
	nodes := map[string][]edge{}

	for _, line := range strings.Split(string(s), "\n") {
		items := strings.Split(line, " ")

		v, _ := strconv.Atoi(items[4])

		nodes[items[0]] = append(nodes[items[0]], edge{dest: items[2], cost: v})

		//ensure the dest node is in our list of nodes
		nodes[items[2]] = append(nodes[items[2]], edge{dest: items[0], cost: v})
	}

	pq := &shortestPath{}
	heap.Init(pq)

	for node := range nodes {
		heap.Push(pq, &path{
			name: node,
			cost: 0,
			path: []string{node},
		})
	}

	logrus.Infof("locations: %#v", strings.Join(slices.Collect(maps.Keys(nodes)), `,`))
	logrus.Infof("%#v", nodes)

	longest := path{}
	for pq.Len() > 0 {
		p := heap.Pop(pq).(*path)

		//return the first path that traversed all nodes.
		// Since we are traversing in least cost order, this is also shortest path
		if len(p.path) == len(nodes) {
			logrus.Infof("found path: cost: %d, path: %s", longest.cost, strings.Join(longest.path, ","))
			if p.cost > longest.cost {
				longest = *p
			}
			continue
		}

		for _, edge := range nodes[p.name] {
			//do not visist a node we have visisted before
			if slices.Contains(p.path, edge.dest) {
				continue
			}

			heap.Push(pq, &path{
				name: edge.dest,
				cost: p.cost + edge.cost,
				path: append(slices.Clone(p.path), edge.dest),
			})
		}
	}

	logrus.Infof("longest - end: %s cost: %d, path: %s", longest.name, longest.cost, strings.Join(longest.path, ","))
	return longest.cost
}

type path struct {
	name  string
	cost  int
	index int
	path  []string
}

type shortestPath []*path

func (pq shortestPath) Len() int { return len(pq) }

func (pq shortestPath) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq shortestPath) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *shortestPath) Push(x any) {
	n := len(*pq)
	path := x.(*path)
	path.index = n
	*pq = append(*pq, path)
}

func (pq *shortestPath) Pop() any {
	old := *pq
	n := len(old)
	path := old[n-1]
	old[n-1] = nil
	path.index = -1
	*pq = old[0 : n-1]
	return path
}
