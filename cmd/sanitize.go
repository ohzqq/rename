package cmd

import (
	"github.com/londek/reactea"
	"github.com/ohzqq/rename/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sanitizeCmd represents the sanitize command
var sanitizeCmd = &cobra.Command{
	Use:   "sanitize",
	Short: "sanitize filenames",
	Long:  `remove special characters, spaces, etc from file names`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("sanitize", true)
		names := ValidateArgs(args)
		pre := ui.New(names).Route("padding")
		program := reactea.NewProgram(pre)

		if err := program.Start(); err != nil {
			panic(err)
		}

		//fmt.Println(t)
	},
}

func init() {
	rootCmd.AddCommand(sanitizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sanitizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sanitizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
