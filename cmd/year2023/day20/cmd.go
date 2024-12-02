package day20

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day20",
	Short: "day20",
	Long:  `day20`,
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

type moduleQueue []module

func (m moduleQueue) Push(in module) moduleQueue {
	m = append(m, in)
	return m
}

func (m moduleQueue) Pop() (moduleQueue, module) {
	out := m[0]
	m = m[1:]
	return m, out
}

type notify struct {
	name string
	high bool
}

func part1(s string) int64 {
	// var score int = 0

	notifyChan := make(chan notify)

	modules := parse(s, notifyChan)

	low, high := 0, 0

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		for sig := range notifyChan {
			if sig.high {
				high++
				continue
			}
			low++
		}
		wg.Done()
	}(wg)

	var m module

	for i := 0; i < 1000; i++ {
		notifyChan <- notify{`button push`, false}
		q := moduleQueue{modules[`broadcaster`]}

		for len(q) > 0 {
			q, m = q.Pop()

			for _, destination := range m.pulse() {
				q = q.Push(destination)
			}
		}
	}

	close(notifyChan)

	wg.Wait()

	return int64(low * high)
}

var search string = `rx`

func part2(s string) int64 {
	// var score int = 0

	notifyChan := make(chan notify)

	modules := parse(s, notifyChan)

	low, high := 0, 0

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		for sig := range notifyChan {
			// fmt.Printf("%s sent %t\n", sig.name, sig.high)
			if sig.high {
				high++
				continue
			}
			low++
		}
		wg.Done()
	}(wg)

	var m module
	buttonPresses := 0

	cycles := map[string]int{}

	triggerModules := []string{}
	for _, m := range modules {
		for _, dest := range m.Destinations() {
			if dest.Name() == search {
				triggerModules = append(triggerModules, m.Name())
				break
			}
		}
	}

	for _, m := range modules {
		for _, dest := range m.Destinations() {
			if slices.Contains(triggerModules, dest.Name()) {
				cycles[m.Name()] = 0
				break
			}
		}
	}

OUTER:
	for {
		buttonPresses++
		notifyChan <- notify{`button push`, false}
		q := moduleQueue{modules[`broadcaster`]}

		for len(q) > 0 {
			q, m = q.Pop()

			if v, ok := cycles[m.Name()]; ok {
				if v == 0 {
					if m.willSend() {
						cycles[m.Name()] = buttonPresses
					}
				}
			}

			cyclesFound := true
			for _, v := range cycles {
				if v == 0 {
					cyclesFound = false
					break
				}
			}

			if cyclesFound {
				break OUTER
			}

			for _, destination := range m.pulse() {
				q = q.Push(destination)
			}
		}
	}

	close(notifyChan)

	wg.Wait()

	lcm := 0

	fmt.Printf("cycles :%#v\n", cycles)
	for _, v := range cycles {
		if lcm == 0 {
			lcm = v
		}
		lcm = lcm * v / gcd(lcm, v)
	}

	return int64(lcm)
}

func parse(s string, notifyChan chan notify) map[string]module {
	modules := map[string]module{}

	//generate modules
	for _, line := range strings.Split(s, "\n") {
		items := strings.Split(line, ` -> `)
		name := items[0][1:]

		if items[0] == `broadcaster` {
			name = `broadcaster`
		}

		switch items[0][0] {
		case '%':
			modules[name] = &flipFlop{name: name, notify: notifyChan}
		case '&':
			modules[name] = &conjunction{name: name, notify: notifyChan}
		default:
			modules[name] = &broadcaster{name: name, notify: notifyChan}
		}
	}

	//associate modules
	for _, line := range strings.Split(s, "\n") {
		items := strings.Split(line, ` -> `)
		name := items[0][1:]

		if items[0] == `broadcaster` {
			name = `broadcaster`
		}

		destinations := strings.Split(items[1], ", ")
		for _, destination := range destinations {
			m := modules[name]
			destM := modules[destination]

			if destM == nil {
				destM = &base{name: destination}
				modules[destination] = destM
			}

			m.addDestination(destM)
			destM.addSource(m)
		}
	}

	return modules
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
