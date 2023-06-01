package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
)

// lowerCamelCmd represents the lower command
var lowerCamelCmd = &cobra.Command{
	Use:   "lower_camel",
	Short: "names to lowerCamel",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.SetCase(opt.LowerCamel)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(lowerCamelCmd)
}
