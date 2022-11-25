package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	libraryVersion = "0.1"
	UserAgent      = "vmc-go-client/" + libraryVersion
	mediaType      = "application/json"
)

// Client struct to define connection details and credentials
type Client struct {
	CspUrl       string `json:"csp_url"`
	Port         string `json:"port"`
	RefreshToken string `json:"refreshtoken"`
	AccessToken  string `json:"accesstoken"`
	Log          string `json:"log"`
	Client       *http.Client
}

// NewClient creates a new API Client
func NewClient(clientConfig Client) *Client {

	if clientConfig.Client == nil {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}
		clientConfig.Client = &http.Client{Transport: tr}
	}

	c := &Client{
		CspUrl:       clientConfig.CspUrl,
		Port:         clientConfig.Port,
		RefreshToken: clientConfig.RefreshToken,
		AccessToken:  clientConfig.AccessToken,
		Log:          clientConfig.Log,
		Client:       clientConfig.Client}

	return c
}

// LoadConfig loads Client config from json file
func LoadConfig(filename string) *Client {

	var appConfig Client

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(bytes, &appConfig)
	if err != nil {
		fmt.Println(err)
	}
	return &appConfig
}
