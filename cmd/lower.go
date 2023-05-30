package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/name"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lowerCamelCmd represents the lower command
var lowerCamelCmd = &cobra.Command{
	Use:   "lower_camel",
	Short: "names to lowerCamel",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("casing", name.LowerCamel)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(lowerCamelCmd)
}
