package vmc

type SupportWindow struct {
	DurationHours int `json:"duration_hours"`
	Metadata      struct {
	} `json:"metadata"`
	Sddcs           []string `json:"sddcs"`
	Seats           int      `json:"seats"`
	StartDay        string   `json:"start_day"`
	StartHour       int      `json:"start_hour"`
	SupportWindowID string   `json:"support_window_id"`
}

type SupportWindows []SupportWindow
