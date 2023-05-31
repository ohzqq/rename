package cmd

import (
	"log"
	"os"

	"github.com/ohzqq/rename/batch"
	"github.com/ohzqq/rename/name"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	viper.SetDefault("sep", "_")
	viper.SetDefault("min", 1)
	viper.SetDefault("max", -1)
	viper.SetDefault("casing", name.Snake)
	viper.SetDefault("pad_position", name.PosEnd)
	viper.SetDefault("zeroes", -1)
	viper.SetDefault("asciiify", false)
	viper.SetDefault("sanitize", false)

	rootCmd.PersistentFlags().BoolP("interactive", "i", false, "run tui to interactively rename files")

	rootCmd.PersistentFlags().String("sep", "", "separator for joining words")
	viper.BindPFlag("sep", rootCmd.PersistentFlags().Lookup("sep"))

	rootCmd.PersistentFlags().Int("min", 1, "staring num for enumeration")
	viper.BindPFlag("min", rootCmd.PersistentFlags().Lookup("min"))

	rootCmd.PersistentFlags().Int("max", -1, "end num for enumeration")
	viper.BindPFlag("max", rootCmd.PersistentFlags().Lookup("max"))

	rootCmd.PersistentFlags().IntP("zeroes", "d", -1, "zero pad files")
	viper.BindPFlag("zeroes", rootCmd.PersistentFlags().Lookup("zeroes"))

	rootCmd.PersistentFlags().BoolP("sanitize", "z", false, "asiify and remove characters")
	viper.BindPFlag("sanitize", rootCmd.PersistentFlags().Lookup("sanitize"))

	rootCmd.PersistentFlags().Bool("cwd", false, "use cwd as basename")
	viper.BindPFlag("cwd", rootCmd.PersistentFlags().Lookup("cwd"))

	rootCmd.PersistentFlags().StringP("prefix", "p", "", "add a prefix")
	viper.BindPFlag("prefix", rootCmd.PersistentFlags().Lookup("prefix"))
	rootCmd.PersistentFlags().StringP("suffix", "s", "", "add a suffix")
	viper.BindPFlag("suffix", rootCmd.PersistentFlags().Lookup("suffix"))

	rootCmd.PersistentFlags().StringP("find", "f", "", "find regex")
	viper.BindPFlag("find", rootCmd.PersistentFlags().Lookup("find"))
	rootCmd.PersistentFlags().StringP("replace", "r", "", "replace regex")
	viper.BindPFlag("replace", rootCmd.PersistentFlags().Lookup("replace"))
	rootCmd.MarkFlagsRequiredTogether("find", "replace")
}

func ValidateArgs(args []string) *batch.Names {
	switch len(args) {
	case 0:
		log.Fatal("requires either a glob or list of files")
	case 1:
		return batch.Glob(args[0])
	default:
		return batch.Files(args)
	}
	return batch.New()
}
