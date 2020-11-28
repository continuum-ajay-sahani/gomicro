package persistent

import "practice/gomicro/src/entity"

// RiderStorage maintain rider information in the system
var RiderStorage map[string]entity.Rider

// DriverStorage maintain driver information in the system
var DriverStorage map[string]entity.Driver

// VehicleStorage maintain vehicle information in the system
var VehicleStorage map[string]entity.Vehicle

// BookingStorage maintain booking information in the system
var BookingStorage map[string]entity.Booking

// Load init persistancy
func Load() {
	RiderStorage = make(map[string]entity.Rider)
	DriverStorage = make(map[string]entity.Driver)
	VehicleStorage = make(map[string]entity.Vehicle)
	BookingStorage = make(map[string]entity.Booking)
}
