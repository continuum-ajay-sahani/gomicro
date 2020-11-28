package config

import (
	"errors"
	"fmt"
	"practice/gomicro/src/entity"
	"practice/gomicro/src/lib"
)

// Configuration singleton object for application configuration
var Configuration entity.AppConfig

// FilePath holds the path of application configuration
var FilePath = "../../config.json"

// Load app config once
var Load = func() error {
	err := lib.ReadJSONFile(FilePath, &Configuration)
	if err != nil {
		errMsg := fmt.Sprintf("configuration not found: %v = ", err)
		return errors.New(errMsg)
	}

	return nil
}
