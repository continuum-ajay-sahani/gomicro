package model

import (
	"practice/gomicro/src/app"
	"practice/gomicro/src/entity"
)

// AppVersion return application version information
var AppVersion = func() entity.AppVersion {
	return app.Service.ConfigService.Version
}
