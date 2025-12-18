// Copyright (c) ZStack.io, Inc.

package view

// GetCandidateZonesClustersHostsForCreatingVmView GetCandidateZonesClustersHostsForCreatingVm
type GetCandidateZonesClustersHostsForCreatingVmView struct {
	Zones []ZoneInventoryView `json:"zones,omitempty"`
	Clusters []ClusterInventoryView `json:"clusters,omitempty"`
	Hosts []HostInventoryView `json:"hosts,omitempty"`
	Success bool `json:"success,omitempty"`
}

