// Copyright (c) ZStack.io, Inc.

package view

// GetHostPowerStatusEventView GetHostPowerStatusEvent
type GetHostPowerStatusEventView struct {
	Inventory HostIpmiInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

