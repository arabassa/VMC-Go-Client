package vcdr

import "encoding/json"

type CloudFileSystemsResponse struct {
	CloudFileSystems []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"cloud_file_systems"`
}

type CloudFileSystemDetails struct {
	CapacityGib    json.Number `json:"capacity_gib"`
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	RecoverySddcID string      `json:"recovery_sddc_id"`
	UsedGib        json.Number `json:"used_gib"`
}
