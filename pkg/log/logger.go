package log

import (
	"go.uber.org/zap"
)

// Setup helps us initialising our Logger and Sugar
func Setup(debug bool) (zap.SugaredLogger, error) {
	logger, err := initLogger(debug)
	if err != nil {
		return zap.SugaredLogger{}, err
	}
	return *logger.Sugar(), nil
}

func initLogger(debug bool) (*zap.Logger, error) {
	if debug {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}
