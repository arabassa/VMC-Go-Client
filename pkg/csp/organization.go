package csp

type OrganizationDto struct {
	AuditLogsInstanceID             string `json:"auditLogsInstanceId"`
	DisplayName                     string `json:"displayName"`
	EnforceUserAPITokenMfa          bool   `json:"enforceUserApiTokenMfa"`
	IsMfaRequired                   bool   `json:"isMfaRequired"`
	Language                        string `json:"language"`
	Locale                          string `json:"locale"`
	MaxAllowedAuthExemptedUserCount int    `json:"maxAllowedAuthExemptedUserCount"`
	Metadata                        struct {
		Key string `json:"key"`
	} `json:"metadata"`
	Name          string `json:"name"`
	ParentRefLink string `json:"parentRefLink"`
	RefLink       string `json:"refLink"`
}
type SearchGroupsResponse struct {
	Results []struct {
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
}

type OrgLinkedIdpResponse struct {
	IdpDetails struct {
		IdpDisplayName string   `json:"idpDisplayName"`
		IdpDomains     []string `json:"idpDomains"`
		IdpID          string   `json:"idpId"`
	} `json:"idpDetails"`
	MappedToIdp bool `json:"mappedToIdp"`
}

type PagedResponseExpandedAuthClientDto struct {
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
		SourceOrg struct {
			DisplayName string `json:"displayName"`
			OrgID       string `json:"orgId"`
			ShortID     string `json:"shortId"`
		} `json:"sourceOrg"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}

type UserOrganizationInvitationResponse struct {
	CustomGroups []struct {
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
	} `json:"customGroups"`
	CustomGroupsIds []string `json:"customGroupsIds"`
	CustomRoles     []struct {
		CreatedBy       string `json:"createdBy"`
		CreatedDate     string `json:"createdDate"`
		ExpiresAt       int64  `json:"expiresAt"`
		LastUpdatedBy   string `json:"lastUpdatedBy"`
		LastUpdatedDate string `json:"lastUpdatedDate"`
		MembershipType  string `json:"membershipType"`
		Name            string `json:"name"`
		Resource        string `json:"resource"`
	} `json:"customRoles"`
	ExpirationTime    int      `json:"expirationTime"`
	GeneratedAt       int      `json:"generatedAt"`
	GeneratedBy       string   `json:"generatedBy"`
	InvitationLink    string   `json:"invitationLink"`
	InvitedByUsername string   `json:"invitedByUsername"`
	OrgRoleNames      []string `json:"orgRoleNames"`
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
	RedeemedAt       int    `json:"redeemedAt"`
	RedeemedBy       string `json:"redeemedBy"`
	RefLink          string `json:"refLink"`
	RevokedAt        int    `json:"revokedAt"`
	RevokedBy        string `json:"revokedBy"`
	ServiceRolesDtos []struct {
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
	} `json:"serviceRolesDtos"`
	Status   string `json:"status"`
	Username string `json:"username"`
}

type OrganizationRolesResponse struct {
	OrgRolesData []struct {
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
	} `json:"orgRolesData"`
	RefLinks []string `json:"refLinks"`
}

type PagedResponseOrganizationTrustDto struct {
	NextLink string `json:"nextLink"`
	PrevLink string `json:"prevLink"`
	Results  []struct {
		AllowedScopes struct {
			AllScopes          bool `json:"allScopes"`
			OrganizationScopes struct {
				AllRoles bool `json:"allRoles"`
				Roles    []struct {
					Name      string   `json:"name"`
					Resources []string `json:"resources"`
				} `json:"roles"`
			} `json:"organizationScopes"`
			ServicesScopes []struct {
				AllRoles bool `json:"allRoles"`
				Roles    []struct {
					Name      string   `json:"name"`
					Resources []string `json:"resources"`
				} `json:"roles"`
				ServiceDefinitionID string `json:"serviceDefinitionId"`
			} `json:"servicesScopes"`
		} `json:"allowedScopes"`
		ExpiresAt  int    `json:"expiresAt"`
		Status     string `json:"status"`
		TrustID    string `json:"trustId"`
		TrustedOrg struct {
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
			Name        string `json:"name"`
		} `json:"trustedOrg"`
		TrusteeOrg struct {
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
			Name        string `json:"name"`
		} `json:"trusteeOrg"`
		Type string `json:"type"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}

type PagedResponseExpandedTypedUserDto struct {
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
		OrgID             string `json:"orgId"`
		OrganizationRoles []struct {
			CreatedBy   string   `json:"createdBy"`
			CreatedDate string   `json:"createdDate"`
			DisplayName string   `json:"displayName"`
			ExpiresAt   int64    `json:"expiresAt"`
			GroupIds    []string `json:"groupIds"`
			Groups      []struct {
				Description  string   `json:"description"`
				DisplayName  string   `json:"displayName"`
				Domain       string   `json:"domain"`
				GroupType    string   `json:"groupType"`
				ID           string   `json:"id"`
				OwnerOrgID   string   `json:"ownerOrgId"`
				SharedOrgIds []string `json:"sharedOrgIds"`
				UsersCount   int      `json:"usersCount"`
			} `json:"groups"`
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
				CreatedBy   string   `json:"createdBy"`
				CreatedDate string   `json:"createdDate"`
				ExpiresAt   int64    `json:"expiresAt"`
				GroupIds    []string `json:"groupIds"`
				Groups      []struct {
					Description  string   `json:"description"`
					DisplayName  string   `json:"displayName"`
					Domain       string   `json:"domain"`
					GroupType    string   `json:"groupType"`
					ID           string   `json:"id"`
					OwnerOrgID   string   `json:"ownerOrgId"`
					SharedOrgIds []string `json:"sharedOrgIds"`
					UsersCount   int      `json:"usersCount"`
				} `json:"groups"`
				LastUpdatedBy   string `json:"lastUpdatedBy"`
				LastUpdatedDate string `json:"lastUpdatedDate"`
				MembershipType  string `json:"membershipType"`
				Name            string `json:"name"`
				Resource        string `json:"resource"`
			} `json:"serviceRoles"`
		} `json:"serviceRoles"`
		User struct {
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
		} `json:"user"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}

type PagedResponseOrganizationDto struct {
	NextLink string `json:"nextLink"`
	PrevLink string `json:"prevLink"`
	Results  []struct {
		AuditLogsInstanceID             string `json:"auditLogsInstanceId"`
		DisplayName                     string `json:"displayName"`
		EnforceUserAPITokenMfa          bool   `json:"enforceUserApiTokenMfa"`
		IsMfaRequired                   bool   `json:"isMfaRequired"`
		Language                        string `json:"language"`
		Locale                          string `json:"locale"`
		MaxAllowedAuthExemptedUserCount int    `json:"maxAllowedAuthExemptedUserCount"`
		Metadata                        struct {
			Key string `json:"key"`
		} `json:"metadata"`
		Name          string `json:"name"`
		ParentRefLink string `json:"parentRefLink"`
		RefLink       string `json:"refLink"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}

