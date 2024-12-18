package day17

import (
	"aoc/cmd/common"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day17",
	Long:  `day17`,
	Use:   "day17",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1, "part 1")
	common.Run(parent, command, 1, part2, "part 2")
}
func key(r, c int) string {
	return fmt.Sprintf("%d:%d", r, c)
}

var mem = map[string]int{}
var reRegister = regexp.MustCompile(`Register (\w): (\d+)`)

var reProgram = regexp.MustCompile(`Program: (.*)$`)

func part1(s []byte) string {
	logrus.SetLevel(logrus.DebugLevel)
	for _, match := range reRegister.FindAllStringSubmatch(string(s), len(s)) {
		v, _ := strconv.Atoi(match[2])
		mem[match[1]] = v
	}

	out := run(s)
	return strings.Join(strings.Split(out, ``), `,`)
}

func part2(s []byte) int {
	logrus.SetLevel(logrus.InfoLevel)
	for _, match := range reRegister.FindAllStringSubmatch(string(s), len(s)) {
		v, _ := strconv.Atoi(match[2])
		mem[match[1]] = v
	}
	program := strings.Split(reProgram.FindAllStringSubmatch(string(s), len(s))[0][1], `,`)

	q := []int{}
	new_q := []int{0}

	logrus.Infof("PROGRAM: %s", program)

	for i := len(program) - 1; i >= 0; i-- {
		q = slices.Clone(new_q)
		new_q = []int{}
		for _, n := range q {
			n = n << 3
			for j := range 8 {
				mem[`A`] = n + j
				mem[`B`] = 0
				mem[`C`] = 0
				out := run(s)

				if string(out[0]) == program[i] {
					logrus.Infof("MEM_A_%d: %b (%d)", i, n+j, n+j)
					new_q = append(new_q, n+j)
				}

				if out == strings.Join(program, ``) {
					logrus.Infof("match %d", n+j)
					return n + j
				}
			}
		}
	}
	logrus.Infof("new queue: %#v", new_q)
	logrus.Infof("final queue: %#v", q)
	return 0
}

func run(s []byte) string {

	program := strings.Split(reProgram.FindAllStringSubmatch(string(s), len(s))[0][1], `,`)

	out := []string{}

	logrus.Debugf("MEM_A_0: %b (%d)", mem[`A`], mem[`A`])

	for i := 0; i < len(program); i++ {
		// for i, v := range program {
		v := program[i]
		if i%2 == 0 {
			continue
		}

		if program[i-1] == `3` {
			if mem[`A`] == 0 {
				continue
			}
			vi, _ := strconv.Atoi(v)
			i = vi - 1
			continue
		}

		fn := getOp(program[i-1])
		o := fn(v)
		if o != `` {
			out = append(out, o)
			logrus.Debugf("MEM_A_%d: %b (%d)", len(out), mem[`A`], mem[`A`])
		}
	}

	return strings.Join(out, ``)
}

func getOp(s string) func(string) string {
	switch s {
	case `0`:
		return adv
	case `5`:
		return out
	case `2`:
		return bst
	case `1`:
		return bxl
	case `7`:
		return cdv
	case `4`:
		return bxc
	}

	logrus.Fatalf("opcode not defined %s", s)
	return nil
}

func getValue(s string) int {
	switch s {
	case `4`:
		return mem[`A`]
	case `5`:
		return mem[`B`]
	case `6`:
		return mem[`C`]
	case `7`:
		logrus.Fatal(`Unknown condition occured`)

	}
	v, _ := strconv.Atoi(s)
	return v
}

func adv(s string) string {
	in := getValue(s)
	num := mem[`A`]
	den := int(math.Pow(2, float64(in)))
	mem[`A`] = num / den
	return ``
}

func out(s string) string {
	in := getValue(s)
	v := in % 8
	return strings.Join(strings.Split(fmt.Sprintf("%d", v), ``), `,`)
}

func bst(s string) string {
	in := getValue(s)
	v := in % 8
	mem[`B`] = v
	return ``
}

func bxl(s string) string {
	in, _ := strconv.Atoi(s)
	mem[`B`] = mem[`B`] ^ in
	return ``
}

// The bxc instruction (opcode 4)
// calculates the bitwise XOR of register B and register C,
// then stores the result in register B.
// (For legacy reasons, this instruction reads an operand but ignores it.)
func bxc(in string) string {
	mem[`B`] = mem[`B`] ^ mem[`C`]
	return ``
}

// The cdv instruction (opcode 7) works exactly like the
// adv instruction except that the result is stored in the C register.
// ÷ (The numerator is still read from the A register.)
func cdv(s string) string {
	in := getValue(s)
	num := mem[`A`]
	den := int(math.Pow(2, float64(in)))
	if den == 0 {
		mem[`C`] = 0
		return ``
	}
	mem[`C`] = num / den
	return ``
}
