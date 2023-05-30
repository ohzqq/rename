package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/name"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// upperCmd represents the upper command
var upperCmd = &cobra.Command{
	Use:   "upper",
	Short: "names to upper case",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("casing", name.Upper)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(upperCmd)
}
