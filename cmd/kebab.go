package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
)

// kebabCmd represents the kebab command
var kebabCmd = &cobra.Command{
	Use:   "kebab glob|file...",
	Short: "all files to kebab",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.SetCase(opt.Kebab)
		View(args)
	},
}

func init() {
	rootCmd.AddCommand(kebabCmd)
}
