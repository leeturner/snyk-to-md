package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	log := zaptest.NewLogger(t)
	logger := *log.Sugar()
	tests := []struct {
		testName           string
		jsonInput          string
		displayOnlySummary bool
		displayRemediation bool
		expectedMD         string
		expectedErr        error
	}{
		{
			testName:           "invalid snyk json",
			jsonInput:          `{"invalid": "json"}`,
			displayOnlySummary: false,
			displayRemediation: false,
			expectedMD:         "",
			expectedErr:        errors.New("unable to parse snyk json"),
		},
		{
			testName:           "valid snyk json with 1 vulnerability",
			jsonInput:          readTestFileContents("test-report-1-vuln.json"),
			displayOnlySummary: false,
			displayRemediation: false,
			expectedMD:         readTestFileContents("expected-test-report-1-vuln.md"),
			expectedErr:        nil,
		},
	}

	for _, data := range tests {
		t.Run(data.testName, func(t *testing.T) {
			resultMD, resultErr := Convert(data.jsonInput, data.displayOnlySummary, data.displayRemediation, logger)
			assert.Equal(t, data.expectedMD, resultMD)
			assert.Equal(t, data.expectedErr, resultErr)
		})
	}
}

func readTestFileContents(fileName string) string {
	contents, err := os.ReadFile("../../test-data/" + fileName)
	if err != nil {
		return ""
	}
	return string(contents)
}
