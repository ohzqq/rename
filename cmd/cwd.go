package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cwdCmd represents the cwd command
var cwdCmd = &cobra.Command{
	Use:   "cwd",
	Short: "rename files using the cwd or base dir",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set(opt.Dir, true)
		if cfg.Zeroes() == -1 {
			cfg.SetZeroes(0)
		}
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(cwdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cwdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cwdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
