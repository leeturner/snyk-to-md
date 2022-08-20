package service

import (
	"go.uber.org/zap"
)

func Convert(jsonInput string, logger zap.SugaredLogger) (string, error) {
	logger.Debugf("Converting jsonInput... \n raw jsonInput: %s", jsonInput)
	// TODO conversion
	return jsonInput, nil
}
