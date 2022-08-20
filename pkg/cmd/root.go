package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
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

			// if we have something in the input flag then we need load the contents of that file which will error if
			// it doesn't exist an error if not.  Then we need to load the json file and parse it.  The input parameter
			// will override whatever is being piped on the command line
			if input != "" {
				file, err := os.ReadFile(input)
				if err != nil {
					_, _ = fmt.Fprintln(os.Stderr, "Unable to open file for reading -", input)
					os.Exit(1)
				}
				fmt.Println(string(file))
			} else {
				// if nothing in the input flag then we can load from stdin
				stat, err := os.Stdin.Stat()
				if err != nil {
					panic(err)
				}
				if (stat.Mode() & os.ModeCharDevice) == 0 {
					bytes, err := io.ReadAll(os.Stdin)
					if err != nil {
						panic(err)
					}
					str := string(bytes)
					fmt.Println(str)
				} else {
					fmt.Println("no pipe :(")
				}
			}
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
