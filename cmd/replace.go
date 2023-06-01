package cmd

import (
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
		View(args)
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)
}
