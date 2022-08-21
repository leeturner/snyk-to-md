package cmd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	log := zaptest.NewLogger(t)
	logger := *log.Sugar()
	tests := []struct {
		testName      string
		jsonInputFile string
		expectedJson  string
		expectedErr   error
	}{
		{
			testName:      "input file does not exist",
			jsonInputFile: "./doesnt-exist.txt",
			expectedJson:  "",
			expectedErr:   errors.New("could not open the file"),
		},
		{
			testName:      "input file does exist",
			jsonInputFile: "../../test-data/dummy-report.json",
			expectedJson:  `{"ok": false}`,
			expectedErr:   nil,
		},
	}

	for _, data := range tests {
		t.Run(data.testName, func(t *testing.T) {
			resultMD, resultErr := getFileContent(data.jsonInputFile, logger)
			assert.Equal(t, data.expectedJson, resultMD)
			assert.Equal(t, data.expectedErr, resultErr)
		})
	}
}
