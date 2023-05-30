package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/name"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// kebabCmd represents the kebab command
var kebabCmd = &cobra.Command{
	Use:   "kebab",
	Short: "all files to kebab",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("casing", name.Kebab)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(kebabCmd)
}
