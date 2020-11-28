package service

import (
	"net/http"
	"practice/gomicro/src/lib"
	"practice/gomicro/src/model"
)

// handlerAppVersion used for application version handler
var handlerAppVersion = func(w http.ResponseWriter, r *http.Request) {
	lib.RenderJSON(w, r, model.AppVersion())
}
