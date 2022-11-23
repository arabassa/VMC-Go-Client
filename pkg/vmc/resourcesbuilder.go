package vmc

import (
	"fmt"
	"vmc-go-client/pkg/core"
	"vmc-go-client/pkg/utils"
)

const name = "VMC"
const scheme = "https"
const url = "vmc.vmware.com"

var readMethods = []string{"GET"}
var allMethods = []string{"GET", "POST", "DELETE", "PATCH"}

var resources = map[string]core.Endpoint{
	"Org":                            {"/vmc/api/orgs/{org}", readMethods, initType(new(Org)), queryparams["Org"]},
	"Orgs":                           {"/vmc/api/orgs", readMethods, initType(new(Orgs)), queryparams["Orgs"]},
	"OrgPaymentMethods":              {"/vmc/api/orgs/{org}/payment-methods", readMethods, initType(new(PaymentMethodInfos)), queryparams["OrgPaymentMethods"]},
	"OrgProviders":                   {"/vmc/api/orgs/{org}/providers", readMethods, initType(new(AwsCloudProviders)), queryparams["OrgProviders"]},
	"Reservations":                   {"/vmc/api/orgs/{org}/reservations", readMethods, initType(new(Reservations)), queryparams["Reservations"]},
	"Sddcs":                          {"/vmc/api/orgs/{org}/sddcs", []string{"GET", "POST"}, initType(new(Sddcs)), queryparams["Sddcs"]},
	"Sddc":                           {"/vmc/api/orgs/{org}/sddcs/{sddc}", []string{"GET", "DELETE"}, initType(new(Sddc)), queryparams["Sddc"]},
	"ProvisionSpec":                  {"/vmc/api/orgs/{org}/sddcs/provision-spec", readMethods, initType(new(ProvisionSpec)), queryparams["ProvisionSpec"]},
	"Cluster":                        {"/vmc/api/orgs/{org}/sddcs/{sddc}/primarycluster", readMethods, initType(new(Cluster)), queryparams["Cluster"]},
	"VsanConfigConstraints":          {"/vmc/api/orgs/{org}/storage/cluster-constraints", readMethods, initType(new(VsanConfigConstraints)), queryparams["VsanConfigConstraints"]},
	"VsanClusterReconfigConstraints": {"/vmc/api/orgs/{org}/sddcs/{sddc}/clusters/{cluster}/config/constraints", readMethods, initType(new(VsanClusterReconfigConstraints)), queryparams["VsanClusterReconfigConstraints"]},
	"StorageSpecs":                   {"/vmc/api/orgs/{org}/sddcs/{sddc}/clusters/{cluster}/esxs/{esx}/storage-spec", readMethods, initType(new(VsanDiskgroupMapping)), queryparams["StorageSpecs"]},
	"ServiceQuotaRequest":            {"/vmc/api/orgs/{org}/aws/resources/servicequotas/requests/{serviceQuotaRequestId}Id", readMethods, initType(new(ServiceQuotaRequest)), queryparams["ServiceQuotaRequest"]},
	"ServiceQuotaRequests":           {"/vmc/api/orgs/{org}/aws/resources/servicequotas/requests", readMethods, initType(new(ServiceQuotaRequests)), queryparams["ServiceQuotaRequests"]},
	"AccountLink":                    {"/vmc/api/orgs/{org}/account-link", readMethods, initType(new(AccountLink)), queryparams["AccountLink"]},
	"ConnectedAccounts":              {"/vmc/api/orgs/{org}/account-link/connected-accounts", readMethods, initType(new(AwsCustomerConnectedAccounts)), queryparams["ConnectedAccounts"]},
	"CompatibleSubnets":              {"/vmc/api/orgs/{org}/account-link/compatible-subnets", readMethods, initType(new(AwsCompatibleSubnets)), queryparams["CompatibleSubnets"]},
	"AwsSddcConnections":             {"/vmc/api/orgs/{org}/account-link/sddc-connections", readMethods, initType(new(AwsSddcConnections)), queryparams["AwsSddcConnections"]},
	"Credential":                     {"/vmc/api/orgs/{org}/sddcs/{sddc}/addons/" + Credentials_GET_ADDON_TYPE_HCX + "/credentials/{name}", readMethods, initType(new(NewCredential)), queryparams["Credential"]},
	"Credentials":                    {"/vmc/api/orgs/{org}/sddcs/{sddc}/addons/" + Credentials_GET_ADDON_TYPE_HCX + "/credentials", readMethods, initType(new(NewCredentials)), queryparams["Credentials"]},
	"CustomerSupport":                {"/vmc/api/customer-support/inputs", readMethods, initType(new(SddcInput)), queryparams["CustomerSupport"]},
	"MsftLicensing":                  {"/vmc/api/orgs/{org}/tos", readMethods, initType(new(SddcInput)), queryparams["MsftLicensing"]},
	"Subscriptions":                  {"/vmc/api/orgs/{org}/subscriptions", readMethods, initType(new(Subscriptions)), queryparams["Subscriptions"]},
	"Subscription":                   {"/vmc/api/orgs/{org}/subscriptions/{subscription}", readMethods, initType(new(Subscription)), queryparams["Subscription"]},
	"OfferInstances":                 {"/vmc/api/orgs/{org}/subscriptions/offer-instances", readMethods, initType(new(OfferInstances)), queryparams["OfferInstances"]},
	"Products":                       {"/vmc/api/orgs/{org}/subscriptions/products", readMethods, initType(new(SubscriptionProducts)), queryparams["Products"]},
	"Tbrs":                           {"/vmc/api/orgs/{org}/tbrs/support-window", readMethods, initType(new(SupportWindows)), queryparams["Tbrs"]},
	"Tasks":                          {"/vmc/api/orgs/{org}/tasks", readMethods, initType(new(Tasks)), queryparams["Tasks"]}, //heavy call, better use it with filter-query params
	"Task":                           {"/vmc/api/orgs/{org}/tasks/{task}", readMethods, initType(new(Task)), queryparams["Task"]},
	"NetworkConnChecker":             {"/vmc/api/orgs/{org}/sddcs/{sddc}/networking/connectivity-tests", readMethods, initType(new(ConnectivityValidationGroup)), queryparams["NetworkConnChecker"]},
}

// params to set variables and construct our API resources
var params = map[string]string{
	"{org}":                   "",
	"{sddc}":                  "",
	"{cluster}":               "",
	"{esx}":                   "",
	"{serviceQuotaRequestId}": "",
	"{name}":                  "",
	"{subscription}":          "",
	"{task}":                  "",
}

// queryparams for specific resource types
var queryparams = map[string]map[string]string{
	"VsanConfigConstraints": {"num_hosts": "10", "provider": "AWS"},
	"Sddcs":                 {"validateOnly": "true"},
	"CompatibleSubnets":     {"linkedAccountId": "", "forceRefresh": "true"},
	"CustomerSupport":       {"orgId": "", "userID": ""},
	"MsftLicensing":         {"termsId": ""},
	"OfferInstances":        {"product": "", "product_type": "", "region": "", "type": ""},
	"Tasks":                 {"$filter": "(updated gt 2022-07-15)"},
}

func BuildConfig() {
	config := core.ResourceFactory{
		ApiName:   name,
		ApiScheme: scheme,
		ApiUrl:    url,
		Resources: resources,
		Params:    params,
	}
	fmt.Println("#### Auto-Generated Config file for ", name, " API ####")
	fmt.Println(utils.Marshal(config))
}
