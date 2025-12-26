// Copyright (c) ZStack.io, Inc.

package view

// GetCandidateL2NetworksForAttachingClusterView GetCandidateL2NetworksForAttachingCluster
type GetCandidateL2NetworksForAttachingClusterView struct {
	Inventories []L2NetworkDataView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

