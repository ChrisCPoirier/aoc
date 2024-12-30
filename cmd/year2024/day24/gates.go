package day24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type gate struct {
	name         string
	index, value int
	op           string
	set          bool
	a, b         *gate
}

func (g *gate) Set(v int) *gate {
	g.value = v
	g.set = true
	return g
}

func (g *gate) Compute() int {
	switch g.op {
	case `AND`:
		return g.a.value & g.b.value
	case `XOR`:
		return g.a.value ^ g.b.value
	case `OR`:
		return g.a.value | g.b.value
	default:
		logrus.Fatalf("unknown op: %s", g.op)
	}
	return 0
}

type gates map[string]*gate

func initGates(s string) gates {
	g := gates{}
	for _, line := range strings.Split(s, "\n") {
		items := strings.Split(line, `: `)

		name := items[0]
		v, err := strconv.Atoi(items[1])

		if err != nil {
			logrus.Fatal(err)
		}

		g.Add(name).Set(v)
	}
	return g
}

func getGateInstructions(s string) ([]string, []string, []string, []string) {
	a := []string{}
	op := []string{}
	b := []string{}
	dest := []string{}

	for _, line := range strings.Split(s, "\n") {
		items := strings.Split(line, ` `)
		a = append(a, items[0])
		op = append(op, items[1])
		b = append(b, items[2])
		dest = append(dest, items[4])
	}
	return a, op, b, dest
}

func NewGates(s []byte) gates {
	sections := strings.Split(string(s), "\n\n")
	gates := initGates(sections[0])

	aGates, op, bGates, destGates := getGateInstructions(sections[1])

	for i := range destGates {
		d := gates.Add(destGates[i])
		a := gates.Add(aGates[i])
		b := gates.Add(bGates[i])
		d.a = a
		d.b = b

		d.op = op[i]
	}

	return gates
}

func (g gates) Add(s string) *gate {
	if v, ok := g[s]; ok {
		return v
	}

	id := len(g)

	g[s] = &gate{name: s, index: id}
	return g[s]
}

func (g gates) Swap(a, b *gate) {
	a.index, b.index, a.a, a.b, b.a, b.b = b.index, a.index, b.a, b.b, a.a, a.b
}

func (g gates) Result() string {
	zGates := []string{}

	for _, gate := range g {
		if !strings.HasPrefix(gate.name, `z`) {
			continue
		}
		zGates = append(zGates, gate.name)
	}

	slices.Sort(zGates)
	slices.Reverse(zGates)

	out := ""
	for _, gate := range zGates {
		out += fmt.Sprintf("%b", g[gate].value)
	}
	return out
}

func (g gates) Solve() error {
	queue := []*gate{}
	for _, v := range g {
		queue = append(queue, v)
	}

	var q *gate
	itterSinceProcessed := 0
	for len(queue) > 0 {
		q, queue = queue[0], queue[1:]
		if strings.HasPrefix(q.name, `x`) || strings.HasPrefix(q.name, `y`) {
			continue
		}

		if q.a == nil || q.b == nil {
			return fmt.Errorf("unsolvable configuration")
		}

		//if one of the values has not been set, put it on the back of the stack
		if !q.a.set || !q.b.set {
			itterSinceProcessed++
			if itterSinceProcessed > len(queue)+10 {
				return fmt.Errorf("endless loop condition met %d", itterSinceProcessed)
			}
			queue = append(queue, q)
			continue
		}
		itterSinceProcessed = 0
		v := q.Compute()

		q.Set(v)
	}

	return nil
}

func (g gates) FindGateBySources(a, b, op string) *gate {
	for _, gate := range g {
		if gate.a == nil || gate.b == nil {
			continue
		}

		if gate.op != op {
			continue
		}

		if gate.a.name == a && gate.b.name == b ||
			gate.b.name == a && gate.a.name == b {
			return gate
		}
	}
	return nil
}

func (g gates) FindGateBySource(a, op string) *gate {
	for _, gate := range g {
		if gate.a == nil || gate.b == nil {
			continue
		}

		if gate.op != op {
			continue
		}

		if gate.a.name == a ||
			gate.b.name == a {
			return gate
		}
	}
	return nil
}

func (g gates) Fix(gt *gate) []string {
	// logrus.Infof("fixing gate: %s", gt.name)
	if gt.name == `z00` {
		if !((gt.a.name == `x00` && gt.b.name == `y00`) ||
			(gt.a.name == `y00` && gt.b.name == `x00`)) || gt.op != `XOR` {
			sg := g.FindGateBySources(`x00`, `y00`, `XOR`)
			g.Swap(gt, sg)

			return []string{gt.name, sg.name}
		}
		return []string{}
	}

	if gt.name == `z45` {
		if gt.op != `OR` {
			return []string{gt.name}
		}
		return []string{}
	}

	// if the start operator is incorrect, swap it out for the correct one
	// needs to attmept find and fix for htis condition

	if gt.op != `XOR` {
		og := g.FindGateBySources(fmt.Sprintf(`x%s`, gt.name[1:]), fmt.Sprintf(`y%s`, gt.name[1:]), `XOR`)
		sg := g.FindGateBySource(og.name, `XOR`)
		g.Swap(gt, sg)

		return []string{gt.name, sg.name}
	}

	og := g.FindGateBySources(fmt.Sprintf(`x%s`, gt.name[1:]), fmt.Sprintf(`y%s`, gt.name[1:]), `XOR`)
	sg := g.FindGateBySource(og.name, `XOR`)

	// We could not locate the XOR value
	if sg == nil {
		// we know the previous gate is accurate, use that to build the next gates tree
		gateNum, _ := strconv.Atoi(gt.name[1:])
		// -- to get previous gate num
		gateNum--
		prevGate := g[fmt.Sprintf("z%0d", gateNum)]

		//The carry bit from previous number
		c1 := g.FindGateBySources(prevGate.a.name, prevGate.b.name, `AND`)
		c2 := g.FindGateBySource(c1.name, `OR`)

		actualGate := g.FindGateBySource(c2.name, `XOR`)

		if actualGate != gt {
			g.Swap(actualGate, gt)
			return []string{actualGate.name, gt.name}
		}

		incorrectGate := gt.a

		if gt.a == c2 {
			incorrectGate = gt.b
		}

		g.Swap(og, incorrectGate)
		return []string{og.name, incorrectGate.name}

	}

	return []string{}

}
