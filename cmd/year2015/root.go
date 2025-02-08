package year2015

import (
	"aoc/cmd/year2015/day1"
	"aoc/cmd/year2015/day10"
	"aoc/cmd/year2015/day11"
	"aoc/cmd/year2015/day12"
	"aoc/cmd/year2015/day13"
	"aoc/cmd/year2015/day14"
	"aoc/cmd/year2015/day2"
	"aoc/cmd/year2015/day3"
	"aoc/cmd/year2015/day4"
	"aoc/cmd/year2015/day5"
	"aoc/cmd/year2015/day6"
	"aoc/cmd/year2015/day7"
	"aoc/cmd/year2015/day8"
	"aoc/cmd/year2015/day9"

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
	Cmd.AddCommand(day6.Cmd)
	Cmd.AddCommand(day7.Cmd)
	Cmd.AddCommand(day8.Cmd)
	Cmd.AddCommand(day9.Cmd)
	Cmd.AddCommand(day10.Cmd)
	Cmd.AddCommand(day11.Cmd)
	Cmd.AddCommand(day12.Cmd)
	Cmd.AddCommand(day13.Cmd)
	Cmd.AddCommand(day14.Cmd)
	//Cmd.AddCommand(day15.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
