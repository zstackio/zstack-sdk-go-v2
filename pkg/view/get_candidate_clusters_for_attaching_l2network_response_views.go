// Copyright (c) ZStack.io, Inc.

package view

// GetCandidateClustersForAttachingL2NetworkView GetCandidateClustersForAttachingL2Network
type GetCandidateClustersForAttachingL2NetworkView struct {
	Inventories []ClusterInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

