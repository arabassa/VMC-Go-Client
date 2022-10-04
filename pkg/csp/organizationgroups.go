package csp

type ExpandedGroupDto struct {
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
	Description       string `json:"description"`
	DisplayName       string `json:"displayName"`
	Domain            string `json:"domain"`
	GroupType         string `json:"groupType"`
	ID                string `json:"id"`
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
	OwnerOrgID   string `json:"ownerOrgId"`
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
	SharedOrgIds []string `json:"sharedOrgIds"`
	UsersCount   int      `json:"usersCount"`
}

type PagedResponseExpandedGroupDto struct {
	NextLink string `json:"nextLink"`
	PrevLink string `json:"prevLink"`
	Results  []struct {
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
		Description       string `json:"description"`
		DisplayName       string `json:"displayName"`
		Domain            string `json:"domain"`
		GroupType         string `json:"groupType"`
		ID                string `json:"id"`
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
		OwnerOrgID   string `json:"ownerOrgId"`
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
		SharedOrgIds []string `json:"sharedOrgIds"`
		UsersCount   int      `json:"usersCount"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}

type PagedResponseGroupDto struct {
	NextLink string `json:"nextLink"`
	PrevLink string `json:"prevLink"`
	Results  []struct {
		Description  string   `json:"description"`
		DisplayName  string   `json:"displayName"`
		Domain       string   `json:"domain"`
		GroupType    string   `json:"groupType"`
		ID           string   `json:"id"`
		OwnerOrgID   string   `json:"ownerOrgId"`
		SharedOrgIds []string `json:"sharedOrgIds"`
		UsersCount   int      `json:"usersCount"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}

type PagedResponseBaseUserWithProfileDto struct {
	NextLink string `json:"nextLink"`
	PrevLink string `json:"prevLink"`
	Results  []struct {
		Accessible  bool   `json:"accessible"`
		Acct        string `json:"acct"`
		Domain      string `json:"domain"`
		Email       string `json:"email"`
		FirstName   string `json:"firstName"`
		IdpID       string `json:"idpId"`
		LastName    string `json:"lastName"`
		UserID      string `json:"userId"`
		UserProfile struct {
			AlternativeEmail string `json:"alternativeEmail"`
			Language         string `json:"language"`
			Locale           string `json:"locale"`
		} `json:"userProfile"`
		Username string `json:"username"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}
