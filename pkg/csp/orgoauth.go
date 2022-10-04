package csp

type PagedResponseOrgOAuthAppResponse struct {
	NextLink string `json:"nextLink"`
	PrevLink string `json:"prevLink"`
	Results  []struct {
		AccessTokenTTL                int      `json:"accessTokenTTL"`
		AdditionalAttributeMasks      []string `json:"additionalAttributeMasks"`
		AllowOpenRedirectUris         bool     `json:"allowOpenRedirectUris"`
		AllowedActorsAudienceExchange []string `json:"allowedActorsAudienceExchange"`
		AllowedActorsClientDelegate   []string `json:"allowedActorsClientDelegate"`
		AllowedOrgs                   []struct {
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
			Name        string `json:"name"`
		} `json:"allowedOrgs"`
		AllowedScopes struct {
			GeneralScopes      []string `json:"generalScopes"`
			OrganizationScopes struct {
				AllPermissions bool     `json:"allPermissions"`
				AllRoles       bool     `json:"allRoles"`
				KeptInToken    []string `json:"keptInToken"`
				Permissions    []struct {
					PermissionID string   `json:"permissionId"`
					Resources    []string `json:"resources"`
				} `json:"permissions"`
				Roles []struct {
					Name     string `json:"name"`
					Resource string `json:"resource"`
				} `json:"roles"`
			} `json:"organizationScopes"`
			ServicesScopes []struct {
				AllPermissions bool     `json:"allPermissions"`
				AllRoles       bool     `json:"allRoles"`
				KeptInToken    []string `json:"keptInToken"`
				Permissions    []struct {
					PermissionID string   `json:"permissionId"`
					Resources    []string `json:"resources"`
				} `json:"permissions"`
				Roles []struct {
					Name     string `json:"name"`
					Resource string `json:"resource"`
				} `json:"roles"`
				ServiceDefinitionID string `json:"serviceDefinitionId"`
			} `json:"servicesScopes"`
		} `json:"allowedScopes"`
		CreatedAt                         int      `json:"createdAt"`
		CreatedBy                         string   `json:"createdBy"`
		Description                       string   `json:"description"`
		DisplayName                       string   `json:"displayName"`
		ForcePkce                         bool     `json:"forcePkce"`
		GrantTypes                        []string `json:"grantTypes"`
		GroupDomainAppendedInIDToken      bool     `json:"groupDomainAppendedInIDToken"`
		ID                                string   `json:"id"`
		Immutable                         bool     `json:"immutable"`
		LastUpdatedAt                     int      `json:"lastUpdatedAt"`
		LastUpdatedBy                     string   `json:"lastUpdatedBy"`
		MaxAdditionalAttributesInIDToken  int      `json:"maxAdditionalAttributesInIdToken"`
		MaxCharactersInAccessToken        int      `json:"maxCharactersInAccessToken"`
		MaxGroupsInIDToken                int      `json:"maxGroupsInIdToken"`
		OrganizationID                    string   `json:"organizationId"`
		OwnerOnlySecretRotation           bool     `json:"ownerOnlySecretRotation"`
		PostLogoutRedirectUris            []string `json:"postLogoutRedirectUris"`
		PublicClient                      bool     `json:"publicClient"`
		RedirectUris                      []string `json:"redirectUris"`
		RefreshTokenTTL                   int      `json:"refreshTokenTTL"`
		SecretRotationExpirationInSeconds int      `json:"secretRotationExpirationInSeconds"`
		ServiceDefinitionID               string   `json:"serviceDefinitionId"`
	} `json:"results"`
	TotalResults int `json:"totalResults"`
}

type OrgOAuthAppResponse struct {
	AccessTokenTTL                int      `json:"accessTokenTTL"`
	AdditionalAttributeMasks      []string `json:"additionalAttributeMasks"`
	AllowOpenRedirectUris         bool     `json:"allowOpenRedirectUris"`
	AllowedActorsAudienceExchange []string `json:"allowedActorsAudienceExchange"`
	AllowedActorsClientDelegate   []string `json:"allowedActorsClientDelegate"`
	AllowedOrgs                   []struct {
		DisplayName string `json:"displayName"`
		ID          string `json:"id"`
		Name        string `json:"name"`
	} `json:"allowedOrgs"`
	AllowedScopes struct {
		GeneralScopes      []string `json:"generalScopes"`
		OrganizationScopes struct {
			AllPermissions bool     `json:"allPermissions"`
			AllRoles       bool     `json:"allRoles"`
			KeptInToken    []string `json:"keptInToken"`
			Permissions    []struct {
				PermissionID string   `json:"permissionId"`
				Resources    []string `json:"resources"`
			} `json:"permissions"`
			Roles []struct {
				Name     string `json:"name"`
				Resource string `json:"resource"`
			} `json:"roles"`
		} `json:"organizationScopes"`
		ServicesScopes []struct {
			AllPermissions bool     `json:"allPermissions"`
			AllRoles       bool     `json:"allRoles"`
			KeptInToken    []string `json:"keptInToken"`
			Permissions    []struct {
				PermissionID string   `json:"permissionId"`
				Resources    []string `json:"resources"`
			} `json:"permissions"`
			Roles []struct {
				Name     string `json:"name"`
				Resource string `json:"resource"`
			} `json:"roles"`
			ServiceDefinitionID string `json:"serviceDefinitionId"`
		} `json:"servicesScopes"`
	} `json:"allowedScopes"`
	CreatedAt                         int      `json:"createdAt"`
	CreatedBy                         string   `json:"createdBy"`
	Description                       string   `json:"description"`
	DisplayName                       string   `json:"displayName"`
	ForcePkce                         bool     `json:"forcePkce"`
	GrantTypes                        []string `json:"grantTypes"`
	GroupDomainAppendedInIDToken      bool     `json:"groupDomainAppendedInIDToken"`
	ID                                string   `json:"id"`
	Immutable                         bool     `json:"immutable"`
	LastUpdatedAt                     int      `json:"lastUpdatedAt"`
	LastUpdatedBy                     string   `json:"lastUpdatedBy"`
	MaxAdditionalAttributesInIDToken  int      `json:"maxAdditionalAttributesInIdToken"`
	MaxCharactersInAccessToken        int      `json:"maxCharactersInAccessToken"`
	MaxGroupsInIDToken                int      `json:"maxGroupsInIdToken"`
	OrganizationID                    string   `json:"organizationId"`
	OwnerOnlySecretRotation           bool     `json:"ownerOnlySecretRotation"`
	PostLogoutRedirectUris            []string `json:"postLogoutRedirectUris"`
	PublicClient                      bool     `json:"publicClient"`
	RedirectUris                      []string `json:"redirectUris"`
	RefreshTokenTTL                   int      `json:"refreshTokenTTL"`
	SecretRotationExpirationInSeconds int      `json:"secretRotationExpirationInSeconds"`
	ServiceDefinitionID               string   `json:"serviceDefinitionId"`
}
