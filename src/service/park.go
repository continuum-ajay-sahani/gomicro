package service

import (
	"net/http"
	"practice/gomicro/src/app"
	"practice/gomicro/src/entity"
	"practice/gomicro/src/lib"
	"practice/gomicro/src/model"
)

// parking used for application to  parking vehicle
var parking = func(w http.ResponseWriter, r *http.Request) {
	app.Service.LoggerService.Info("parking vehicle")
	var pkLot entity.ParkingLot
	lib.ParseHTTPRequestBody(r, &pkLot)
	err := model.Parking(pkLot)
	if err != nil {
		lib.SendError(w, r, http.StatusBadRequest, "car parking error", err)
		return
	}
	lib.RenderJSONCreated(w, r, &pkLot)
}

// parking used for application to  parkout vehicle
var parkout = func(w http.ResponseWriter, r *http.Request) {
	app.Service.LoggerService.Info("parkout vehicle")
	var pkLot entity.ParkingLot
	lib.ParseHTTPRequestBody(r, &pkLot)
	err := model.Parkout(pkLot)
	if err != nil {
		lib.SendError(w, r, http.StatusBadRequest, "car parkout error", err)
		return
	}
	lib.RenderJSONCreated(w, r, &pkLot)
}
