package year2015

import (
	"aoc/cmd/year2015/day1"
	"aoc/cmd/year2015/day2"
	"aoc/cmd/year2015/day3"
	"aoc/cmd/year2015/day4"
	"aoc/cmd/year2015/day5"

	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2015",
	Short: "2015 - Pew Pew",
	Long:  `2015 is a command line utility to super charge your Advent of Code experience`,
}

func init() {
	Cmd.AddCommand(day1.Cmd)
	Cmd.AddCommand(day2.Cmd)
	Cmd.AddCommand(day3.Cmd)
	Cmd.AddCommand(day4.Cmd)
	Cmd.AddCommand(day5.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
