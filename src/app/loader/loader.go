package loader

import (
	"os"
	"practice/gomicro/src/app"
	"practice/gomicro/src/config"
	"practice/gomicro/src/logger"
	"practice/gomicro/src/persistent"
)

// LoadApplicationServices loads all partial configurations of components
// and populates the app.App with the configuration data
func LoadApplicationServices() error {
	initOSSignalChannel()
	app.Service = &app.App{}

	err := config.Load()
	if err != nil {
		return err
	}
	app.Service.ConfigService = config.Configuration

	err = logger.Load()
	if err != nil {
		return err
	}
	app.Service.LoggerService = logger.Log
	persistent.Load()

	return nil
}

var initOSSignalChannel = func() {
	app.OSSingnal = make(chan os.Signal, 1)
}