type PagedResponseOrganizationMemberForTrustDto struct {
	NextLink string `json:"nextLink"`
	PrevLink string `json:"prevLink"`
	Results  []struct {
		DisplayName       string `json:"displayName"`
		ID                string `json:"id"`
		Name              string `json:"name"`
		OrganizationTrust struct {
			AllowedScopes struct {
				AllScopes          bool `json:"allScopes"`
				OrganizationScopes struct {
					AllRoles bool `json:"allRoles"`
					Roles    []struct {
						Name      string   `json:"name"`
						Resources []string `json:"resources"`
					} `json:"roles"`
				} `json:"organizationScopes"`
				ServicesScopes []struct {
					AllRoles bool `json:"allRoles"`
					Roles    []struct {
						Name      string   `json:"name"`
						Resources []string `json:"resources"`
					} `json:"roles"`
					ServiceDefinitionID string `json:"serviceDefinitionId"`
				} `json:"servicesScopes"`
			} `json:"allowedScopes"`
			ExpiresAt  int    `json:"expiresAt"`
			Status     string `json:"status"`
			TrustID    string `json:"trustId"`
			TrustedOrg struct {
				DisplayName string `json:"displayName"`
				ID          string `json:"id"`
				Name        string `json:"name"`
			} `json:"trustedOrg"`
			TrusteeOrg struct {
				DisplayName string `json:"displayName"`
				ID          string `json:"id"`
				Name        string `json:"name"`
			} `json:"trusteeOrg"`
			Type string `json:"type"`
		} `json:"organizationTrust"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}
