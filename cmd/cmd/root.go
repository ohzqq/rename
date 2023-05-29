package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	NameSep string
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
	rootCmd.PersistentFlags().StringVarP(&NameSep, "separator", "s", "", "separator for joining words")
}

func ValidateArgs(args []string) []string {
	switch len(args) {
	case 0:
		log.Fatal("requires either a glob or list of files")
		return args
	case 1:
		files, err := filepath.Glob(args[0])
		if err != nil {
			log.Fatal(err)
		}
		return files
	default:
		return args
	}
}
