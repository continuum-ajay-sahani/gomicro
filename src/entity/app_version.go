package entity

// AppVersion contain application version specific information
type AppVersion struct {
	ServiceName string `json:"ServiceName"`
	Major       string `json:"Major"`
	Minor       string `json:"Minor"`
	Patch       string `json:"Patch"`
}
