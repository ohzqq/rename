package cmd

import (
	"github.com/spf13/cobra"
)

// sanitizeCmd represents the sanitize command
var sanitizeCmd = &cobra.Command{
	Use:   "sanitize",
	Short: "sanitize filename",
	Long:  `remove special characters, spaces, etc from file names`,
	Run: func(cmd *cobra.Command, args []string) {
		files := ValidateArgs(args)
		for _, file := range files {
			println(file)
		}
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
