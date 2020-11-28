package logger

import (
	"errors"
	"fmt"
	"practice/gomicro/src/lib"

	"go.uber.org/zap"
)

// Log singleton object for application log
var Log *zap.Logger

// Load app log once
var Load = func() error {
	var err error
	Log, err = lib.GetLogger()

	if err != nil {
		errMsg := fmt.Sprintf("error in loading log = %v", err)
		return errors.New(errMsg)
	}
	return nil
}
