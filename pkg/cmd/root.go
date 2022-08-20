package cmd

import (
	"fmt"
	"github.com/leeturner/snyk-to-md/pkg/service"
	"github.com/spf13/cobra"
	"io"
	"os"
)

const (
	inputFlag = "input"
)

var (
	// flags
	input string

	// definition of the root command
	rootCmd = &cobra.Command{
		Use:     "snyk-to-md",
		Short:   "Export test json reports from the snyk CLI to markdown ",
		Long:    `The Snyk JSON to Markdown Mapper takes the json outputted from "snyk test --json" and creates a local markdown file displaying the vulnerabilities discovered.`,
		Example: "snyk-to-md",
		Run: func(cmd *cobra.Command, args []string) {
			// if we have something in the input flag then we need load the contents of that file which will error if
			// it doesn't exist an error if not.  Then we need to load the json file and parse it.  The input parameter
			// will override whatever is being piped on the command line
			if cmd.Flags().Changed(inputFlag) {
				exists, err := doesFileExist(input)
				if err != nil || !exists {
					_, _ = fmt.Fprintln(os.Stderr, err.Error())
					os.Exit(1)
				}
				// we know the file exists
				markdown, err := service.Convert(input)
				if err != nil {
					_, _ = fmt.Fprintln(os.Stderr, err.Error())
					os.Exit(1)
				}
				fmt.Printf(markdown)
				return
			}

			// if nothing in the input flag then we can load from stdin
			stat, err := os.Stdin.Stat()
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("could't read from stdin. error: %w", err))
				os.Exit(1)
			}
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				bytes, err := io.ReadAll(os.Stdin)
				if err != nil {
					panic(err)
				}
				str := string(bytes)
				fmt.Println(str)
				return
			}

			fmt.Println("No input flag specified and nothing piped in on the command line")
		},
	}
)

func init() {
	fs := rootCmd.PersistentFlags()
	fs.StringVarP(&input, inputFlag, "i", "", "input path from where to read the json. Defaults to stdin")
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
