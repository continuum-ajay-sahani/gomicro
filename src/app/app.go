package app

import (
	"practice/gomicro/src/entity"

	"go.uber.org/zap"
)

// App defines the whole configuration of the Application.
// It incorporates partial configurations of components
type App struct {
	ConfigService entity.AppConfig
	LoggerService *zap.Logger
}

// Service is the whole configuration of the Application
var Service *App
