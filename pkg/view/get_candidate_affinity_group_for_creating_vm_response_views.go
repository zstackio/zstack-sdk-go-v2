// Copyright (c) ZStack.io, Inc.

package view

// GetCandidateAffinityGroupForCreatingVmView GetCandidateAffinityGroupForCreatingVm
type GetCandidateAffinityGroupForCreatingVmView struct {
	Inventories []AffinityGroupInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

