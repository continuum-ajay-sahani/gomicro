package persistent

import "practice/gomicro/src/entity"

// ParkingLot maintain offset of user
var ParkingLot map[string]entity.Admin

// Load init persistancy
func Load() {
	ParkingLot = make(map[string]entity.Admin)
}

// CreateLot create parking lot into the system
func CreateLot(ac entity.AdminCreate) error {
	bikeLot := make(map[string]entity.Vehicle)
	carLot := make(map[string]entity.Vehicle)
	admin := entity.Admin{}
	admin.BikeLot = bikeLot
	admin.CarLot = carLot
	admin.Create = ac
	ParkingLot[ac.ID] = admin
	return nil
}
