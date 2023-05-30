package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/name"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// camelCmd represents the camel command
var camelCmd = &cobra.Command{
	Use:   "camel",
	Short: "names to camel",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("casing", name.Camel)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(camelCmd)
}
