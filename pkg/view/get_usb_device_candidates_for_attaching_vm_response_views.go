// Copyright (c) ZStack.io, Inc.

package view

// GetUsbDeviceCandidatesForAttachingVmView GetUsbDeviceCandidatesForAttachingVm
type GetUsbDeviceCandidatesForAttachingVmView struct {
	Inventories []UsbDeviceInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

