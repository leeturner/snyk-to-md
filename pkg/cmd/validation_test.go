package cmd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoesFileExist(t *testing.T) {
	tests := []struct {
		testName     string
		fileName     string
		expectedBool bool
		expectedErr  error
	}{
		{
			testName:     "file does not exist",
			fileName:     "abc",
			expectedBool: false,
			expectedErr:  errors.New("could not open the file"),
		},
		{
			testName:     "file does exist",
			fileName:     "../../test-data/dummy-report.json",
			expectedBool: true,
			expectedErr:  nil,
		},
	}

	for _, data := range tests {
		t.Run(data.testName, func(t *testing.T) {
			resultBool, resultErr := doesFileExist(data.fileName)
			assert.Equal(t, data.expectedBool, resultBool)
			assert.Equal(t, data.expectedErr, resultErr)
		})
	}
}
