package core

import (
	utils "vmc-go-client/pkg/utils"
)

// Post will POST to a factory API endpoint identified by its friendly name and will also send JSON body data if defined
func Post(rf ResourceFactory, s string, data []byte) interface{} {
	u, _ := FindResource(s, rf, false)

	_, body := utils.DoPostHttp(u.String(), data)

	return body
}
