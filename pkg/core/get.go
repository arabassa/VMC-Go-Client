package core

import (
	"reflect"
	utils "vmc-go-client/pkg/utils"

	"fmt"
)

//Get will GET a factory API resource identified by its friendly name returning an empty interface which can be asserted to its struct type
func Get(rf ResourceFactory, s string) interface{} {
	u, t := FindResource(s, rf)

	r := reflect.New(t.Elem()).Interface()

	_, body := utils.DoGetHttp(u.String())

	if jsonerr := utils.Unmarshal([]byte(body), &r); jsonerr != nil {
		fmt.Println("Unmarshalling error: ", jsonerr)
	}
	return r
}
