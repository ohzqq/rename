package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/spf13/cobra"
)

// lowerCmd represents the lower command
var lowerCmd = &cobra.Command{
	Use:   "lower glob|file...",
	Short: "names to lower case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg.ToLower(true)
		View(args)
	},
}

func init() {
	rootCmd.AddCommand(lowerCmd)
}
