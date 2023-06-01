package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/spf13/cobra"
)

// sanitizeCmd represents the sanitize command
var sanitizeCmd = &cobra.Command{
	Use:   "sanitize",
	Short: "sanitize filenames",
	Long:  `remove special characters, spaces, etc from file names`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg.Tidy(true)
		View(args)
	},
}

func init() {
	rootCmd.AddCommand(sanitizeCmd)
}
