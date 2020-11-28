package model

import (
	"practice/gomicro/src/app"
	"practice/gomicro/src/entity"
	"reflect"
	"testing"
)

func TestAppVersion(t *testing.T) {
	old := app.Service
	as := &app.App{}
	cs := entity.AppConfig{}
	newVersion := entity.AppVersion{
		ServiceName: "test",
		Major:       "1",
		Minor:       "2",
		Patch:       "3",
	}
	cs.Version = newVersion
	as.ConfigService = cs
	app.Service = as
	app.Service.ConfigService.Version = newVersion
	defer func() {
		app.Service = old
	}()

	v := AppVersion()
	if !reflect.DeepEqual(v, newVersion) {
		t.Fatal("invalid version object")
	}
}
