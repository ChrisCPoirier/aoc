package day4

import (
	"aoc/cmd/common"
	"crypto/md5"
	"encoding/hex"
	"fmt"

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
	common.Run(parent, command, 1, part1, `part 1`)
	common.Run(parent, command, 1, part2, `part 2`)
}

func part1(s []byte) int {
	return findHashStartingWithNZeros(s, 5)
}

func part2(s []byte) int {
	return findHashStartingWithNZeros(s, 6)
}

func findHashStartingWithNZeros(in []byte, count int) int {
	inc := -1

	for {
		inc++
		out := md5.Sum([]byte(fmt.Sprintf("%s%d", in, inc)))

		if !startsWithNZeros(out, count) {
			continue
		}

		return inc

	}

	return inc
}

func startsWithNZeros(in [16]byte, count int) bool {
	for i := range min(len(in), count/2) {
		if in[i] != '\x00' {
			return false
		}
	}

	h := hex.EncodeToString(in[:])
	for _, v := range h[:min(len(h), count)] {
		if v != '0' {
			return false
		}
	}

	return true

}
