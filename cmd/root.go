package cmd

import (
	"log"
	"os"

	"github.com/londek/reactea"
	"github.com/ohzqq/rename/batch"
	"github.com/ohzqq/rename/opt"
	"github.com/ohzqq/rename/ui"
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
	viper.SetDefault(opt.Sep, "_")
	viper.SetDefault(opt.Casing, opt.Snake)
	viper.SetDefault(opt.Start, 1)
	viper.SetDefault(opt.Position, opt.End)
	viper.SetDefault(opt.Zeroes, -1)
	viper.SetDefault(opt.Asciiify, false)
	viper.SetDefault(opt.Tidy, false)
	viper.SetDefault(opt.Lower, false)
	viper.SetDefault(opt.Upper, false)

	rootCmd.PersistentFlags().BoolP("interactive", "i", false, "run tui to interactively rename files")

	rootCmd.PersistentFlags().String(opt.Sep, "", "separator for joining words")
	viper.BindPFlag(opt.Sep, rootCmd.PersistentFlags().Lookup(opt.Sep))

	rootCmd.PersistentFlags().StringP(opt.Name, "n", "", "use as basename")
	viper.BindPFlag(opt.Name, rootCmd.PersistentFlags().Lookup(opt.Name))

	rootCmd.PersistentFlags().IntP(opt.Start, "m", 1, "min num for enumeration")
	viper.BindPFlag(opt.Start, rootCmd.PersistentFlags().Lookup(opt.Start))

	rootCmd.PersistentFlags().IntP(opt.Zeroes, "z", -1, "zero pad files")
	viper.BindPFlag(opt.Zeroes, rootCmd.PersistentFlags().Lookup(opt.Zeroes))

	rootCmd.PersistentFlags().BoolP(opt.Tidy, "t", false, "asciify and remove special characters")
	viper.BindPFlag(opt.Tidy, rootCmd.PersistentFlags().Lookup(opt.Tidy))

	rootCmd.PersistentFlags().StringP(opt.Casing, "c", opt.Snake, "choose case [camel, kebab, lowerCamel, snake]")
	viper.BindPFlag(opt.Casing, rootCmd.PersistentFlags().Lookup(opt.Casing))
	rootCmd.PersistentFlags().BoolP(opt.Upper, "u", false, "upper case name")
	viper.BindPFlag(opt.Upper, rootCmd.PersistentFlags().Lookup(opt.Upper))
	rootCmd.PersistentFlags().BoolP(opt.Lower, "l", false, "lower case name")
	viper.BindPFlag(opt.Lower, rootCmd.PersistentFlags().Lookup(opt.Lower))

	rootCmd.PersistentFlags().BoolP(opt.Dir, "d", false, "use dir as basename")
	viper.BindPFlag(opt.Dir, rootCmd.PersistentFlags().Lookup(opt.Dir))

	rootCmd.PersistentFlags().StringP(opt.Prefix, "p", "", "add a prefix")
	viper.BindPFlag(opt.Prefix, rootCmd.PersistentFlags().Lookup(opt.Prefix))
	rootCmd.PersistentFlags().StringP(opt.Suffix, "s", "", "add a suffix")
	viper.BindPFlag(opt.Suffix, rootCmd.PersistentFlags().Lookup(opt.Suffix))

	rootCmd.PersistentFlags().StringP(opt.Find, "f", "", "find regex")
	viper.BindPFlag(opt.Find, rootCmd.PersistentFlags().Lookup(opt.Find))
	rootCmd.PersistentFlags().StringP(opt.Replace, "r", "", "replacement string")
	viper.BindPFlag(opt.Replace, rootCmd.PersistentFlags().Lookup(opt.Replace))
	rootCmd.MarkFlagsRequiredTogether(opt.Find, opt.Replace)
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

func View(args []string) {
	names := ValidateArgs(args)
	pre := ui.New(names)
	if !viper.GetBool("interactive") {
		pre.Route(ui.View)
	}
	program := reactea.NewProgram(pre)

	if err := program.Start(); err != nil {
		panic(err)
	}
}
