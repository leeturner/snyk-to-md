package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	// flags
	input string

	// definition of the root command
	rootCmd = &cobra.Command{
		Use:   "snyk-to-md",
		Short: "Export test json reports from the snyk CLI to markdown ",
		Long:  `The Snyk JSON to Markdown Mapper takes the json outputted from "snyk test --json" and creates a local markdown file displaying the vulnerabilities discovered.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Root Command Executing")
			fmt.Println("The input flag was :", input)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "input path from where to read the json. Defaults to stdin")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
