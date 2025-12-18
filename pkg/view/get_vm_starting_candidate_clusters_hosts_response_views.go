// Copyright (c) ZStack.io, Inc.

package view

// GetVmStartingCandidateClustersHostsView GetVmStartingCandidateClustersHosts
type GetVmStartingCandidateClustersHostsView struct {
	Hosts []HostInventoryView `json:"hosts,omitempty"`
	Clusters []ClusterInventoryView `json:"clusters,omitempty"`
}

