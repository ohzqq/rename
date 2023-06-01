package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
)

// kebabCmd represents the kebab command
var kebabCmd = &cobra.Command{
	Use:   "kebab",
	Short: "all files to kebab",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.SetCase(opt.Kebab)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(kebabCmd)
}
