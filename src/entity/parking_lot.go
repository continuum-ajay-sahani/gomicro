package entity

// ParkingLot will contain information about the lot
type ParkingLot struct {
	AdminID string  `json:"AdminID"`
	Vehicle Vehicle `json:"Vehicle"`
	//InTime  int64
	//OutTime int64
}
