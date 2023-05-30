package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace regex str file...",
	Short: "search and replace text",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("find", args[0])
		viper.Set("replace", args[1])
		names := ValidateArgs(args[2:]).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
