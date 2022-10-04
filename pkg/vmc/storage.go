package vmc

type VsanConfigConstraints struct {
	MaxCapacity                int64   `json:"max_capacity"`
	MinCapacity                int64   `json:"min_capacity"`
	NumHosts                   int64   `json:"num_hosts"`
	RecommendedCapacities      []int64 `json:"recommended_capacities"`
	SupportedCapacityIncrement int64   `json:"supported_capacity_increment"`
}

type VsanDiskgroupMapping struct {
	MappingType string `json:"mapping_type"`
}

type VsanClusterReconfigConstraints struct {
	AvailableCapacities map[string]VsanAvailableCapacity `json:"available_capacities"`
	DefaultCapacities   map[string]VsanAvailableCapacity `json:"default_capacities"`
	DefaultConfigBiasId string                           `json:"default_reconfig_bias_id"`
	Hosts               int                              `json:"hosts"`
	ReconfigBiases      []VsanClusterReconfigBias        `json:"reconfig_biases"`
}

type VsanAvailableCapacity struct {
	Cost    string `json:"cost"`
	Quality string `json:"quality"`
	Size    int64  `json:"size"`
}

type VsanClusterReconfigBias struct {
	FullDescription  string `json:"full_description"`
	Id               string `json:"id"`
	ShortDescription string `json:"short_description"`
}
