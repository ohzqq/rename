package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/spf13/cobra"
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace regex str glob|file...",
	Short: "search and replace text",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		cfg.SetFind(args[0])
		cfg.SetReplace(args[1])
		View(args[2:])
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)
}
