// Copyright (c) ZStack.io, Inc.

package view

// GetPciDeviceCandidatesForAttachingVmView GetPciDeviceCandidatesForAttachingVm
type GetPciDeviceCandidatesForAttachingVmView struct {
	Inventories []PciDeviceInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

