package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
)

// snakeCmd represents the snake command
var snakeCmd = &cobra.Command{
	Use:   "snake glob|file...",
	Short: "names to snake",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.SetCase(opt.Snake)
		View(args)
	},
}

func init() {
	rootCmd.AddCommand(snakeCmd)
}
