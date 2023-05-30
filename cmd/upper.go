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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upperCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upperCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
