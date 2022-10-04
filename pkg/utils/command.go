package utils

import (
	"flag"
	"os"
	"strings"
)

var (
	api      = flag.String("api", "", "API type")
	method   = flag.String("method", "", "HTTP method")
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
