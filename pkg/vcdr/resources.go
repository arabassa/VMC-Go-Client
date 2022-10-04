package vcdr

import (
	"reflect"
)

var VCDRResourceMap = map[string]reflect.Type{
	"CloudFileSystem":          initType(new(CloudFileSystemDetails)),
	"CloudFileSystems":         initType(new(CloudFileSystemsResponse)),
	"ProtectedSites":           initType(new(GetProtectedSitesResponse)),
	"ProtectedSite":            initType(new(ProtectedSiteDetails)),
	"ProtectedVms":             initType(new(GetProtectedVirtualMachinesResponse)),
	"ProtectionGroups":         initType(new(GetProtectionGroupsResponse)),
	"ProtectionGroup":          initType(new(ProtectionGroupDetails)),
	"ProtectionGroupSnapshots": initType(new(GetProtectionGroupsSnapshotsResponse)),
	"ProtectionGroupSnapshot":  initType(new(ProtectionGroupSnapshotDetails)),
	"RecoverySddcs":            initType(new(GetRecoverySddcResponse)),
	"RecoverySddcDetails":      initType(new(RecoverySddcDetails)),
}

func initType(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}
