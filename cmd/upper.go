package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/spf13/cobra"
)

// upperCmd represents the upper command
var upperCmd = &cobra.Command{
	Use:   "upper glob|file...",
	Short: "names to UPPER case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg.ToUpper(true)
		View(args)
	},
}

func init() {
	rootCmd.AddCommand(upperCmd)
}
