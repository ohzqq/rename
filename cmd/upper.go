package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
)

// upperCmd represents the upper command
var upperCmd = &cobra.Command{
	Use:   "upper",
	Short: "names to upper case",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.SetCase(opt.Upper)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(upperCmd)
}
