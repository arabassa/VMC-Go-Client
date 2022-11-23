package main

import (
	"fmt"
	"vmc-go-client/pkg/core"
	"vmc-go-client/pkg/csp"
	"vmc-go-client/pkg/utils"
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
		c := utils.RunCli()

		switch c["api"] {
		case "CSP":
			factory := core.LoadFactory("factory-csp.json", csp.CSPResourceMap)
			switch c["method"] {
			case "GET":
				fmt.Println(utils.Marshal(factory.Get(c["resource"])))
			case "LIST":
				core.ListMethods(factory)
			default:
				fmt.Println("Invalid method for", factory.ApiName, "API")
			}
		case "VCDR":
			factory := core.LoadFactory("factory-vcdr.json", vcdr.VCDRResourceMap)
			switch c["method"] {
			case "GET":
				fmt.Println(utils.Marshal(factory.Get(c["resource"])))
			case "LIST":
				core.ListMethods(factory)
			default:
				fmt.Println("Invalid method for", factory.ApiName, "API")
			}
		case "VMC":
			factory := core.LoadFactory("factory-vmc.json", vmc.VMCResourceMap)
			switch c["method"] {
			case "GET":
				fmt.Println(utils.Marshal(factory.Get(c["resource"])))
			case "POST":
				fmt.Println(utils.Marshal(factory.Post((c["resource"]), utils.LoadJsonFile("post-data.json"))))
			case "DELETE":
				fmt.Println(utils.Marshal(factory.Delete(c["resource"])))
			case "LIST":
				core.ListMethods(factory)
			default:
				fmt.Println("Invalid method for", factory.ApiName, "API")
			}
		}
	}
}
