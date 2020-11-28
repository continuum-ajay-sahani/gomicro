package persistent

import "practice/gomicro/src/entity"

// RiderStorage maintain rider information in the system
var riderStorage map[string]entity.Rider

// DriverStorage maintain driver information in the system
var driverStorage map[string]entity.Driver

// VehicleStorage maintain vehicle information in the system
var vehicleStorage map[string]entity.Vehicle

// BookingStorage maintain booking information in the system
var bookingStorage map[string]entity.Booking

// RatingStorage maintain booking information in the system
var ratingStorage map[string]entity.Rating

// Load init persistancy
func Load() {
	riderStorage = make(map[string]entity.Rider)
	driverStorage = make(map[string]entity.Driver)
	vehicleStorage = make(map[string]entity.Vehicle)
	bookingStorage = make(map[string]entity.Booking)
	ratingStorage = make(map[string]entity.Rating)
}

// AddRider insert rider information in the system
func AddRider(riderID string, rider entity.Rider) {
	riderStorage[riderID] = rider
}

// AddDriver insert driver information in the system
func AddDriver(driverID string, driver entity.Driver) {
	driverStorage[driverID] = driver
}

// AddVehicle insert vehicle information in the system
func AddVehicle(vehicleID string, vehicle entity.Vehicle) {
	vehicleStorage[vehicleID] = vehicle
}

// AddBooking insert booking information in the system
func AddBooking(bookingID string, booking entity.Booking) {
	bookingStorage[bookingID] = booking
}

// AddRating insert rating information in the system
func AddRating(driverID string, rating entity.Rating) {
	ratingStorage[driverID] = rating
}
