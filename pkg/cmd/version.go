package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of snyk-to-md",
	Long:  `All software has versions. This is synk-to-md's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("snyk-to-md version v0.1")
	},
}
