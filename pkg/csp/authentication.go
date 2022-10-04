package csp

type PublicKeyDto struct {
	Alg    string `json:"alg"`
	Issuer string `json:"issuer"`
	Keys   struct {
	} `json:"keys"`
	Value string `json:"value"`
}
