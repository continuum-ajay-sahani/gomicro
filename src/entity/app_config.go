package entity

// AppConfig contain application configuration information
type AppConfig struct {
	ListenURL  string     `json:"ListenURL"`
	URLPrefix  string     `json:"URLPrefix"`
	APIVersion string     `json:"APIVersion"`
	Log        AppLog     `json:"Log"`
	Version    AppVersion `json:"Version"`
}
