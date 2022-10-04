
#  VMC-Go-Client

VMC-Go-Client is a Go developed client to interact with VMware Cloud APIs. A command-line client is included for easy access. You can also work directly with the different data types with minimal coding.

In this initial release v0.1 client is read-only (only GET operations available) for these 3 APIs:
- VMware CSP (https://developer.vmware.com/apis/csp/)
- VMC on AWS (https://developer.vmware.com/apis/vmc/latest/)
- VMware VCDR (https://developer.vmware.com/apis/vcdr/latest/)

# Pre-requisites

You'll need Go installed in your operating system. VMC-Go-Client has been tested with Go versions **1.18.3** and **1.19.1**.

## Installation

- Clone the repository
```
git clone https://github.com/arabassa/VMC-Go-Client.git
```
- Build the package
```
go build
```
- Run client
```
./vmc-go-client
```

# Usage and examples
## Built-in command line

By default the client will run with the command cli mode:
```
% ./vmc-go-client
Usage of ./vmc-go-client:
-api string
        API type: [CSP, VMC, VCDR]
-method string
        Method:   [GET, LIST]
-resource string
       Resource type
```

Run the client with method `list` for the Client to print all the methods for an API:

```
% ./vmc-go-client -api VCDR -method LIST
Method   Resource
======   ===================
GET      ProtectedSites
Params:  {cloud_file_system_id} 

GET      ProtectionGroupSnapshots
Params:  {cloud_file_system_id} {protection_group_id} 

GET      ProtectionGroups
Params:  {cloud_file_system_id} 

GET      ProtectionGroupSnapshot
Params:  {cloud_file_system_id} {protection_group_id} {snapshot_id} 

GET      RecoverySddcDetails
Params:  {recovery_sddc_id} 

GET      RecoverySddcs

GET      CloudFileSystem
Params:  {cloud_file_system_id} 

GET      CloudFileSystems

GET      ProtectedSite
Params:  {cloud_file_system_id} {protected_site_id} 

GET      ProtectedVms
Params:  {cloud_file_system_id} 

GET      ProtectionGroup
Params:  {cloud_file_system_id} {protection_group_id} 
```

Before calling the APIs for resources you'll need set your CSP Refresh Token in the `config.json` file:

```
{
"csp_url": "https://console.cloud.vmware.com",
"refreshtoken": "plht0b_myFakeRefreshToken_vjYJrAadf2c6cTFS"
}
```

Now you can GET a resource with the command cli:
```
% ./vmc-go-client -api VCDR -method GET -resource CloudFileSystems
2022/10/04 17:50:55 API Login HTTP Response Status: 200 OK
2022/10/04 17:50:55 HTTP GET https://vcdr-3-xxx-xxx-xxx.app.vcdr.vmware.com/api/vcdr/v1alpha/cloud-file-systems
2022/10/04 17:50:55 HTTP Response Status: 200 OK
{
  "cloud_file_systems": [
    {
      "id": "g7h17yft-masked-cloud-fs-90a7-f17542a362f25",
      "name": "cloud-file-system-01"
    }
  ]
}
```

Many API calls will need parameters to fully construct the URL, for instance getting information of a VCDR Cloud File System:

```
% ./vmc-go-client -api VCDR -method GET -resource CloudFileSystem
Missing parameter for API Call: {cloud_file_system_id}
Review or update your API Factory json for missing Params
panic: Error: Exiting API Client.
```

You can see with the List method all the parameters needed for each API resource. The client also informs you, in this case we need to provide `{cloud_file_system_id}`.

This has to be setup for the API you're talking to in the correspondent factory JSON. We have the Cloud File System ID from the previous call, so in this case we will edit `factory-vcdr.json`, adding it to the Params section in the end:
```
"Params": {
  "{cloud_file_system_id}": "g7h17yft-masked-cloud-fs-90a7-f17542a362f25",
  "{protected_site_id}": "",
  "{protection_group_id}": "",
  "{recovery_sddc_id}": "",
  "{snapshot_id}": ""
}
```

Now I can repeat the previous execution:

```
% ./vmc-go-client -api VCDR -method GET -resource CloudFileSystem 
2022/10/04 17:53:12 API Login HTTP Response Status: 200 OK
2022/10/04 17:53:13 HTTP GET https://vcdr-3-212-171-179.app.vcdr.vmware.com/api/vcdr/v1alpha/cloud-file-systems/g7h17yft-masked-cloud-fs-90a7-f17542a362f25
2022/10/04 17:53:13 HTTP Response Status: 200 OK
{
  "capacity_gib": 107928.93593750056,
  "id": "g7h17yft-masked-cloud-fs-90a7-f17542a362f25",
  "name": "cloud-file-system-01",
  "recovery_sddc_id": "bb8ggbb6-sddc-masked-a8f2-fg8ba0dff2f",
  "used_gib": 56.8125
}
```

## Automating or creating your own functions

You can set variable `climode = false` in `main.go` to disable the CLI mode and add your custom behaviour or just write your own main file or functions. 

All implemented API calls have their correspondent structs to Marshal and Unmarshal the API JSON data so you can assert all the responses with the correspondent Struct types:

```
factory := core.LoadFactory("factory-vcdr.json", vcdr.VCDRResourceMap)
object := factory.Get("CloudFileSystem")
fmt.Println("raw print of the API response")
fmt.Print(object, "\n")
fmt.Println("raw print of the API response with struct fields")
fmt.Printf("%+v\n", object)
fmt.Println("marshalled print of the API response")
fmt.Println(utils.Marshal(object))
fmt.Println("assert struct type to access specific data, get only UsedGib from Cloud File System details")
fmt.Println(object.(*vcdr.CloudFileSystemDetails).UsedGib)
```  

Generates: 
```
raw print of the API response:
&{107928.93593750056 g7h17yft-masked-cloud-fs-90a7-f17542a362f25 cloud-file-system-01 ab478bb6-c6db-47e1-a8f2-feeb8ba0bf2f 59.5546875}
raw print of the API response with struct fields:
&{CapacityGib:107928.93593750056 ID:g7h17yft-masked-cloud-fs-90a7-f17542a362f25 Name:cloud-file-system-01 RecoverySddcID:bb8ggbb6-sddc-masked-a8f2-fg8ba0dff2f UsedGib:59.5546875}
marshalled print of the API response:
{
  "capacity_gib": 107928.93593750056,
  "id": "g7h17yft-masked-cloud-fs-90a7-f17542a362f25",
  "name": "cloud-file-system-01",
  "recovery_sddc_id": "bb8ggbb6-sddc-masked-a8f2-fg8ba0dff2f",
  "used_gib": 59.5546875
}
assert struct type to access specific data, get only UsedGib from Cloud File System details:
59.5546875
```

## Factories
Each API has its own `factory-api-name.json`file which contains all the API paths and definitions and also placeholders for the different params. These files are dinamycally generated with the Resource files from each API and additions of more APIs to the client can be easily done this way.

When using the client, edits on the JSON file params are expected to either put valid URL of a VCDR endpoint or to pass the different needed parameters for the API calls.

If you need at any point the regenerate the JSON file this can be done automatically setting the`buildafactories` variable to `true` on main. It will print the JSON file for all the APIs. Alternatively, you can just use the `BuildConfig()` methods of the different APIs.


# Design thoughts
This client has been developed with some considerations:
- Primary goal of this project is to be a learning journey for anyone interested in the VMC APIs
- Easy automation of tasks
- Use of the standard Go libraries, no external dependencies added
- Facilitate adoption/addition of new APIs

## To-Do
- Implement remaining HTTP Methods (PUT, POST, DELETE)
- Update or add new resources based on the new HTTP methods
- Add new APIs
- Refactor, clean, improve

## Where can I find  additional documentation of VMC-Go-Client?

A series of blog posts with examples will be created.
