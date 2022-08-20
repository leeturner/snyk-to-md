package cmd

import (
	"errors"
	"os"
)

func doesFileExist(input string) (bool, error) {
	_, err := os.Open(input)
	if err != nil {
		return false, errors.New("could not open the file")
	}
	return true, nil
}
