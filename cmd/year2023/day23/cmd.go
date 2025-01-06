package day23

import (
	"aoc/cmd/grid"
	"fmt"
	"maps"
	"os"
	"slices"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day23",
	Short: "day23",
	Long:  `day23`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(string(b), true))
	logrus.Infof("score part2: %d", part2(string(b), false))
	// logrus.Infof("score part2: %d", part2(string(b)))

}

type step struct {
	r, c    int
	path    [][]int
	visited map[string]bool
}

func (s step) Key() string {
	return fmt.Sprintf("%d:%d", s.r, s.c)
}

func (s step) Next(nr, nc int) step {
	return step{
		r:       nr,
		c:       nc,
		path:    append(slices.Clone(s.path), []int{nr, nc}),
		visited: maps.Clone(s.visited),
	}
}

var d = map[string][]int{
	"v": grid.DIR_DOWN,
	">": grid.DIR_RIGHT,
	"<": grid.DIR_LEFT,
	"^": grid.DIR_UP,
}

func part1(s string, forceDirection bool) int {
	return part2(s, forceDirection)
}

type edge struct {
	r, c, er, ec int
	cost         int
}

func (e edge) Key() string {

	sr, sc := e.r, e.c
	er, ec := e.er, e.ec

	//make sure they are ordered to ensure we do not revisit across duplicated edges
	if sr > er || (sr == er && sc > ec) {
		sr, sc, er, ec = er, ec, sr, sc
	}

	return fmt.Sprintf("%d:%d:%d:%d", sr, sc, er, ec)
}

type connection struct {
	tr, tc  int
	nodes   [][]int
	edges   map[string]edge
	visited map[string]bool
}

func (c connection) cost() int {
	cost := 0
	for _, e := range c.edges {
		cost += e.cost
	}
	//  - start/end
	return cost
}

func part2(s string, forceDirection bool) int {

	g := grid.New(s, ``)

	sr := 0
	sc := 1
	er := len(g) - 1
	ec := len(g[er-1]) - 2

	nodes := map[string][]int{grid.Key(sr, sc): {sr, sc}, grid.Key(er, ec): {er, ec}}

	//Find Junctions
	// Based on the test input and a review of the input data we have determinedc that this is a single width path maze.
	// Meaning that we can reduce our number of test points down to just the junctions (an point that splits into multiple directions)
	// we can then BFS between all junctions and generate a length of edges between each junction (or node)
	// Start and End are also nodes
	// This will reduce our complexity from size of grid to total edge size between S->E junctions
	// The route with longest route wins
	for r, row := range g {
		for c, v := range row {
			neighbors := 0

			if v != `.` {
				continue
			}

			for _, dir := range grid.DIR_CROSS {
				if !g.InBound(r+dir[0], c+dir[1]) {
					continue
				}

				if g[r+dir[0]][c+dir[1]] == `#` {
					continue
				}

				neighbors++
			}

			//if we have three or greater connections on a single point we are a junction/node
			if neighbors >= 3 {
				nodes[grid.Key(r, c)] = []int{r, c}
			}
		}
	}

	edges := []edge{}

	for _, n := range nodes {
		edges = append(edges, findEdges(g, n[0], n[1], nodes, forceDirection)...)
	}

	queue := []connection{}

	for _, e := range edges {
		if e.r == sr && e.c == sc {
			queue = append(queue, connection{tr: e.r, tc: e.c, nodes: [][]int{}, edges: map[string]edge{}, visited: map[string]bool{}})
		}
	}

	completed := []connection{}
	var q connection

	for len(queue) > 0 {
		q, queue = queue[0], queue[1:]

		if _, ok := q.visited[grid.Key(q.tr, q.tc)]; ok {
			continue
		}

		for _, e := range edges {
			if q.tr != e.r && q.tc != e.c {
				continue
			}

			nNodes := append(slices.Clone(q.nodes), []int{e.er, e.ec})
			nEdges := maps.Clone(q.edges)
			nEdges[e.Key()] = e

			nVisisted := maps.Clone(q.visited)

			nVisisted[grid.Key(e.r, e.c)] = true

			c := connection{tr: e.er, tc: e.ec, nodes: nNodes, edges: nEdges, visited: nVisisted}

			if e.er == er && e.ec == ec {
				completed = append(completed, c)
				continue
			}

			queue = append(queue, c)
		}
	}

	score := completed[0].cost()
	longest := completed[0]

	for _, complete := range completed {
		if complete.cost() > score {
			score = complete.cost()
			longest = complete
		}
		// score = max(score, complete.cost())
	}
	logrus.Infof("%#v", len(completed))
	logrus.Infof("%#v", longest.edges)
	logrus.Infof("%#v", longest.nodes)

	return score
}

func findEdges(g grid.Strings, sr, sc int, ends map[string][]int, forceDirection bool) []edge {
	edges := []edge{}
	queue := []step{{r: sr, c: sc, path: [][]int{{sr, sc}}, visited: map[string]bool{}}}

	var q step

	for len(queue) > 0 {
		q, queue = queue[0], queue[1:]

		if !g.InBound(q.r, q.c) {
			continue
		}

		if g[q.r][q.c] == `#` {
			continue
		}

		if ok := q.visited[q.Key()]; ok {
			continue
		}

		q.visited[q.Key()] = true

		if _, ok := ends[grid.Key(q.r, q.c)]; ok && sr != q.r && sc != q.c {
			edges = append(edges, edge{r: sr, c: sc, er: q.r, ec: q.c, cost: len(q.path) - 2})
			continue
		}

		if forceDirection && g[q.r][q.c] != `.` {
			if dir, ok := d[g[q.r][q.c]]; ok {
				next := q.Next(q.r+dir[0], q.c+dir[1])
				if ok := q.visited[next.Key()]; ok {
					continue
				}
				queue = append(queue, next)
			}
			continue
		}

		for _, dir := range grid.DIR_CROSS {
			if !g.InBound(q.r+dir[0], q.c+dir[1]) {
				continue
			}

			if g[q.r+dir[0]][q.c+dir[1]] == `#` {
				continue
			}

			next := q.Next(q.r+dir[0], q.c+dir[1])
			if ok := q.visited[next.Key()]; ok {
				continue
			}

			queue = append(queue, next)
		}
	}

	return edges

}
