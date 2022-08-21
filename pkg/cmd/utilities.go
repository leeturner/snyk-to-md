package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
)

// getContent returns the content of the input file if input was given, or stdin if it wasn't.
func getContent(inputProvided bool, inputFile string, logger zap.SugaredLogger) (string, error) {
	if inputProvided {
		return getFileContent(inputFile, logger)
	}
	return getStdInContent(logger)
}

// getFileContent returns the content of the input file
func getFileContent(jsonInput string, logger zap.SugaredLogger) (string, error) {
	logger.Debug("Reading from file...")
	contents, err := os.ReadFile(jsonInput)
	if err != nil {
		return "", errors.New("could not open the file")
	}
	logger.Debug("Reading from file was a success.")
	return string(contents), nil
}

// getStdInContent returns the content of the stdin
func getStdInContent(logger zap.SugaredLogger) (string, error) {
	logger.Debug("Reading from STDIN...")
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", fmt.Errorf("Couldn't get the STDIN stat: %s", err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("Couldn't read from STDIN: %s", err)
		}
		return string(bytes), nil
	}
	return "", errors.New("No data provided on STDIN")
}

// exportResults exports a string to the STDOUT, except if an output file was provided,
// in which case it exports the string to the output file.
func exportResults(outputProvided bool, fileName string, contents string, logger zap.SugaredLogger) error {
	if outputProvided {
		return writeToFile(fileName, contents, logger)
	}
	writeToStdOut(contents, logger)
	return nil
}

func writeToFile(fileName string, contents string, logger zap.SugaredLogger) error {
	logger.Debug("Writing to file...")
	err := os.WriteFile(fileName, []byte(contents), 0644)
	if err != nil {
		return fmt.Errorf("Couldn't write to file: %s", err)
	}
	return err
}

func writeToStdOut(contents string, logger zap.SugaredLogger) {
	logger.Debug("Writing to STDOUT...")
	fmt.Println(contents)
}
