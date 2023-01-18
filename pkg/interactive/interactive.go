package interactive

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"vmc-go-client/pkg/core"
	"vmc-go-client/pkg/csp"
	"vmc-go-client/pkg/utils"
	"vmc-go-client/pkg/vcdr"
	"vmc-go-client/pkg/vmc"
)

var apis = "[CSP, VMC, VCDR]"
var methods = "[LIST, GET, POST, DELETE]"

func RunCli() {

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	api := flag.String("api", "", "API type: "+apis)
	method := flag.String("method", "", "List API resources or HTTP method: "+methods)
	resource := flag.String("resource", "", "Resource type")

	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
	}

	var cliConf = map[string]string{
		"api":      strings.ToUpper(*api),
		"method":   strings.ToUpper(*method),
		"resource": *resource,
	}
	Interact(cliConf)
}

func Interact(c map[string]string) {
	switch c["api"] {
	case "CSP":
		factory := core.LoadFactory("factory-csp.json", csp.CSPResourceMap)
		switch c["method"] {
		case "GET":
			fmt.Println(utils.Marshal(factory.Get(c["resource"])))
		case "LIST":
			core.ListMethods(factory)
		default:
			fmt.Println("Exiting API Client. Invalid method for", factory.ApiName, "API")
		}
	case "VCDR":
		factory := core.LoadFactory("factory-vcdr.json", vcdr.VCDRResourceMap)
		switch c["method"] {
		case "GET":
			fmt.Println(utils.Marshal(factory.Get(c["resource"])))
		case "LIST":
			core.ListMethods(factory)
		default:
			fmt.Println("Exiting API Client. Invalid method for", factory.ApiName, "API")
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
			fmt.Println("Exiting API Client. Invalid method for", factory.ApiName, "API")
		}
	default:
		fmt.Println("Exiting API Client. Invalid API Type", c["api"])
	}
}
