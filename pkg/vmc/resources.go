package vmc

import (
	"reflect"
)

var VMCResourceMap = map[string]reflect.Type{
	"Org":                            initType(new(Org)),
	"Orgs":                           initType(new(Orgs)),
	"OrgPaymentMethods":              initType(new(PaymentMethodInfos)),
	"OrgProviders":                   initType(new(AwsCloudProviders)),
	"Reservations":                   initType(new(Reservations)),
	"Sddcs":                          initType(new(Sddcs)),
	"Sddc":                           initType(new(Sddc)),
	"ProvisionSpec":                  initType(new(ProvisionSpec)),
	"Cluster":                        initType(new(Cluster)),
	"VsanConfigConstraints":          initType(new(VsanConfigConstraints)),
	"VsanClusterReconfigConstraints": initType(new(VsanClusterReconfigConstraints)),
	"StorageSpecs":                   initType(new(VsanDiskgroupMapping)),
	"ServiceQuotaRequest":            initType(new(ServiceQuotaRequest)),
	"ServiceQuotaRequests":           initType(new(ServiceQuotaRequests)),
	"AccountLink":                    initType(new(AccountLink)),
	"ConnectedAccounts":              initType(new(AwsCustomerConnectedAccounts)),
	"CompatibleSubnets":              initType(new(AwsCompatibleSubnets)),
	"AwsSddcConnections":             initType(new(AwsSddcConnections)),
	"Credential":                     initType(new(NewCredential)),
	"Credentials":                    initType(new(NewCredentials)),
	"CustomerSupport":                initType(new(SddcInput)),
	"MsftLicensing":                  initType(new(SddcInput)),
	"Subscriptions":                  initType(new(Subscriptions)),
	"Subscription":                   initType(new(Subscription)),
	"OfferInstances":                 initType(new(OfferInstances)),
	"Products":                       initType(new(SubscriptionProducts)),
	"Tbrs":                           initType(new(SupportWindows)),
	"Tasks":                          initType(new(Tasks)),
	"Task":                           initType(new(Task)),
	"NetworkConnChecker":             initType(new(ConnectivityValidationGroup)),
}

func initType(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}
