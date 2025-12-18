// Copyright (c) ZStack.io, Inc.

package view

// GetCandidateVMForAttachingAffinityGroupView GetCandidateVMForAttachingAffinityGroup
type GetCandidateVMForAttachingAffinityGroupView struct {
	Inventories []VmInstanceInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

