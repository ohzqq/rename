package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dirCmd represents the cwd command
var dirCmd = &cobra.Command{
	Use:   "dir glob|file...",
	Short: "rename files using the cwd or base dir",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set(opt.Dir, true)
		if cfg.Zeroes() == -1 {
			cfg.SetZeroes(0)
		}
		View(args)
	},
}

func init() {
	rootCmd.AddCommand(dirCmd)
}
