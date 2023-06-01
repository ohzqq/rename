package cmd

import (
	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
)

// lowerCamelCmd represents the lower command
var lowerCamelCmd = &cobra.Command{
	Use:   "lower_camel glob|file...",
	Short: "names to lowerCamel",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.SetCase(opt.LowerCamel)
		View(args)
	},
}

func init() {
	rootCmd.AddCommand(lowerCamelCmd)
}
