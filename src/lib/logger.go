package lib

import "go.uber.org/zap"

var logger *zap.Logger

// GetLogger return logger object for application logging
var GetLogger = func() (*zap.Logger, error) {
	var err error
	logger, err = zap.NewProduction()
	return logger, err
}
