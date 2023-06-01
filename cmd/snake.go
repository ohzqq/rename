package cmd

import (
	"fmt"

	"github.com/ohzqq/rename/opt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// snakeCmd represents the snake command
var snakeCmd = &cobra.Command{
	Use:   "snake",
	Short: "names to snake",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set(opt.Casing, opt.Snake)
		names := ValidateArgs(args).Transform()
		fmt.Println(names)
	},
}

func init() {
	rootCmd.AddCommand(snakeCmd)
}
