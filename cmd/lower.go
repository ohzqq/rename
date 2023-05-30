package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/name"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lowerCmd represents the lower command
var lowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "names to lower case",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("casing", name.Lower)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(lowerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lowerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lowerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
