package cmd

import (
	"errors"
	"go.uber.org/zap"
	"io"
	"os"
)

// getContent returns the content of the input file if input was given, or stdin if it wasn't.
func getContent(inputProvided bool, inputFile string, logger zap.SugaredLogger) (string, error) {
	if inputProvided {
		return getFileContent(inputFile, logger)
	}
	return getStsInContent()
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

// getStsInContent returns the content of the stdin
func getStsInContent() (string, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	return "", errors.New("no data provided on STDIN")
}
