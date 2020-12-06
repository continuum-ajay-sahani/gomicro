package service

import (
	"fmt"
	"net/http"
	"practice/gomicro/src/app"
	"practice/gomicro/src/entity"
	"practice/gomicro/src/lib"
	"practice/gomicro/src/model"
)

// adminCreate used for application to create parking lot handler
var adminCreate = func(w http.ResponseWriter, r *http.Request) {
	app.Service.LoggerService.Info("admin create parking lot")
	var ac entity.AdminCreate
	err := lib.ParseHTTPRequestBody(r, &ac)
	if err != nil {
		app.Service.LoggerService.Error(err.Error())
		return
	}
	fmt.Println(ac.ID)
	//lib.ParseRequestBody(r, &adminCreate)
	model.AdminCreate(ac)
	lib.RenderJSONCreated(w, r, &ac)
}
