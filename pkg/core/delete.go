package core

import (
	utils "vmc-go-client/pkg/utils"
)

// Delete will DELETE to a factory API endpoint identified by its friendly name
func Delete(rf ResourceFactory, s string) interface{} {
	u, _ := FindResource(s, rf, true, "DELETE")

	_, body := utils.DoDeleteHttp(u.String())

	return body
}
