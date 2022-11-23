package utils

import (
	"flag"
	"os"
	"strings"
)

var apis = "[CSP, VMC, VCDR]"
var methods = "[LIST, GET, POST, DELETE]"

var (
	api      = flag.String("api", "", "API type: "+apis)
	method   = flag.String("method", "", "List API resources or HTTP method: "+methods)
	resource = flag.String("resource", "", "Resource type")
)

func RunCli() map[string]string {

	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()

	}

	var cliConf = map[string]string{
		"api":      strings.ToUpper(*api),
		"method":   strings.ToUpper(*method),
		"resource": *resource,
	}
	return cliConf
}
