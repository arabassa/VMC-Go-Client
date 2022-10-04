package vmc

type Sddcs []Sddc

type Sddc struct {
	UserId             string             `json:"user_id"`
	UserName           string             `json:"user_name"`
	Created            string             `json:"created"`
	Version            int                `json:"version"`
	Id                 string             `json:"id"`
	UpdatedByUserId    string             `json:"updated_by_user_id"`
	UpdatedByUserName  string             `json:"updated_by_user_name"`
	Updated            string             `json:"updated"`
	AccountLinkState   string             `json:"account_link_state"`
	ExpirationDate     string             `json:"expiration_date"`
	Name               string             `json:"name"`
	OrgId              string             `json:"org_id"`
	Provider           string             `json:"provider"`
	SddcResourceConfig SddcResourceConfig `json:"resource_config"`
	SddcAccessState    string             `json:"sddc_access_state"`
	SddcState          string             `json:"sddc_state"`
	SddcType           string             `json:"sddc_type"`
}

type SddcResourceConfig struct {
	Provider                    string                 `json:"provider"`
	AvailabilityZones           []string               `json:"availability_zones"`
	CloudPassword               string                 `json:"cloud_password"`
	CloudUserGroup              string                 `json:"cloud_user_group"`
	CloudUserName               string                 `json:"cloud_user_name"`
	Clusters                    []Cluster              `json:"clusters"`
	CustomProps                 map[string]string      `json:"custom_properties"`
	CvdsEnabled                 bool                   `json:"cdvs_enabled"`
	DeploymentTime              string                 `json:"deployment_time"`
	DnsMgmtPrivateIP            bool                   `json:"dns_with_management_vm_private_ip"`
	EsxClusterId                string                 `json:"esx_cluster_id"`
	EsxHostSubnet               string                 `json:"esx_host_subnet"`
	EsxHosts                    []AwsEsxHost           `json:"esx_hosts"`
	ManagementDs                string                 `json:"management_ds"`
	ManagementRp                string                 `json:"management_rp"`
	MgmtApplianceNetworkName    string                 `json:"mgmt_appliance_network_name"`
	MgwId                       string                 `json:"mgw_id"`
	MsftLicensingConfig         MsftLicensingConfig    `json:"msft_license_config"`
	NfsMode                     bool                   `json:"nfs_mode"`
	NsxApiPubEndpoint           string                 `json:"nsx_api_public_endpoint_url"`
	NsxCloudAdmin               string                 `json:"nsx_cloud_admin"`
	NsxCloudAdminPassword       string                 `json:"nsx_cloud_admin_password"`
	NsxControllerIps            []string               `json:"nsx_controller_ips"`
	NsxMgrLoginUrl              string                 `json:"nsx_mgr_login_url"`
	NsxMgrMgmtIp                string                 `json:"nsx_mgr_management_ip"`
	NsxMgrUrl                   string                 `json:"nsx_mgr_url"`
	NsxNativeClient             CspOauthClient         `json:"nsx_native_client"`
	Nsxt                        bool                   `json:"nsxt"`
	NsxtAddons                  NsxtAddons             `json:"nsxt_addons"`
	OutpostConfig               OutpostConfig          `json:"outpost_config"`
	PopAgentXeniConnection      PopAgentXeniConnection `json:"pop_agent_xeni_connection"`
	Cgws                        []string               `json:"cgws"`
	PscIp                       string                 `json:"psc_ip"`
	PscMgmtIp                   string                 `json:"psc_management_ip"`
	PscUrl                      string                 `json:"psc_url"`
	Region                      string                 `json:"region"`
	SddcDesiredState            bool                   `json:"sddc_desired_state"`
	SddcId                      string                 `json:"sddc_id"`
	SddcManifest                SddcManifest           `json:"sddc_manifest"`
	SddcNetworks                []string               `json:"sddc_networks"`
	SddcSecurity                SddcSecurity           `json:"sddc_security"`
	SddcSize                    SddcSize               `json:"sddc_size"`
	SkipCreatingVxlan           bool                   `json:"skip_creating_vxlan"`
	SsoDomain                   string                 `json:"sso_domain"`
	TwoHostnameVcDeployment     bool                   `json:"two_hostname_vc_deployment"`
	VcContainerizedPermsEnabled bool                   `json:"vc_containerized_permissions_enabled"`
	VcInstanceId                string                 `json:"vc_instance_id"`
	VcIp                        string                 `json:"vc_ip"`
	VcMgmtIp                    string                 `json:"vc_management_ip"`
	VcPubIp                     string                 `json:"vc_public_ip"`
	VcUrl                       string                 `json:"vc_url"`
	VlcmEnabled                 bool                   `json:"vlcm_enabled"`
	VxlanSubnet                 string                 `json:"vxlan_subnet"`
	WitnessAz                   string                 `json:"witness_availability_zone"`
	SddcLinkConfig              []SddcLinkConfig       `json:"account_link_sddc_config"`
	BackupRestoreBucket         string                 `json:"backup_restore_bucket"`
	InstanceProfileInfo         InstanceProfileInfo    `json:"esx_instance_profile"`
	KmsVpcEndpoint              KmsVpcEndpoint         `json:"kms_vpc_endpoint"`
	MaxNumPublicIp              int                    `json:"max_num_public_ip"`
	NsxServiceClient            CspOauthClient         `json:"nsx_service_Client"`
	PublicIpPool                []SddcPublicIp         `json:"public_ip_pool"`
	VpcInfo                     VpcInfo                `json:"vpc_info"`
	VpcInfoPeeredAgent          VpcInfo                `json:"vpc_info_peered_agent"`
	VsanEncryptionConfig        VsanEncryptionConfig   `json:"vsan_encryption_config"`
}

