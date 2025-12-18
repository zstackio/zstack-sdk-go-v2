// Copyright (c) ZStack.io, Inc.

package view

// UpdateHostIpmiEventView UpdateHostIpmiEvent
type UpdateHostIpmiEventView struct {
	HostIpmiInventory HostIpmiInventoryView `json:"hostIpmiInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

