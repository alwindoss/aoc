/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
	},
}

var (
	year   int
	day    int
	puzzle int
)

func init() {
	rootCmd.AddCommand(runCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")
	runCmd.PersistentFlags().IntVarP(&year, "year", "y", 0, "--year 2023")
	runCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "--day 1")
	runCmd.PersistentFlags().IntVarP(&puzzle, "puzzle", "p", 0, "--puzzle 2")
	runCmd.MarkFlagRequired("year")
	runCmd.MarkFlagRequired("day")
	runCmd.MarkFlagRequired("puzzle")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
