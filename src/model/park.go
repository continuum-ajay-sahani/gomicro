package model

import (
	"errors"
	"practice/gomicro/src/entity"
	"practice/gomicro/src/persistent"
)

// constants
const (
	BIKE = 1
	CAR  = 2
)

// Parking do parking
var Parking = func(ad entity.ParkingLot) error {
	db := persistent.ParkingLot
	v, ok := db[ad.AdminID]
	if !ok {
		return errors.New("admin id not found")
	}
	vehicleType := ad.Vehicle.VehicleType
	if vehicleType == BIKE {
		if len(v.BikeLot) >= v.Create.Bike {
			return errors.New("bike parking full")
		}
		// create intime
		v.BikeLot[ad.Vehicle.Number] = ad.Vehicle
	}
	if vehicleType == CAR {
		if len(v.CarLot) >= v.Create.Car {
			return errors.New("car parking full")
		}
		v.CarLot[ad.Vehicle.Number] = ad.Vehicle
	}
	return nil
}

// Parkout do parkout
var Parkout = func(ad entity.ParkingLot) error {
	db := persistent.ParkingLot
	v, ok := db[ad.AdminID]
	if !ok {
		return errors.New("admin id not found")
	}
	vehicleType := ad.Vehicle.VehicleType
	if vehicleType == BIKE {
		_, k := v.BikeLot[ad.Vehicle.Number]
		if !k {
			return errors.New("bike details not found")
		}
		delete(v.BikeLot, ad.Vehicle.Number)
	}
	if vehicleType == CAR {
		_, k := v.CarLot[ad.Vehicle.Number]
		if !k {
			return errors.New("car details not found")
		}
		delete(v.CarLot, ad.Vehicle.Number)
	}
	return nil
}
