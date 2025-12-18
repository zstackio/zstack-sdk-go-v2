// Copyright (c) ZStack.io, Inc.

package view

// GetCandidateAffinityGroupForAttachingVmView GetCandidateAffinityGroupForAttachingVm
type GetCandidateAffinityGroupForAttachingVmView struct {
	Inventories []AffinityGroupInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

