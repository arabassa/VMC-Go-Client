package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// DoGetHttp authenticates against CSP and performs HTTP request returning the response and the body
func DoGetHttp(url string) (*http.Response, string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("New Request Failed: %s", err)
		return nil, ""
	}

	cli := Auth()
	addHttpHeader(req, "csp-auth-token", cli.AccessToken)          //CSP
	addHttpHeader(req, "authorization", "Bearer "+cli.AccessToken) //some VMC implements
	addHttpHeader(req, "x-da-access-token", cli.AccessToken)       //VCDR

	resp, err := cli.Client.Do(req)
	log.Println("HTTP GET " + url)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return nil, ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 400 {
		fmt.Println(string(body))
		panic("Non 2XX HTTP Response: Exiting API Client.")
	}
	return resp, string(body)
}

// DoPostHttp authenticates against CSP and performs HTTP request returning the response and the body
func DoPostHttp(url string, data []byte) (*http.Response, string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("New Request Failed: %s", err)
		return nil, ""
	}

	fmt.Println(req.Body)
	cli := Auth()
	addHttpHeader(req, "content-type", "application/json")         //media content-type
	addHttpHeader(req, "csp-auth-token", cli.AccessToken)          //CSP
	addHttpHeader(req, "authorization", "Bearer "+cli.AccessToken) //some VMC implements
	addHttpHeader(req, "x-da-access-token", cli.AccessToken)       //VCDR

	resp, err := cli.Client.Do(req)
	log.Println("HTTP POST " + url + "\n" + "JSON Data:")
	log.Println(string(data))

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return nil, ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 400 {
		fmt.Println(string(body))
		panic("Non 2XX HTTP Response: Exiting API Client.")
	}
	return resp, string(body)
}

// addHttpHeader is a helper method to easily add HTTP headers to a request
func addHttpHeader(r *http.Request, k string, v string) *http.Request {
	r.Header.Set(k, v)
	return r
}
