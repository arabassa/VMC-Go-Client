package vmc

type Reservation struct {
	InMaintenanceMode   bool   `json:"in_maintenance_mode"`
	InMaintenanceWindow bool   `json:"in_maintenance_window"`
	ReservationID       string `json:"reservation_id"`
	ReservationSchedule struct {
	} `json:"reservation_schedule"`
	SddcID string `json:"sddc_id"`
}

type Reservations []Reservation
