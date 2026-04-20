package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	semVer            string
	commitHash        string
	isGitDirty        string
	isSnapshot        string
	goOs              string
	goArch            string
	gitUrl            string
	builtBranch       string
	builtDate         string
	semReleaseVersion = semVer +
		func(prefix string, value string) string {
			if value == "" {
				return ""
			}

			return prefix + value
		}("+", goArch) +
		func(prefix string, value string) string {
			if value == "" {
				return ""
			}

			return prefix + value
		}(".", builtBranch) +
		func(prefix string, value string) string {
			if value == "" {
				return ""
			}

			return prefix + value
		}(".", commitHash)
)

var appConfig config

var rootCmd = &cobra.Command{
	Use:               "gencvcpass",
	Args:              cobra.NoArgs,
	Short:             "Generates CVC password",
	Long:              `Generates consonant-vowel-consonant patterned password`,
	Version:           semReleaseVersion,
	DisableAutoGenTag: true,
	CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	RunE:              runRootCommand,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	initLoggers()

	rootCmd.PersistentFlags().IntVar(&loggingLevel, "logging", 0,
		fmt.Sprintf("logging level [0...%d] (default 0)", MAX_LOGGING_LEVEL))

	rootCmd.Flags().IntVarP(&appConfig.setsNum, "sets", "s", 4, "number of sets between separators")
	rootCmd.Flags().IntVarP(&appConfig.upperNum, "upper", "u", 2, "number of uppercase letters")
	rootCmd.Flags().IntVarP(&appConfig.digitsNum, "digits", "d", 2, "number of digits")
	rootCmd.Flags().StringVar(&appConfig.separator, "sep", "-", "separator character")
	rootCmd.Flags().BoolVarP(&appConfig.lessNonPolish, "less-non-polish", "l", false, "do not use non-polish consonants")

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.Flags().SortFlags = false
	cobra.EnableCommandSorting = false
}
