package year2024

import (
	"aoc/cmd/year2024/day1"
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
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