type Cluster struct {
	ClusterId                   string                        `json:"cluster_id"`
	AwsKmsInfo                  AwsKmsInfo                    `json:"aws_kms_info"`
	EntityCapacity              EntityCapacity                `json:"cluster_capacity"`
	ClusterName                 string                        `json:"cluster_name"`
	ClusterState                string                        `json:"cluster_state"`
	CustomProps                 map[string]string             `json:"custom_properties"`
	EsxHostInfo                 EsxHostInfo                   `json:"esx_host_info"`
	EsxHostList                 []AwsEsxHost                  `json:"esx_host_list"`
	HostCpuCores                int32                         `json:"host_cpu_cores_count"`
	HyperThreadingEnabled       bool                          `json:"hyper_threading_enabled"`
	MsftLicensingConfig         MsftLicensingConfig           `json:"msft_license_config"`
	PartitionPlacementGroupInfo []PartitionPlacementGroupInfo `json:"partition_placement_group_info"`
	VsanWitness                 AwsWtinessEsx                 `json:"vsan_witness"`
	WcpDetails                  WcpDetails                    `json:"wcp_details"`
}

type AwsKmsInfo struct {
	AmazonResourceName string `json:"amazon_resource_name"`
}

type EntityCapacity struct {
	CpuCapacityGhz     float32 `json:"cpu_capacity_ghz"`
	MemoryCapacityGib  int     `json:"memory_capacity_gib"`
	NumSockets         int     `json:"number_of_sockets"`
	NumSsds            int     `json:"number_of_ssds"`
	StorageCapacityGib int     `json:"storage_capacity_gib"`
	TotalNumCores      int     `json:"total_number_of_cores"`
}

type EsxHostInfo struct {
	InstanceType string `json:"instance_type"`
}

type AwsEsxHost struct {
	Provider             string            `json:"provider"`
	AvailabilityZone     string            `json:"availability_zone"`
	CustomProps          map[string]string `json:"custom_properties"`
	EsxState             string            `json:"esx_state"`
	Hostname             string            `json:"hostname"`
	InstanceType         string            `json:"instance_type"`
	MacAddress           string            `json:"mac_address"`
	Name                 string            `json:"name"`
	EsxId                string            `json:"esx_id"`
	EniList              []EniInfo         `json:"eni_list"`
	InternalPublicIpPool []SddcPublicIp    `json:"internal_public_ip_pool"`
	PartitionNumber      int32             `json:"partition_number"`
	XeniInfo             XeniInfo          `json:"xeni_info"`
}

type EniInfo struct {
	AssociationId        string   `json:"association_id"`
	AttachmentId         string   `json:"attachment_id"`
	DeviceIndex          int      `json:"device_index"`
	Id                   string   `json:"id"`
	InstanceId           string   `json:"instance_id"`
	MacAddress           string   `json:"mac_address"`
	PortGroup            string   `json:"portgroup"`
	PrivateIp            string   `json:"private_ip"`
	PublicIps            []string `json:"public_ips"`
	SecondaryIps         []string `json:"secondary_ips"`
	SecurityGroupId      string   `json:"security_group_id"`
	SourceDestCheckFalse bool     `json:"source_dest_check_false"`
	SubnetId             string   `json:"subnet_id"`
	VmkId                string   `json:"vmk_id"`
}

type SddcPublicIp struct {
	PublicIp            string `json:"public_ip"`
	AllocationId        string `json:"allocation_id"`
	AssociatedPrivateIp string `json:"associated_private_ip"`
	DnatRuleId          string `json:"dnat_rule_id"`
	Name                string `json:"name"`
	SnatRuleId          string `json:"snat_rule_id"`
}

type XeniInfo struct {
	AssociationId string `json:"association_id"`
	TrunkEniId    string `json:"trunk_eni_id"`
	XeniId        string `json:"x_eni_id"`
}

