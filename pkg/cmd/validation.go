package cmd

import (
	"errors"
	"go.uber.org/zap"
	"os"
)

func getFileContent(jsonInput string, logger zap.SugaredLogger) (string, error) {
	logger.Debug("Reading from file...")
	contents, err := os.ReadFile(jsonInput)
	if err != nil {
		return "", errors.New("could not open the file")
	}
	logger.Debug("Reading from file was a success.")
	return string(contents), nil
}
