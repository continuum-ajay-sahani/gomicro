package entity

// Vehicle holds information about vehicle
type Vehicle struct {
	CarNumber string
	Lat       float64
	Lon       float64
	Type      string
	Available bool
	DriverID  string
}