type MsftLicensingConfig struct {
	AcademicLicense  bool   `json:"academic_license"`
	MssqlLicensing   string `json:"mssql_licensing"`
	WindowsLicensing string `json:"windows_licensing"`
}

type PartitionPlacementGroupInfo struct {
	AvailabilityZone    string   `json:"availability_zone"`
	PartitionGroupNames []string `json:"partition_group_names"`
}

type AwsWtinessEsx struct {
	Provider   string  `json:"provider"`
	EnumState  string  `json:"enum_state"`
	EsxId      string  `json:"esx_id"`
	Hostname   string  `json:"hostname"`
	MacAddress string  `json:"mac_address"`
	Name       string  `json:"name"`
	EniInfo    EniInfo `json:"eni_info"`
	InstanceId string  `json:"instance_id"`
}

type WcpDetails struct {
	EgressCidr  string `json:"egress_cidr"`
	IngressCidr string `json:"ingress_cidr"`
	PodCidr     string `json:"pod_cidr"`
	ServiceCidr string `json:"service_cidr"`
	WcpStatus   string `json:"wcp_status"`
}

type CspOauthClient struct {
	Id             string   `json:"id"`
	Secret         string   `json:"secret"`
	AccessTokenTtl int      `json:"accessTokenTTL"`
	GrantTypes     []string `json:"grant_types"`
	OrgId          string   `json:"org_id"`
	RedirectUri    string   `json:"redirect_uri"`
	RedirectUris   []string `json:"redirect_uris"`
	RefresTokenTtl int      `json:"refreshTokenTTL"`
	ResourceLink   string   `json:"resource_link"`
}

type NsxtAddons struct {
	EnableNsxAdvancedAddon bool `json:"enable_nsx_advanced_addon"`
}

type OutpostConfig struct {
	Mocked    bool   `json:"mocked"`
	OutpostId string `json:"outpost_id"`
}

type PopAgentXeniConnection struct {
	DefaultSubnetRoute string  `json:"default_subnet_route"`
	EniInfo            EniInfo `json:"eni_info"`
}

type SddcManifest struct {
	EbsBackedVsanConfig EbsBackedVsanConfig `json:"ebs_backed_vsan_config"`
	EsxAmiInfo          AmiInfo             `json:"esx_ami"`
	NsxtAmiInfo         AmiInfo             `json:"esx_nsxt_ami"`
	GlcmBundle          GlcmBundle          `json:"glcm_bundle"`
	Metadata            Metadata            `json:"metadata"`
	PopInfo             PopInfo             `json:"pop_info"`
	VmcInternalVersion  string              `json:"vmc_internal_info"`
	VmcVersion          string              `json:"vmc_version"`
	VsanWitnessAmi      AmiInfo             `json:"vsan_witness_ami"`
}

type EbsBackedVsanConfig struct {
	InstanceType string `json:"instance_type"`
}

type AmiInfo struct {
	Id           string `json:"id"`
	InstanceType string `json:"instance_type"`
	Name         string `json:"name"`
	Region       string `json:"region"`
}

type GlcmBundle struct {
	Id       string `json:"id"`
	S3Bucket string `json:"s3bucket"`
}

type Metadata struct {
	CycleId   string `json:"cycle_id"`
	Timestamp string `json:"timestamp"`
}

type PopInfo struct {
	AmiInfos        map[string]PopAmiInfo     `json:"ami_infos"`
	CreatedAt       string                    `json:"created_at"`
	Id              string                    `json:"id"`
	ManifestVersion string                    `json:"manifest_version"`
	ServiceInfos    map[string]PopServiceInfo `json:"service_infos"`
}

type PopAmiInfo struct {
	Id           string `json:"id"`
	InstanceType string `json:"instance_type"`
	Name         string `json:"name"`
	Region       string `json:"region"`
	Type         string `json:"type"`
}

type PopServiceInfo struct {
	Build   string `json:"build"`
	Cln     string `json:"cln"`
	Service string `json:"service"`
	Version string `json:"version"`
}

type SddcSecurity struct {
	Hardened bool   `json:"hardened"`
	Profile  string `json:"profile"`
}

type SddcSize struct {
	NsxSize string `json:"nsx_size"`
	Size    string `json:"size"`
	VcSize  string `json:"vc_size"`
}

type SddcLinkConfig struct {
	ConnectedAccountId string   `json:"connected_account_id"`
	CustomerSubnetIds  []string `json:"customer_subnet_ids"`
}

type InstanceProfileInfo struct {
	InstanceProfileName string `json:"instance_profile_name"`
	PolicyName          string `json:"policy_name"`
	RoleName            string `json:"role_name"`
}

type KmsVpcEndpoint struct {
	NetworkInterfaceIds []string `json:"network_interface_ids"`
	VpcEndpointId       string   `json:"vpc_endpoint_id"`
}

