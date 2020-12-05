package service

import (
	"fmt"
	"net/http"
	"practice/gomicro/src/app"
	"practice/gomicro/src/entity"
	"practice/gomicro/src/lib"
	"practice/gomicro/src/model"
)

// parking used for application to  parking vehicle
var parking = func(w http.ResponseWriter, r *http.Request) {
	app.Service.LoggerService.Info("parking vehicle")
	pkLot := entity.ParkingLot{}
	lib.ParseRequestBody(r, &pkLot)
	err := model.Parking(pkLot)
	if err != nil {
		fmt.Println(err)
		//lib.SendError(w, r, 51, "car parking error", err)
		//return
	}
	lib.RenderJSONCreated(w, r, &pkLot)
}

// parking used for application to  parkout vehicle
var parkout = func(w http.ResponseWriter, r *http.Request) {
	app.Service.LoggerService.Info("parkout vehicle")
	pkLot := entity.ParkingLot{}
	lib.ParseRequestBody(r, &pkLot)
	model.Parkout(pkLot)
	lib.RenderJSONCreated(w, r, &pkLot)
}
