package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
)

// camelCmd represents the camel command
var camelCmd = &cobra.Command{
	Use:   "camel",
	Short: "names to camel",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.SetCase(opt.Camel)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(camelCmd)
}
