package year2022

import (
	"aoc/cmd/year2022/day2"
	"aoc/cmd/year2022/day3"
	"aoc/cmd/year2022/day4"
	"aoc/cmd/year2022/day5"
	"aoc/cmd/year2022/day6"
	"aoc/cmd/year2022/day7"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2022",
	Short: "2022 - Pew Pew",
	Long:  `2022 is a command line utility to super charge your Advent of Code experience`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	Cmd.AddCommand(day2.Cmd)
	Cmd.AddCommand(day3.Cmd)
	Cmd.AddCommand(day4.Cmd)
	Cmd.AddCommand(day5.Cmd)
	Cmd.AddCommand(day6.Cmd)
	Cmd.AddCommand(day7.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
