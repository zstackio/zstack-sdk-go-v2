// Copyright (c) ZStack.io, Inc.

package view

// GetScsiLunCandidatesForAttachingVmView GetScsiLunCandidatesForAttachingVm
type GetScsiLunCandidatesForAttachingVmView struct {
	Inventories []ScsiLunInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

