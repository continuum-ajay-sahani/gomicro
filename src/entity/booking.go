package entity

// Booking contain car booking information
type Booking struct {
	ID          string
	RiderUserID string
	CarNumber   string
	StartTime   int64
	EndTime     int64
	Status      string
}
