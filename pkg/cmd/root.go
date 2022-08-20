package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/leeturner/snyk-to-md/pkg/log"
	"github.com/leeturner/snyk-to-md/pkg/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	inputFlag = "input"
	debugFlag = "debug"
)

type Config struct {
	Input string `mapstructure:"input"`
	Debug bool   `mapstructure:"debug"`
}

var (
	config Config

	// definition of the root command
	rootCmd = &cobra.Command{
		Use:     "snyk-to-md",
		Short:   "Export test json reports from the snyk CLI to markdown ",
		Long:    `The Snyk JSON to Markdown Mapper takes the json outputted from "snyk test --json" and creates a local markdown file displaying the vulnerabilities discovered.`,
		Example: "snyk-to-md",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := log.Setup(config.Debug)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			// if we have something in the input flag then we need load the contents of that file which will error if
			// it doesn't exist an error if not.  Then we need to load the json file and parse it.  The input parameter
			// will override whatever is being piped on the command line
			if cmd.Flags().Changed(inputFlag) {
				contents, err := getFileContent(config.Input, logger)
				if err != nil {
					logger.Error(err.Error())
					os.Exit(1)
				}
				resultMarkdown, err := service.Convert(contents, logger)
				if err != nil {
					logger.Error(err.Error())
					os.Exit(1)
				}
				fmt.Println(resultMarkdown)
				return
			}

			// if nothing in the input flag then we can load from stdin
			stat, err := os.Stdin.Stat()
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				bytes, err := io.ReadAll(os.Stdin)
				if err != nil {
					logger.Error(err.Error())
					os.Exit(1)
				}
				contents := string(bytes)
				resultMarkdown, err := service.Convert(contents, logger)
				if err != nil {
					logger.Error(err.Error())
					os.Exit(1)
				}
				fmt.Println(resultMarkdown)
				return
			}

			fmt.Println("No input flag specified and nothing piped in on the command line")
		},
	}
)

// Main() runs the available Cobra commands
func Main() {
	cobra.CheckErr(Execute())
}

func Execute() error {
	if err := initFlags(); err != nil {
		return err
	}
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func initFlags() error {
	fs := rootCmd.PersistentFlags()
	fs.StringP("input", "i", "", "input path from where to read the json. Defaults to stdin")
	fs.BoolP("debug", "d", false, "determines whether the application runs with debug log messages enable")

	if err := viper.BindPFlags(fs); err != nil {
		return err
	}
	cobra.OnInitialize(initConfig)
	return nil
}

func initConfig() {
	replacer := strings.NewReplacer("-", "_") // allowing env vars like 'example-variable' to be defined as EXAMPLE_VARIABLE
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()     // read value from ENV variables
	viper.Unmarshal(&config) // storing the whole config in a single struct
}
