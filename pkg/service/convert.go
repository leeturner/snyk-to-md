package service

import (
	"log"
	"os"
)

func Convert(jsonInput string) (string, error) {
	log.Println("Converting ..")
	log.Printf("jsonInput: %s", jsonInput)

	contents, err := os.ReadFile(jsonInput)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
