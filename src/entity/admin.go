package entity

// AdminCreate contain information regarding admin api
type AdminCreate struct {
	ID   string `json:"ID"`
	Bike int    `json:"Bike"`
	Car  int    `json:"Car"`
}

// Admin data structure to hold entire data
type Admin struct {
	ID      string
	Create  AdminCreate
	BikeLot map[string]Vehicle
	CarLot  map[string]Vehicle
}
