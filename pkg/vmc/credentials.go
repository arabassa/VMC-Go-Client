package vmc

const Credentials_GET_ADDON_TYPE_HCX = "HCX"

type NewCredential struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type NewCredentials []NewCredential
