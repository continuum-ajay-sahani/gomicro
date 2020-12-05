package service

import (
	"net/http"
	"practice/gomicro/src/app"
	"practice/gomicro/src/lib"
	"practice/gomicro/src/model"
)

// handlerAppHealth used for application health handler
var handlerAppHealth = func(w http.ResponseWriter, r *http.Request) {
	app.Service.LoggerService.Info("health api")
	//routeVars := mux.Vars(r)
	lib.RenderJSON(w, r, model.AppHealth())
}