type VpcInfo struct {
	ApiAssociationId         string                    `json:"api_association_id"`
	ApiSubnetId              string                    `json:"api_subnet_id"`
	AssociationId            string                    `json:"association_id"`
	AvailableZones           []AvailableZoneInfo       `json:"available_zones"`
	EdgeAssociationId        string                    `json:"edge_association_id"`
	EdgeSubnetId             string                    `json:"edge_subnet_id"`
	EsxPublicSecurityGroupId string                    `json:"esx_public_security_group_id"`
	EsxSecurityGroupId       string                    `json:"esx_security_group_id"`
	Id                       string                    `json:"id"`
	InternetGatewayId        string                    `json:"internet_gateway_id"`
	NetworkType              string                    `json:"network_type"`
	PeeringConnectionId      string                    `json:"peering_connection_id"`
	PrivateAssociationId     string                    `json:"private_association_id"`
	PrivateSubnetId          string                    `json:"private_subnet_id"`
	Provider                 string                    `json:"provider"`
	RouteTableId             string                    `json:"route_table_id"`
	RouteTables              map[string]RouteTableInfo `json:"routetables"`
	SecurityGroupId          string                    `json:"security_group_id"`
	SubnetId                 string                    `json:"subnet_id"`
	TgwIps                   map[string][]string       `json:"tgw_ips"`
	TrafficGroupEdgeVmIps    []string                  `json:"traffic_group_edge_vm_ips"`
	VcdrEnis                 []EniInfo                 `json:"vcdr_enis"`
	VgwId                    string                    `json:"vgw_id"`
	VgwRouteTableId          string                    `json:"vgw_route_table_id"`
	VifIds                   []string                  `json:"vif_ids"`
	VmSecurityGroupId        string                    `json:"vm_security_group_id"`
	VpcCidr                  string                    `json:"vpc_cidr"`
}

type AvailableZoneInfo struct {
	Id      string   `json:"id"`
	Subnets []Subnet `json:"subnets"`
}

type Subnet struct {
	Name        string                 `json:"name"`
	RouteTables []SubnetRouteTableInfo `json:"route_tables"`
	SubnetId    string                 `json:"subnet_id"`
}

type SubnetRouteTableInfo struct {
	AssociationId string `json:"association_id"`
	RouteTableId  string `json:"routetable_id"`
	SubnetId      string `json:"subnet_id"`
}

type RouteTableInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type VsanEncryptionConfig struct {
	Certificate string `json:"certificate"`
	Port        int    `json:"port"`
}

type ProvisionSpec struct {
	Provider map[string]SddcConfigSpec `json:"provider"`
}

type SddcConfigSpec struct {
	RegionDisplayNames map[string]string     `json:"region_display_names"`
	SddcTypeConfigSpec map[string]ConfigSpec `json:"sddc_type_config_spec"`
}

type ConfigSpec struct {
	ExpiryInDays              int                                    `json:"expiry_in_days"`
	Availability              map[string][]InstanceTypeConfig        `json:"availability"`
	SddcSizes                 []string                               `json:"sddc_sizes"`
	PoweredBy                 map[string]PoweredByInstanceTypeConfig `json:"powered_by"`
	PoweredByOutpostAvailable bool                                   `json:"powered_by_outpost_available"`
}

type InstanceTypeConfig struct {
	InstanceType                   string         `json:"instance_type"`
	DisplayName                    string         `json:"display_name"`
	Label                          string         `json:"label"`
	Description                    string         `json:"description"`
	EntityCapacity                 EntityCapacity `json:"entity_capacity"`
	Hosts                          []int          `json:"hosts"`
	CpuCores                       []int          `json:"cpu_cores"`
	SupportedCpuCores              int            `json:"supported_cpu_cores"`
	HyperThreadingSupported        bool           `json:"hyper_threading_supported"`
	InstanceProvisioningErrorCause string         `json:"instanceProvisioningErrorCause"`
	TwoNodeSupported               bool           `json:"two_node_supported"`
}

type PoweredByInstanceTypeConfig struct {
	Description                    string         `json:"description"`
	DisplayName                    string         `json:"display_name"`
	EntityCapacity                 EntityCapacity `json:"entity_capacity"`
	InstanceType                   string         `json:"instance_type"`
	Label                          string         `json:"label"`
	CpuCores                       []int          `json:"cpu_cores"`
	Hosts                          []int          `json:"hosts"`
	HyperThreadingSupported        bool           `json:"hyper_threading_supported"`
	InstanceProvisioningErrorCause string         `json:"InstanceProvisioningErrorCause"`
	PoweredById                    string         `json:"powered_by_id"`
	PoweredByType                  string         `json:"powered_by_type"`
}
