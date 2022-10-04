package csp

type OidcUserInfoDto struct {
	Acct          string   `json:"acct"`
	Context       string   `json:"context"`
	ContextName   string   `json:"context_name"`
	Domain        string   `json:"domain"`
	Email         string   `json:"email"`
	EmailVerified bool     `json:"email_verified"`
	FamilyName    string   `json:"family_name"`
	GivenName     string   `json:"given_name"`
	GroupIds      []string `json:"group_ids"`
	GroupNames    []string `json:"group_names"`
	Sub           string   `json:"sub"`
	Username      string   `json:"username"`
}

type CspOpenIdConfigurationDto struct {
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	ClaimsSupported                   []string `json:"claimsSupported"`
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported"`
	EndSessionEndpoint                string   `json:"end_session_endpoint"`
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	Issuer                            string   `json:"issuer"`
	JwksURI                           string   `json:"jwks_uri"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	ScopesSupported                   []string `json:"scopes_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
}
