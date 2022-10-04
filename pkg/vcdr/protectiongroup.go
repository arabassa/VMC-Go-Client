package vcdr

import "encoding/json"

type GetProtectionGroupsResponse struct {
	Cursor           string `json:"cursor"`
	ProtectionGroups []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"protection_groups"`
}

type ProtectionGroupDetails struct {
	Health       string `json:"health"`
	ID           string `json:"id"`
	MembersSpecs []struct {
		VcenterFolderPaths    []string `json:"vcenter_folder_paths"`
		VcenterID             string   `json:"vcenter_id"`
		VcenterTags           []string `json:"vcenter_tags"`
		VcenterVMNamePatterns []string `json:"vcenter_vm_name_patterns"`
	} `json:"members_specs"`
	Name            string `json:"name"`
	ProtectedSiteID string `json:"protected_site_id"`
	ScheduleSpecs   []struct {
		Name         string      `json:"name"`
		RetentionMin json.Number `json:"retention_min"`
		ScheduleSpec struct {
			DayOfMonth         int `json:"day_of_month"`
			DayOfMonthInterval int `json:"day_of_month_interval"`
			DayOfWeek          int `json:"day_of_week"`
			DayOfWeekInterval  int `json:"day_of_week_interval"`
			Hours              int `json:"hours"`
			HoursInterval      int `json:"hours_interval"`
			Minutes            int `json:"minutes"`
			MinutesInterval    int `json:"minutes_interval"`
			Month              int `json:"month"`
			MonthInterval      int `json:"month_interval"`
		} `json:"schedule_spec"`
	} `json:"schedule_specs"`
	SnapshotFrequencyType    string      `json:"snapshot_frequency_type"`
	SnapshotQuiescingEnabled bool        `json:"snapshot_quiescing_enabled"`
	SnapshotScheduleActive   bool        `json:"snapshot_schedule_active"`
	UsedGib                  json.Number `json:"used_gib"`
}

type GetProtectionGroupsSnapshotsResponse struct {
	Cursor    string `json:"cursor"`
	Snapshots []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"snapshots"`
}

type ProtectionGroupSnapshotDetails struct {
	CreationTimestamp   json.Number `json:"creation_timestamp"`
	ExpirationTimestamp json.Number `json:"expiration_timestamp"`
	FailedVMSnapCount   int         `json:"failed_vm_snap_count"`
	ID                  string      `json:"id"`
	Name                string      `json:"name"`
	TotalUsedDataGib    json.Number `json:"total_used_data_gib"`
	TriggerType         string      `json:"trigger_type"`
	UniqueUsedDataGib   json.Number `json:"unique_used_data_gib"`
	VMCount             int         `json:"vm_count"`
}
