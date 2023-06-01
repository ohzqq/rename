package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/spf13/cobra"
)

// numberCmd represents the numbering command
var numberCmd = &cobra.Command{
	Use:   "number zeroes glob|file...",
	Short: "number the files",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cfg.SetZeroes(args[0])
		View(args[1:])
	},
}

func init() {
	rootCmd.AddCommand(numberCmd)
}
