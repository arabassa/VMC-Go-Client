package core

import (
	utils "vmc-go-client/pkg/utils"
)

// Delete will DELETE to a factory API endpoint identified by its friendly name and will also send JSON body data if defined
func Delete(rf ResourceFactory, s string) interface{} {
	u, _ := FindResource(s, rf, true)

	_, body := utils.DoDeleteHttp(u.String())

	return body
}
