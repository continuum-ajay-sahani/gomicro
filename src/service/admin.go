package service

import (
	"net/http"
	"practice/gomicro/src/app"
	"practice/gomicro/src/entity"
	"practice/gomicro/src/lib"
	"practice/gomicro/src/model"
)

// adminCreate used for application to create parking lot handler
var adminCreate = func(w http.ResponseWriter, r *http.Request) {
	app.Service.LoggerService.Info("admin create parking lot")
	adminCreate := entity.AdminCreate{}
	lib.ParseRequestBody(r, &adminCreate)
	model.AdminCreate(adminCreate)
	lib.RenderJSONCreated(w, r, &adminCreate)
}
