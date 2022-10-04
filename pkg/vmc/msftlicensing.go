package vmc

type TermsOfServiceResult struct {
	Signed  bool   `json:"signed"`
	TermsID string `json:"terms_id"`
}
