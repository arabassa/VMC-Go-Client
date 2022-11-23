package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

type Endpoint struct {
	Path        string
	Method      []string
	Resource    reflect.Type      `json:"-"`
	QueryParams map[string]string `json:",omitempty"`
}

type ResourceFactory struct {
	ApiName   string
	ApiScheme string
	ApiUrl    string
	Resources map[string]Endpoint
	Params    map[string]string
	Get       func(string) interface{}
	Post      func(string, []byte) interface{}
}

/*
FindResource receives an API resource friendly name and returns the constructed link with needed variables/IDs and queryparams.
It will also return the reflect Type of the interface so the GET response can be asserted and the correspondent struct data accessed
*/
func FindResource(r string, rf ResourceFactory, enforce bool) (*url.URL, reflect.Type) {
	var e Endpoint

	for name, values := range rf.Resources {
		if name == r {
			e = values
		}
	}

	if e.Resource == nil {
		panic("Exiting API Client. Unknown resource type: " + r)
	}

	u := url.URL{
		Scheme: rf.ApiScheme,
		Host:   rf.ApiUrl,
		Path:   fillUrl(e.Path, rf.Params),
	}

	if enforce {
		validateUrl(u.String(), rf)
	}

	n := addQueryParams(&u, r, rf)

	return n, e.Resource
}

// fillUrl substitutes the required API params in the URL
func fillUrl(path string, params map[string]string) string {
	for ids, vals := range params {
		if vals != "" {
			path = strings.Replace(path, ids, vals, 1)
		}
	}
	return path
}

// validateUrl checks that parameters needed for the API URL construction are satisfied
func validateUrl(s string, rf ResourceFactory) {
	var sb strings.Builder

	//re-encoding "{" and "}" to check if path paramaters in original format (i.e. {orgId})
	s = strings.Replace(s, "%7B", "{", -1)
	s = strings.Replace(s, "%7D", "}", -1)

	for i := range rf.Params {
		sb.WriteString(i + "|")
	}

	str := sb.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}

	rs := regexp.MustCompile(str)

	missing := rs.FindAllString(s, -1)
	for i := range missing {
		fmt.Println("Missing parameter for API Call: " + missing[i])
	}

	if len(missing) > 0 {
		fmt.Println("Review or update your API Factory json for missing Params")
		panic("Error: Exiting API Client.")
	}
}

// addQueryParams will add trailing query params to the URL based on the resource type
func addQueryParams(u *url.URL, t string, rf ResourceFactory) *url.URL {
	newUrl := *u
	values := newUrl.Query()

	for k, v := range rf.Resources[t].QueryParams {
		values.Add(k, v)
	}

	newUrl.RawQuery = values.Encode()
	return &newUrl
}

// LoadFactory will load the factory struct from the JSON file and make it available to process all the defined resources
func LoadFactory(filename string, resources map[string]reflect.Type) ResourceFactory {
	var rf ResourceFactory

	//factory Getter function
	rf.Get = func(r string) interface{} { return Get(rf, r) }
	//factory Poster function, includes JSON data for the body
	rf.Post = func(r string, d []byte) interface{} { return Post(rf, r, d) }

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error loading config file ", filename, " : ", err)
	}

	err = json.Unmarshal(bytes, &rf)
	if err != nil {
		fmt.Println("Error parsing JSON factory config file ", filename, " : ", err)
		fmt.Println("You can use the BuildConfig() methods of the APIs to regenerate factory config.")
	}

	for k, v := range resources {
		entry := rf.Resources[k]
		entry.Resource = v
		rf.Resources[k] = entry
	}

	return rf
}

// ListMethods is a helper method for pretty print available commands of an API/factory in the cli
func ListMethods(rf ResourceFactory) {
	fmt.Println("Method", "    Resource")
	fmt.Println("======     ===================")
	for a, b := range rf.Resources {
		var params string
		re := regexp.MustCompile(`{([^}]+)}`)
		match := re.FindAllStringSubmatch(b.Path, -1)
		fmt.Println(b.Method, "    ", a)
		for _, i := range match {
			params += i[0] + " "
		}
		if len(params) > 0 {
			fmt.Println("Params:   ", params)
		}
		fmt.Println()
	}
}
