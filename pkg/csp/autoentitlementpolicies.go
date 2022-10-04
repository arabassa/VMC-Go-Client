package csp

type PagedResponseAutoEntitlementPoliciesDto struct {
	NextLink string `json:"nextLink"`
	PrevLink string `json:"prevLink"`
	Results  []struct {
		Description  string `json:"description"`
		DisplayName  string `json:"displayName"`
		ID           string `json:"id"`
		LastEditedAt string `json:"lastEditedAt"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}

type ExpandedAutoEntitlementPoliciesDto struct {
	CustomRoles []struct {
		CreatedBy       string `json:"createdBy"`
		CreatedDate     string `json:"createdDate"`
		ExpiresAt       int64  `json:"expiresAt"`
		LastUpdatedBy   string `json:"lastUpdatedBy"`
		LastUpdatedDate string `json:"lastUpdatedDate"`
		MembershipType  string `json:"membershipType"`
		Name            string `json:"name"`
		Resource        string `json:"resource"`
	} `json:"customRoles"`
	Description       string   `json:"description"`
	DisplayName       string   `json:"displayName"`
	Domains           []string `json:"domains"`
	ID                string   `json:"id"`
	LastEditedAt      string   `json:"lastEditedAt"`
	OrganizationRoles []struct {
		CreatedBy       string `json:"createdBy"`
		CreatedDate     string `json:"createdDate"`
		DisplayName     string `json:"displayName"`
		ExpiresAt       int64  `json:"expiresAt"`
		LastUpdatedBy   string `json:"lastUpdatedBy"`
		LastUpdatedDate string `json:"lastUpdatedDate"`
		MembershipType  string `json:"membershipType"`
		Name            string `json:"name"`
		Resource        string `json:"resource"`
	} `json:"organizationRoles"`
	ServiceRoles []struct {
		ServiceDefinitionID string   `json:"serviceDefinitionId"`
		ServiceRoleNames    []string `json:"serviceRoleNames"`
		ServiceRoles        []struct {
			CreatedBy       string `json:"createdBy"`
			CreatedDate     string `json:"createdDate"`
			ExpiresAt       int64  `json:"expiresAt"`
			LastUpdatedBy   string `json:"lastUpdatedBy"`
			LastUpdatedDate string `json:"lastUpdatedDate"`
			MembershipType  string `json:"membershipType"`
			Name            string `json:"name"`
			Resource        string `json:"resource"`
		} `json:"serviceRoles"`
	} `json:"serviceRoles"`
}
