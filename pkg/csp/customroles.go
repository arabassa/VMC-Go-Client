package csp

type CustomRoleDto struct {
	CreatedBy      string   `json:"createdBy"`
	Description    string   `json:"description"`
	DisplayName    string   `json:"displayName"`
	LastModifiedBy string   `json:"lastModifiedBy"`
	Name           string   `json:"name"`
	Permissions    []string `json:"permissions"`
}

type ResultsDtoCustomRoleDto []CustomRoleDto
