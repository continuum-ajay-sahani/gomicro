package service

import (
	"net/http"
	"practice/gomicro/src/lib"
	"practice/gomicro/src/model"
)

// handlerAppHealth used for application health handler
var handlerAppHealth = func(w http.ResponseWriter, r *http.Request) {
	lib.RenderJSON(w, r, model.AppHealth())
}
