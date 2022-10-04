package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	loginPath   = "/csp/gateway/am/api/auth/api-tokens/authorize?refresh_token="
	contentType = "application/json"
)

//ApiTokens defines the tokens model returned by CSP auth API
type ApiTokens struct {
	IdToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	Expiry       int    `json:"expires_in"`
	Scope        string `json:"scope"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//Auth loads config to client and logins to the API
func Auth() *Client {
	cli := NewClient(*LoadConfig("config.json"))
	return Login(cli)
}

//Login connects to VMware CSP with the refresh-token and sets the access-token in the client
func Login(cli *Client) *Client {
	var apitokens ApiTokens
	url := cli.CspUrl + loginPath + cli.RefreshToken

	resp, err := cli.Client.Post(url, contentType, nil)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
		return nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	json.Unmarshal([]byte(b), &apitokens)

	log.Println("API Login HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	cli.AccessToken = apitokens.AccessToken
	return cli
}
