package vcdr

type GetRecoverySddcResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type RecoverySddcDetails struct {
	AvailabilityZones []string `json:"availability_zones"`
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Region            string   `json:"region"`
}
