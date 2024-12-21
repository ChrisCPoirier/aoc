package year2024

import (
	"aoc/cmd/year2024/day1"
	"aoc/cmd/year2024/day10"
	"aoc/cmd/year2024/day11"
	"aoc/cmd/year2024/day12"
	"aoc/cmd/year2024/day13"
	"aoc/cmd/year2024/day14"
	"aoc/cmd/year2024/day15"
	"aoc/cmd/year2024/day16"
	"aoc/cmd/year2024/day17"
	"aoc/cmd/year2024/day18"
	"aoc/cmd/year2024/day19"
	"aoc/cmd/year2024/day2"
	"aoc/cmd/year2024/day20"
	"aoc/cmd/year2024/day21"
	"aoc/cmd/year2024/day3"
	"aoc/cmd/year2024/day4"
	"aoc/cmd/year2024/day5"
	"aoc/cmd/year2024/day6"
	"aoc/cmd/year2024/day7"
	"aoc/cmd/year2024/day8"
	"aoc/cmd/year2024/day9"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2024",
	Short: "2024 - Pew Pew",
	Long:  `2024 is a command line utility to super charge your Advent of Code experience`,
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
	Cmd.AddCommand(day15.Cmd)
	Cmd.AddCommand(day16.Cmd)
	Cmd.AddCommand(day17.Cmd)
	Cmd.AddCommand(day18.Cmd)
	Cmd.AddCommand(day19.Cmd)
	Cmd.AddCommand(day20.Cmd)
	Cmd.AddCommand(day21.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
