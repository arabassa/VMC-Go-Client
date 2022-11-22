package vcdr

import (
	"fmt"
	core "vmc-go-client/pkg/core"
	"vmc-go-client/pkg/utils"
)

const name = "VCDR"
const scheme = "https"
const url = "vcdr-x-x-x-x.app.vcdr.vmware.com"

var readMethods = []string{"GET"}

var resources = map[string]core.Endpoint{
	"CloudFileSystems":         {"/api/vcdr/v1alpha/cloud-file-systems", readMethods, initType(new(CloudFileSystemsResponse)), queryparams["CloudFileSystems"]},
	"CloudFileSystem":          {"/api/vcdr/v1alpha/cloud-file-systems/{cloud_file_system_id}", readMethods, initType(new(CloudFileSystemDetails)), queryparams["CloudFileSystem"]},
	"ProtectedSites":           {"/api/vcdr/v1alpha/cloud-file-systems/{cloud_file_system_id}/protected-sites", readMethods, initType(new(GetProtectedSitesResponse)), queryparams["ProtectedSites"]},
	"ProtectedSite":            {"/api/vcdr/v1alpha/cloud-file-systems/{cloud_file_system_id}/protected-sites/{protected_site_id}", readMethods, initType(new(ProtectedSiteDetails)), queryparams["ProtectedSite"]},
	"ProtectedVms":             {"/api/vcdr/v1alpha/cloud-file-systems/{cloud_file_system_id}/protected-vms", readMethods, initType(new(GetProtectedVirtualMachinesResponse)), queryparams["ProtectedVms"]},
	"ProtectionGroups":         {"/api/vcdr/v1alpha/cloud-file-systems/{cloud_file_system_id}/protection-groups", readMethods, initType(new(GetProtectionGroupsResponse)), queryparams["ProtectedGroups"]},
	"ProtectionGroup":          {"/api/vcdr/v1alpha/cloud-file-systems/{cloud_file_system_id}/protection-groups/{protection_group_id}", readMethods, initType(new(ProtectionGroupDetails)), queryparams["ProtectedGroup"]},
	"ProtectionGroupSnapshots": {"/api/vcdr/v1alpha/cloud-file-systems/{cloud_file_system_id}/protection-groups/{protection_group_id}/snapshots", readMethods, initType(new(GetProtectionGroupsSnapshotsResponse)), queryparams["ProtectionGroupSnapshots"]},
	"ProtectionGroupSnapshot":  {"/api/vcdr/v1alpha/cloud-file-systems/{cloud_file_system_id}/protection-groups/{protection_group_id}/snapshots/{snapshot_id}", readMethods, initType(new(ProtectionGroupSnapshotDetails)), queryparams["ProtectionGroupSnapshotDetails"]},
	"RecoverySddcs":            {"/api/vcdr/v1alpha/recovery-sddcs", readMethods, initType(new(GetRecoverySddcResponse)), queryparams["RecoverySddcs"]},
	"RecoverySddcDetails":      {"/api/vcdr/v1alpha/recovery-sddcs/{recovery_sddc_id}", readMethods, initType(new(RecoverySddcDetails)), queryparams["RecoverySddcDetails"]},
}

//params to set variables and construct our API resources
var params = map[string]string{
	"{cloud_file_system_id}": "",
	"{protected_site_id}":    "",
	"{protection_group_id}":  "",
	"{snapshot_id}":          "",
	"{recovery_sddc_id}":     "",
}

//queryparams for specific resource types
var queryparams = map[string]map[string]string{
	"ProtectedVms": {"cursor": "", "limit": "0"},
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
