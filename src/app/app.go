package app

import (
	"os"
	"practice/gomicro/src/entity"

	"go.uber.org/zap"
)

// App defines the whole configuration of the Application.
// It incorporates partial configurations of components
type App struct {
	ConfigService entity.AppConfig
	LoggerService *zap.Logger
}

// OSSingnal accept graceful shutdowns when quit via SIGINT (Ctrl+C)
// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
var OSSingnal chan os.Signal

// Service is the whole configuration of the Application
var Service *App
