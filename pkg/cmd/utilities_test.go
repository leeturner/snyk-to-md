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
		testName    string
		jsonInput   string
		expectedMD  string
		expectedErr error
	}{
		{
			testName:    "input file does not exist",
			jsonInput:   "./doesnt-exist.txt",
			expectedMD:  "",
			expectedErr: errors.New("could not open the file"),
		},
		{
			testName:    "input file does exist",
			jsonInput:   "../../test-data/dummy-report.json",
			expectedMD:  `{"ok": false}`,
			expectedErr: nil,
		},
	}

	for _, data := range tests {
		t.Run(data.testName, func(t *testing.T) {
			resultMD, resultErr := getFileContent(data.jsonInput, logger)
			assert.Equal(t, data.expectedMD, resultMD)
			assert.Equal(t, data.expectedErr, resultErr)
		})
	}
}
