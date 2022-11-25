package main

import (
	"vmc-go-client/pkg/csp"
	"vmc-go-client/pkg/interactive"
	"vmc-go-client/pkg/vcdr"
	"vmc-go-client/pkg/vmc"
)

var buildfactories = false
var climode = true

func main() {

	if buildfactories {
		vmc.BuildConfig()
		csp.BuildConfig()
		vcdr.BuildConfig()
	}

	if climode {
		interactive.RunCli()
	}

}
