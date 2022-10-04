package utils

import (
	"encoding/json"
	"fmt"
)

//Unmarshall Json data to struct
func Unmarshal(body []byte, c interface{}) error {
	return json.Unmarshal(body, c)
}

//Marshall struct to Json
func Marshal(c interface{}) string {
	json, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	return string(json)
}
