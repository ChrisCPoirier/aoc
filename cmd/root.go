package cmd

import (
	"fmt"
	"os"

	"aoc/cmd/year2015"
	year2022 "aoc/cmd/year2022"
	year2023 "aoc/cmd/year2023"
	"aoc/cmd/year2024"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Aoc - Pew Pew",
	Long:  `Aoc is a command line utility to super charge your Advent of Code experience`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(year2015.Cmd)
	rootCmd.AddCommand(year2022.Cmd)
	rootCmd.AddCommand(year2023.Cmd)
	rootCmd.AddCommand(year2024.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
