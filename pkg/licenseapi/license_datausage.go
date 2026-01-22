package licenseapi

type UsageData struct {
	// FeatureUsage contains the usage of features
	FeatureUsage map[string]FeatureUsage `json:"featureUsage"`

	// ResourceUsage contains the usage of resources
	ResourceUsage map[string]ResourceCount `json:"resourceUsage"`

	// Details contains the details of the usage data
	Details UsageDataDetails `json:"details"`
}

type UsageDataDetails struct {
	// Nodes contains the details of the nodes
	Nodes []NodeInfo `json:"nodes"`

	// VClusters contains the details of the virtual clusters
	VClusters []VirtualClusterInfo `json:"vClusters"`
}

type FeatureUsage struct {
	Used   bool
	Status string
}

type NodeInfo struct {
	MachineID         string            `json:"machine_id"`
	CreationTimestamp string            `json:"creation_timestamp"`
	Capacity          map[string]string `json:"capacity"`
}

type VirtualClusterInfo struct {
	UID               string   `json:"uid"`
	Name              string   `json:"name"`
	Namespace         string   `json:"namespace"`
	CreationTimestamp string   `json:"creation_timestamp"`
	IsAvailable       bool     `json:"is_available"`
	NodeMachineIDs    []string `json:"node_machine_ids"`
}
