package vmc

import (
	"fmt"
	"vmc-go-client/pkg/core"
	"vmc-go-client/pkg/utils"
)

const name = "VMC"
const scheme = "https"
const url = "vmc.vmware.com"

var resources = map[string]core.Endpoint{
	"Org":                            {"/vmc/api/orgs/{org}", "GET", initType(new(Org)), queryparams["Org"]},
	"Orgs":                           {"/vmc/api/orgs", "GET", initType(new(Orgs)), queryparams["Orgs"]},
	"OrgPaymentMethods":              {"/vmc/api/orgs/{org}/payment-methods", "GET", initType(new(PaymentMethodInfos)), queryparams["OrgPaymentMethods"]},
	"OrgProviders":                   {"/vmc/api/orgs/{org}/providers", "GET", initType(new(AwsCloudProviders)), queryparams["OrgProviders"]},
	"Reservations":                   {"/vmc/api/orgs/{org}/reservations", "GET", initType(new(Reservations)), queryparams["Reservations"]},
	"Sddcs":                          {"/vmc/api/orgs/{org}/sddcs", "GET", initType(new(Sddcs)), queryparams["Sddcs"]},
	"Sddc":                           {"/vmc/api/orgs/{org}/sddcs/{sddc}", "GET", initType(new(Sddc)), queryparams["Sddc"]},
	"ProvisionSpec":                  {"/vmc/api/orgs/{org}/sddcs/provision-spec", "GET", initType(new(ProvisionSpec)), queryparams["ProvisionSpec"]},
	"Cluster":                        {"/vmc/api/orgs/{org}/sddcs/{sddc}/primarycluster", "GET", initType(new(Cluster)), queryparams["Cluster"]},
	"VsanConfigConstraints":          {"/vmc/api/orgs/{org}/storage/cluster-constraints", "GET", initType(new(VsanConfigConstraints)), queryparams["VsanConfigConstraints"]},
	"VsanClusterReconfigConstraints": {"/vmc/api/orgs/{org}/sddcs/{sddc}/clusters/{cluster}/config/constraints", "GET", initType(new(VsanClusterReconfigConstraints)), queryparams["VsanClusterReconfigConstraints"]},
	"StorageSpecs":                   {"/vmc/api/orgs/{org}/sddcs/{sddc}/clusters/{cluster}/esxs/{esx}/storage-spec", "GET", initType(new(VsanDiskgroupMapping)), queryparams["StorageSpecs"]},
	"ServiceQuotaRequest":            {"/vmc/api/orgs/{org}/aws/resources/servicequotas/requests/{serviceQuotaRequestId}Id", "GET", initType(new(ServiceQuotaRequest)), queryparams["ServiceQuotaRequest"]},
	"ServiceQuotaRequests":           {"/vmc/api/orgs/{org}/aws/resources/servicequotas/requests", "GET", initType(new(ServiceQuotaRequests)), queryparams["ServiceQuotaRequests"]},
	"AccountLink":                    {"/vmc/api/orgs/{org}/account-link", "GET", initType(new(AccountLink)), queryparams["AccountLink"]},
	"ConnectedAccounts":              {"/vmc/api/orgs/{org}/account-link/connected-accounts", "GET", initType(new(AwsCustomerConnectedAccounts)), queryparams["ConnectedAccounts"]},
	"CompatibleSubnets":              {"/vmc/api/orgs/{org}/account-link/compatible-subnets", "GET", initType(new(AwsCompatibleSubnets)), queryparams["CompatibleSubnets"]},
	"AwsSddcConnections":             {"/vmc/api/orgs/{org}/account-link/sddc-connections", "GET", initType(new(AwsSddcConnections)), queryparams["AwsSddcConnections"]},
	"Credential":                     {"/vmc/api/orgs/{org}/sddcs/{sddc}/addons/" + Credentials_GET_ADDON_TYPE_HCX + "/credentials/{name}", "GET", initType(new(NewCredential)), queryparams["Credential"]},
	"Credentials":                    {"/vmc/api/orgs/{org}/sddcs/{sddc}/addons/" + Credentials_GET_ADDON_TYPE_HCX + "/credentials", "GET", initType(new(NewCredentials)), queryparams["Credentials"]},
	"CustomerSupport":                {"/vmc/api/customer-support/inputs", "GET", initType(new(SddcInput)), queryparams["CustomerSupport"]},
	"MsftLicensing":                  {"/vmc/api/orgs/{org}/tos", "GET", initType(new(SddcInput)), queryparams["MsftLicensing"]},
	"Subscriptions":                  {"/vmc/api/orgs/{org}/subscriptions", "GET", initType(new(Subscriptions)), queryparams["Subscriptions"]},
	"Subscription":                   {"/vmc/api/orgs/{org}/subscriptions/{subscription}", "GET", initType(new(Subscription)), queryparams["Subscription"]},
	"OfferInstances":                 {"/vmc/api/orgs/{org}/subscriptions/offer-instances", "GET", initType(new(OfferInstances)), queryparams["OfferInstances"]},
	"Products":                       {"/vmc/api/orgs/{org}/subscriptions/products", "GET", initType(new(SubscriptionProducts)), queryparams["Products"]},
	"Tbrs":                           {"/vmc/api/orgs/{org}/tbrs/support-window", "GET", initType(new(SupportWindows)), queryparams["Tbrs"]},
	"Tasks":                          {"/vmc/api/orgs/{org}/tasks", "GET", initType(new(Tasks)), queryparams["Tasks"]}, //heavy call, better use it with filter-query params
	"Task":                           {"/vmc/api/orgs/{org}/tasks/{task}", "GET", initType(new(Task)), queryparams["Task"]},
	"NetworkConnChecker":             {"/vmc/api/orgs/{org}/sddcs/{sddc}/networking/connectivity-tests", "GET", initType(new(ConnectivityValidationGroup)), queryparams["NetworkConnChecker"]},
}

//params to set variables and construct our API resources
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

//queryparams for specific resource types
var queryparams = map[string]map[string]string{
	"VsanConfigConstraints": {"num_hosts": "10", "provider": "AWS"},
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
