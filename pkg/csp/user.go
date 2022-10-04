package csp

type UserDto struct {
	Acct           string `json:"acct"`
	Address        string `json:"address"`
	City           string `json:"city"`
	Company        string `json:"company"`
	Country        string `json:"country"`
	CustomerNumber string `json:"customerNumber"`
	Email          string `json:"email"`
	EmailVerified  bool   `json:"emailVerified"`
	EulaInfo       string `json:"eulaInfo"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Password       string `json:"password"`
	RefLink        string `json:"refLink"`
	State          string `json:"state"`
	Tnc            bool   `json:"tnc"`
	TradeID        string `json:"tradeId"`
	UserID         string `json:"userId"`
	UserProfile    struct {
		AlternativeEmail    string `json:"alternativeEmail"`
		CustomerNumber      string `json:"customerNumber"`
		DefaultOrgID        string `json:"defaultOrgId"`
		IsFederated         bool   `json:"isFederated"`
		Language            string `json:"language"`
		LinkedUserIDAccount string `json:"linkedUserIdAccount"`
		Locale              string `json:"locale"`
	} `json:"userProfile"`
	Username  string `json:"username"`
	Website   string `json:"website"`
	WorkPhone string `json:"workPhone"`
	Zipcode   string `json:"zipcode"`
}

type UserInfoDto struct {
	User struct {
		Acct           string `json:"acct"`
		Address        string `json:"address"`
		City           string `json:"city"`
		Company        string `json:"company"`
		Country        string `json:"country"`
		CustomerNumber string `json:"customerNumber"`
		Email          string `json:"email"`
		EmailVerified  bool   `json:"emailVerified"`
		EulaInfo       string `json:"eulaInfo"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		Password       string `json:"password"`
		RefLink        string `json:"refLink"`
		State          string `json:"state"`
		Tnc            bool   `json:"tnc"`
		TradeID        string `json:"tradeId"`
		UserID         string `json:"userId"`
		UserProfile    struct {
			AlternativeEmail    string `json:"alternativeEmail"`
			CustomerNumber      string `json:"customerNumber"`
			DefaultOrgID        string `json:"defaultOrgId"`
			IsFederated         bool   `json:"isFederated"`
			Language            string `json:"language"`
			LinkedUserIDAccount string `json:"linkedUserIdAccount"`
			Locale              string `json:"locale"`
		} `json:"userProfile"`
		Username  string `json:"username"`
		Website   string `json:"website"`
		WorkPhone string `json:"workPhone"`
		Zipcode   string `json:"zipcode"`
	} `json:"user"`
	UserOrgInfo []struct {
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
		DisplayName string `json:"displayName"`
		Name        string `json:"name"`
		OrgRoles    []struct {
			CreatedBy        string `json:"createdBy"`
			CreatedDate      string `json:"createdDate"`
			DisplayName      string `json:"displayName"`
			ExpiresAt        int    `json:"expiresAt"`
			LastUpdatedBy    string `json:"lastUpdatedBy"`
			LastUpdatedDate  string `json:"lastUpdatedDate"`
			Name             string `json:"name"`
			OrganizationLink string `json:"organizationLink"`
			Resource         string `json:"resource"`
			Visible          bool   `json:"visible"`
		} `json:"orgRoles"`
		ServicesDef []struct {
			RefLink string `json:"refLink"`
			Roles   []struct {
				CreatedBy       string `json:"createdBy"`
				CreatedDate     string `json:"createdDate"`
				ExpiresAt       int64  `json:"expiresAt"`
				LastUpdatedBy   string `json:"lastUpdatedBy"`
				LastUpdatedDate string `json:"lastUpdatedDate"`
				MembershipType  string `json:"membershipType"`
				Name            string `json:"name"`
				Resource        string `json:"resource"`
			} `json:"roles"`
			ServiceDisplayName string   `json:"serviceDisplayName"`
			ServiceName        string   `json:"serviceName"`
			ServiceRoles       []string `json:"serviceRoles"`
		} `json:"servicesDef"`
	} `json:"userOrgInfo"`
}

type UserGroupsResponse struct {
	Groups []struct {
		Description  string   `json:"description"`
		DisplayName  string   `json:"displayName"`
		Domain       string   `json:"domain"`
		GroupType    string   `json:"groupType"`
		ID           string   `json:"id"`
		OwnerOrgID   string   `json:"ownerOrgId"`
		SharedOrgIds []string `json:"sharedOrgIds"`
		UsersCount   int      `json:"usersCount"`
	} `json:"groups"`
}

type RoleDto struct {
	CreatedBy        string `json:"createdBy"`
	CreatedDate      string `json:"createdDate"`
	DisplayName      string `json:"displayName"`
	ExpiresAt        int    `json:"expiresAt"`
	LastUpdatedBy    string `json:"lastUpdatedBy"`
	LastUpdatedDate  string `json:"lastUpdatedDate"`
	Name             string `json:"name"`
	OrganizationLink string `json:"organizationLink"`
	Resource         string `json:"resource"`
	Visible          bool   `json:"visible"`
}

type RolesDto []RoleDto

type UserServiceRolesResponse struct {
	ServiceRoles []struct {
		ServiceDefinitionLink string   `json:"serviceDefinitionLink"`
		ServiceRoleNames      []string `json:"serviceRoleNames"`
		ServiceRoles          []struct {
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

type UserMfaStatusDto struct {
	Activated  bool   `json:"activated"`
	DeviceName string `json:"deviceName"`
	TurnedOn   bool   `json:"turnedOn"`
	Username   string `json:"username"`
}

type UserCustomRolesResponse struct {
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
}
