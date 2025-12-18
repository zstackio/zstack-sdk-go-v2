// Copyright (c) ZStack.io, Inc.

package view

// GetPciDeviceCandidatesForNewCreateVmView GetPciDeviceCandidatesForNewCreateVm
type GetPciDeviceCandidatesForNewCreateVmView struct {
	Inventories []PciDeviceInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

