package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new name glob|file...",
	Short: "new basename for files",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cfg.NewName(args[0])
		if cfg.Zeroes() == -1 {
			cfg.SetZeroes(0)
		}
		View(args[1:])
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
