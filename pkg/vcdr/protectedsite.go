package vcdr

type GetProtectedSitesResponse struct {
	Cursor         string `json:"cursor"`
	ProtectedSites []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"protected_sites"`
}

type ProtectedSiteDetails struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
