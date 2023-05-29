package cmd

import (
	"log"
	"os"

	"github.com/ohzqq/rename"
	"github.com/spf13/cobra"
)

var (
	batch = rename.Rename()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rename",
	Short: "batch rename files",
	Long:  `utility for batch renaming files`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("interactive", "i", false, "run tui to interactively rename files")
	rootCmd.PersistentFlags().StringVarP(&batch.Sep, "separator", "s", "", "separator for joining words")
	rootCmd.PersistentFlags().IntVar(&batch.Min, "min", 1, "staring num for enumeration")
	rootCmd.PersistentFlags().IntVar(&batch.Max, "max", -1, "end num for enumeration")
	rootCmd.PersistentFlags().BoolVarP(&batch.Pad, "pad", "p", false, "zero pad files")
	rootCmd.PersistentFlags().BoolVarP(&batch.Sanitize, "sanitize", "z", false, "asiify and remove characters")
}

func ValidateArgs(args []string) *rename.Batch {
	switch len(args) {
	case 0:
		log.Fatal("requires either a glob or list of files")
	case 1:
		batch.Glob(args[0])
	default:
		batch.Files(args)
	}
	return batch
}
